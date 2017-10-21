package content

import "execution"

type GameState struct {
	ThreadContexts []*execution.ThreadContext
	Level          Level
	GlobalState    *execution.GlobalContext
}

func (g *GameState) ResetForLevel(level Level) {

}

func (g *GameState) GetProgramOfThread(threadId int) []execution.Instruction {
	return *g.Level.ThreadContexts[threadId].Instructions
}
