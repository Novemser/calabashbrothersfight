package main

import (
	"fmt"
	"container/list"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var gameState = new(GameState)
var undoHistory = list.New()

func saveForUndo() {
	var history = History{gameState.threadContexts, gameState.globalState}
	undoHistory.PushBack(history)
}

//TODO
func StartLevel(levelName string) {
	//清空
	for e := undoHistory.Front(); e != nil; e = e.Next() {
		undoHistory.Remove(e)
	}

	// 开始
	var level = Level1
	gameState.globalState = *level.GlobalContext
	gameState.threadContexts = level.ThreadContexts
	gameState.level = *level

	for i := 0; i < len(level.ThreadContexts[0].Instructions); i++ {
		stepThread(0)
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

	for threadId, t := range gameState.level.ThreadContexts {

		if IsThreadFinished(threadId) {
			continue
		}
		var thread = gameState.level.ThreadContexts[threadId]
		var instructions = thread.Instructions
		var threadState = gameState.threadContexts[threadId]
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
	gameState.threadContexts[threadId].Expanded = true
}

func undo() {

	for e := undoHistory.Front(); e != nil; e = e.Next() {
		if e.Next() == nil {
			gameState.globalState = e.Value.(History).globalContext
			gameState.threadContexts = e.Value.(History).threadContext
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
	var threadState = gameState.threadContexts[thread]
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
	var threadState = gameState.threadContexts[threadId]
	var pc = threadState.ProgramCounter

	return pc >= maxInstructions
}

func IsLevelPristine() bool {
	//return undoHistory.length == 0;
	return true
}

//TODO 删掉的东西

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

func main() {
	//startLevel("L1")
	app := iris.New()

	app.Controller("/movies", new(MoviesController))

	app.Run(iris.Addr(":8080"))
}

// MoviesController is our /movies controller.
type MoviesController struct {
	// if you build with go1.8 you have to use the mvc package always,
	// otherwise
	// you can, optionally
	// use the type alias `iris.C`,
	// same for
	// context.Context -> iris.Context,
	// mvc.Result -> iris.Result,
	// mvc.Response -> iris.Response,
	// mvc.View -> iris.View
	mvc.C
}

// Get returns list of the movies
// Demo:
// curl -i http://localhost:8080/movies
func (c *MoviesController) Get() []Movie {
	return movies
}

// GetBy returns a movie
// Demo:
// curl -i http://localhost:8080/movies/1
func (c *MoviesController) GetBy(id int) Movie {
	return movies[id]
}

// PutBy updates a movie
// Demo:
// curl -i -X PUT -F "genre=Thriller" -F "poster=@/Users/kataras/Downloads/out.gif" http://localhost:8080/movies/1
func (c *MoviesController) PutBy(id int) Movie {
	// get the movie
	m := movies[id]

	// get the request data for poster and genre
	file, info, err := c.Ctx.FormFile("poster")
	if err != nil {
		c.Ctx.StatusCode(iris.StatusInternalServerError)
		return Movie{}
	}
	file.Close()            // we don't need the file
	poster := info.Filename // imagine that as the url of the uploaded file...
	genre := c.Ctx.FormValue("genre")

	// update the poster
	m.Poster = poster
	m.Genre = genre
	movies[id] = m

	return m
}

// DeleteBy deletes a movie
// Demo:
// curl -i -X DELETE -u admin:password http://localhost:8080/movies/1
func (c *MoviesController) DeleteBy(id int) iris.Map {
	// delete the entry from the movies slice
	deleted := movies[id].Name
	movies = append(movies[:id], movies[id+1:]...)
	// and return the deleted movie's name
	return iris.Map{"deleted": deleted}
}
