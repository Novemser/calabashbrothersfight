package main

import "fmt"

var gameState = new(GameState)

func saveForUndo() {
	var history = History{gameState.threadStates,}
}
//TODO
func stepThread(thread int) {
	if IsLevelPristine() {
		//第一步执行
	}
	var program = gameState.GetProgramOfThread(thread)
	var threadState = gameState.threadStates[thread]
	var pc = threadState.tc.ProgramCounter

	if IsThreadFinished(thread){
		
	}
	fmt.Println(program, threadState, pc)


}

func IsThreadFinished(threadId int) bool {
	program := gameState.GetProgramOfThread(threadId)
	var maxInstructions = len(program)
	var threadState = gameState.threadStates[threadId]
	var pc = threadState.tc.ProgramCounter

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
