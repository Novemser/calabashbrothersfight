package execution

type Instruction interface {
	GetCode() string
	GetDescription() string
	GetName() string
	GetExpandInstructions() []Instruction
}

type baseInstruction struct {
	Code               string
	Description        string
	Name               string
	ExpandInstructions []Instruction
}

func (i *baseInstruction) GetCode() string {
	return i.Code
}

func (i *baseInstruction) GetDescription() string {
	return i.Description
}

func (i *baseInstruction) GetName() string {
	return i.Name
}

func (i *baseInstruction) GetExpandInstructions() []Instruction {
	return i.ExpandInstructions
}

func moveToNextInstruction(tc *ThreadContext) {
	if tc.Expanded {
		tc.ExpProgramCounter++

		expandInsLen := len(tc.Instructions[tc.ProgramCounter].GetExpandInstructions())
		// End of expandInstruction
		if tc.ExpProgramCounter >= expandInsLen {
			tc.Expanded = false
			tc.ProgramCounter++
			tc.ExpProgramCounter = 0
		}
	} else {
		tc.ProgramCounter++
		tc.ExpProgramCounter = 0
	}
}

type CommentInstruction struct {
	baseInstruction
}

type DummyInstruction struct {
	baseInstruction
}

func (c *CommentInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	moveToNextInstruction(tc)
}

func (d *DummyInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	moveToNextInstruction(tc)
}
