package main

import "execution"

type GameState struct {
	threadContexts []execution.ThreadContext
	level          Level
	globalState    execution.GlobalContext
}

func (g *GameState) ResetForLevel(level Level) {

}

func (g *GameState) GetProgramOfThread(threadId int) []execution.Instruction {
	return g.level.ThreadContexts[threadId].Instructions
}
