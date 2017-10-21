package main

import (
	"fmt"
	"container/list"
	//"github.com/kataras/iris"
	//"github.com/kataras/iris/middleware/logger"
	//"github.com/kataras/iris/middleware/recover"
	//"strconv"
	c "content"
)

var gameState = new(c.GameState)
var undoHistory = list.New()

func saveForUndo() {
	var history = c.History{gameState.ThreadContexts, gameState.GlobalState}
	undoHistory.PushBack(history)
}

//TODO
func startLevel(levelName string) {
	//清空
	for e := undoHistory.Front(); e != nil; e = e.Next() {
		undoHistory.Remove(e)
	}

	// 开始
	var level = c.Level1
	gameState.GlobalState = *level.GlobalContext
	gameState.ThreadContexts = level.ThreadContexts
	gameState.Level = *level

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

	for threadId, t := range gameState.Level.ThreadContexts {

		if IsThreadFinished(threadId) {
			continue
		}
		var thread = gameState.Level.ThreadContexts[threadId]
		var instructions = thread.Instructions
		var threadState = gameState.ThreadContexts[threadId]
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
	gameState.ThreadContexts[threadId].Expanded = true
}

func undo() {

	for e := undoHistory.Front(); e != nil; e = e.Next() {
		if e.Next() == nil {
			gameState.GlobalState = e.Value.(c.History).GlobalContext
			gameState.ThreadContexts = e.Value.(c.History).ThreadContext
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
	var threadState = gameState.ThreadContexts[thread]
	var pc = threadState.ProgramCounter

	if !IsThreadFinished(thread) {
		saveForUndo()
		if threadState.Expanded {
			//展开了的情况
			program[pc].GetExpandInstructions()[threadState.ExpProgramCounter].Execute(&gameState.GlobalState, &threadState)
		} else {
			program[pc].Execute(&gameState.GlobalState, &threadState)
		}
		checkForVictoryConditions()

	}
	fmt.Println(program, threadState, pc)

}

func IsThreadFinished(threadId int) bool {
	program := gameState.GetProgramOfThread(threadId)
	var maxInstructions = len(program)
	var threadState = gameState.ThreadContexts[threadId]
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
	GameStatus  int       `json:"gameStatus"`
}
type Program struct {
	Current          []int   `json:"current"`
	CanStepNext      bool    `json:"canStepNext"`
	CanCurrentExpand bool    `json:"canCurrentExpand"`
	Code             []Coder `json:"code"`
}
type Coder struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Indent int    `json:"indent"`
}
type BackResponse struct {
	Status int       `json:"status"`
	Msg    string    `json:"msg"`
	Data   LevelInfo `json:"data"`
}

var cs = []Coder{
	{
		Name:   "na",
		Code:   "11",
		Indent: 22,
	},
	{
		Name:   "na",
		Code:   "11",
		Indent: 22,
	},
}
var p = []Program{
	{
		CanStepNext:      true,
		Current:          []int{1, 2},
		CanCurrentExpand: false,
		Code:             cs,
	},
	{
		CanStepNext:      false,
		Current:          []int{1, 2},
		CanCurrentExpand: false,
		Code:             cs,
	},
}

var levelInfo = LevelInfo{Title: "a", Description: "d", Programs: p, GameStatus: 0}

func loadLevel(id int, err error) BackResponse {
	backResponse := BackResponse{}
	if err == nil {
		backResponse.Data = levelInfo
		backResponse.Status = 0
		backResponse.Msg = "success"
		//正确解析
	} else {
		backResponse.Data = levelInfo
		backResponse.Status = 1
		backResponse.Msg = "关卡ID无法识别"
		// 不是数字
	}
	return backResponse
}

func main() {
	startLevel("L1")
	//app := iris.New()
	//app.Use(recover.New())
	//app.Use(logger.New())
	//
	//// 加载关卡
	//app.Get("/api/level/{id}", func(ctx iris.Context) {
	//	levelIdStr := ctx.Params().Get("id")
	//	levelId, err := strconv.Atoi(levelIdStr)
	//	ctx.JSON(loadLevel(levelId, err))
	//})
	//
	//app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}
