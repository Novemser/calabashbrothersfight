package main

import "fmt"

var gameState = new(GameState)
var undoHistory = []History{}

func saveForUndo() {
	var history = History{gameState.threadContext, gameState.globalState}
	undoHistory = append(undoHistory, history)

}

//TODO
func startLevel(levelName string) {
	undoHistory = []History{}
}

func stepThread(thread int) {
	if IsLevelPristine() {
		//第一步执行
		//sendEvent('Gameplay', 'level-first-step', gameState.getLevelId());
	}
	var program = gameState.GetProgramOfThread(thread)
	var threadState = gameState.threadContext[thread]
	var pc = threadState.ProgramCounter

	if IsThreadFinished(thread) {
		saveForUndo()
		if threadState.Expanded {
			//展开了的情况
			program[pc].GetExpandInstructions()[threadState.ExpProgramCounter].Execute()
		}
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

func main() {

	fmt.Println(gameState)
	fmt.Printf("Go")
}
