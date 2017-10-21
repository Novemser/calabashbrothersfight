package execution

import "instructions"

type ThreadContext struct{
	ProgramCounter int // program counter
	ExpProgramCounter int // expanded instruction program counter
	Instructions []instructions.Instruction
	Expanded bool
}
