package execution

type ThreadContext struct {
	id                int
	ProgramCounter    int // program counter
	ExpProgramCounter int // expanded instruction program counter
	Instructions      []Instruction
	Expanded          bool
	TempVariable      interface{} // Temp variable to store Value
}

func DefaultThreadContext(id int, insList []Instruction) ThreadContext {
	return *NewThreadContext(id, 0, 0, insList)
}

func NewThreadContext(id, pc, epc int, insList []Instruction) *ThreadContext {
	return &ThreadContext{
		id, pc, epc, insList, false, nil,
	}
}
