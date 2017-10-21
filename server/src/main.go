package main

import (
	"fmt"
	"container/list"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

var gameState = new(GameState)
var undoHistory = list.New()

func saveForUndo() {
	var history = History{gameState.threadContext, gameState.globalState}
	undoHistory.PushBack(history)
}

//TODO
func startLevel(levelName string) {
	//清空
	for e := undoHistory.Front(); e != nil; e = e.Next() {
		undoHistory.Remove(e)
	}
}
func areAllThreadsBlocked() bool {
	return false
}
func areAllThreadsFinished() bool {
	return false
}
func checkForVictoryConditions() {
	var howManyCriticalSections = 0

	for threadId, t := range gameState.level.threads {

		if IsThreadFinished(threadId) {
			continue
		}
		var thread = gameState.level.threads[threadId]
		var instructions = thread.Instructions
		var threadState = gameState.threadContext[threadId]
		var programCounter = threadState.ProgramCounter
		var currentInstruction = instructions[programCounter]
		fmt.Print(currentInstruction)
		//TODO
		//if currentInstruction.isCriticalSection {
		//	howManyCriticalSections++
		//}

		fmt.Println(t)
	}
	if howManyCriticalSections >= 2 {
		win("Two threads were in a critical section at the same time.")
		return
	}
	if areAllThreadsBlocked() {
		win("A deadlock occurred - all threads were blocked simultaneously.")
		return
	}

	if areAllThreadsFinished() {
		lose("All threads of the program ran to the end, so the program was successful. Try to sabotage the program before it finishes.")
		return
	}
}
func win(winInfo string) {
	fmt.Print(winInfo)
}
func lose(loseInfo string) {
	fmt.Print(loseInfo)
}

func expandThread(threadId int) {
	saveForUndo()
	gameState.threadContext[threadId].Expanded = true
}

func undo() {

	for e := undoHistory.Front(); e != nil; e = e.Next() {
		if e.Next() == nil {
			gameState.globalState = e.Value.(History).globalContext
			gameState.threadContext = e.Value.(History).threadContext
			undoHistory.Remove(e)
		}
	}
}

func stepThread(thread int) {
	//if IsLevelPristine() {
	////第一关
	////sendEvent('Gameplay', 'level-first-step', gameState.getLevelId());
	//}
	var program = gameState.GetProgramOfThread(thread)
	var threadState = gameState.threadContext[thread]
	var pc = threadState.ProgramCounter

	if IsThreadFinished(thread) {
		saveForUndo()
		if threadState.Expanded {
			//展开了的情况
			program[pc].GetExpandInstructions()[threadState.ExpProgramCounter].Execute(&gameState.globalState, &threadState)
		} else {
			program[pc].Execute(&gameState.globalState, &threadState)
		}
		checkForVictoryConditions()

	}
	fmt.Println(program, threadState, pc)

}

func IsThreadFinished(threadId int) bool {
	program := gameState.GetProgramOfThread(threadId)
	var maxInstructions = len(program)
	var threadState = gameState.threadContext[threadId]
	var pc = threadState.ProgramCounter

	return pc >= maxInstructions
}

func IsLevelPristine() bool {
	//return undoHistory.length == 0;
	return true
}

// View Model
type LevelInfo struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Programs    []Program `json:"programs"`
}
type Program struct {
	CanStepNext bool  `json:"canStepNext"`
	Current     []int `json:"current"`
}

// Movie is our sample data structure.
type Movie struct {
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Poster string `json:"poster"`
}

// movies contains our imaginary data source.
var movies = []Movie{
	{
		Name:   "Casablanca",
		Year:   1942,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/1.jpg",
	},
	{
		Name:   "Gone with the Wind",
		Year:   1939,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/2.jpg",
	},
	{
		Name:   "Citizen Kane",
		Year:   1941,
		Genre:  "Mystery",
		Poster: "https://iris-go.com/images/examples/mvc-movies/3.jpg",
	},
	{
		Name:   "The Wizard of Oz",
		Year:   1939,
		Genre:  "Fantasy",
		Poster: "https://iris-go.com/images/examples/mvc-movies/4.jpg",
	},
}
var p = []Program{
	{
		CanStepNext: true,
		Current:     []int{1, 2},
	},
	{
		CanStepNext: false,
		Current:     []int{1, 2},
	},
}

var levelInfo = LevelInfo{Title: "a", Description: "d", Programs: p}

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/api/level/{id}", func(ctx iris.Context) {
		//movies[0].Name = ctx.Params().Get("id")
		ctx.JSON(levelInfo)
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}
