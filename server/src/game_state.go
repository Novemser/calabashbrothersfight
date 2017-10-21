package main

import "execution"

type GameState struct {
	threadState int
	level       Level
	globalState int
}

func (g *GameState) ResetForLevel(level Level) {

}

func (g *GameState) GetProgramOfThread(threadId int) []execution.Instruction {
	return g.level.threads[threadId].Instructions
}
