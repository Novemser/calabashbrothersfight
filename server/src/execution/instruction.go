package execution

type Instruction interface {
	GetCode() string
	GetDescription() string
	GetName() string
	GetExpandInstructions() []Instruction
	Execute(gc *GlobalContext, tc *ThreadContext)
	IsBlocking(gc *GlobalContext, tc *ThreadContext) bool
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

func (i *baseInstruction) IsBlocking(gc *GlobalContext, tc *ThreadContext) bool {
	return false
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

type IfInstruction struct {
	baseInstruction
	exp Expression
}

func NewStartIfStatment(exp Expression, name string) IfInstruction {
	base := baseInstruction{
		Code:        IfStart() + AddBraces(exp) + Then(),
		Description: "If statement",
		Name:        name,
	}
	return IfInstruction{base, exp}
}

func (c *CommentInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	moveToNextInstruction(tc)
}

func (d *DummyInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	moveToNextInstruction(tc)
}

func (i *IfInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	if bool(i.exp.Evaluate(gc, tc)) {
		moveToNextInstruction(tc)
	} else {

	}
}

func IfStart() string {
	return "if"
}

func AddBraces(expCode Expression) string {
	return " (" + expCode.GetCode() + ") "
}

func Then() string {
	return "{"
}
