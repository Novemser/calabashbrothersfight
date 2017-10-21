package main

import "execution"

type GameState struct {
	threadStates []ThreadState
	level       Level
	globalState execution.GlobalContext
}

func (g *GameState) ResetForLevel(level Level) {

}

func (g *GameState) GetProgramOfThread(threadId int) []execution.Instruction {
	return g.level.threads[threadId].Instructions
}
