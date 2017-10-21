package main

import (
	"fmt"
	"container/list"
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

		//TODO
		if currentInstruction.isCriticalSection {
			howManyCriticalSections++
		}

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
//func undo() {
//	var last = JSON.parse(undoHistory.pop())
//	gameState.threadState = last.threadState
//	gameState.globalState = last.globalState
//}

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

func main() {

	fmt.Println(gameState)
	fmt.Printf("Go")
}
