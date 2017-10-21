package main

import "instructions"

type GameState struct {
	threadState int
	level       Level
	globalState int
}

func (g *GameState) ResetForLevel(level Level) {

}

func (g *GameState) GetProgramOfThread(threadId int) []instructions.Instruction {
	return g.level.threads.Instructions
}
