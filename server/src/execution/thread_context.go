package execution

type ThreadContext struct{
	ProgramCounter int // program counter
	ExpProgramCounter int // expanded instruction program counter
	Instructions []Instruction
	Expanded bool
}
