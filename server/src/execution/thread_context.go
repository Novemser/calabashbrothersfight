package execution

type ThreadContext struct{
	id int
	ProgramCounter int // program counter
	ExpProgramCounter int // expanded instruction program counter
	Instructions []Instruction
	Expanded bool
	TempVariable interface{} // Temp variable to store value
}
