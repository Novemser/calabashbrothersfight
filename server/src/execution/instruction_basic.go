package execution

import (
	"reflect"
	"fmt"
)

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

		expandInsLen := len((*tc.Instructions)[tc.ProgramCounter].GetExpandInstructions())
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
	expr Expression
}

type EndIfInstruction struct {
	baseInstruction
}

type WhileInstruction struct {
	baseInstruction
	expr Expression
}

type EndWhileInstruction struct {
	baseInstruction
}

type ExpandableInstruction struct {
	baseInstruction
}

type AtomicAssignmentToTemp struct {
	baseInstruction
	expr Expression
}

type AtomicAssignmentFromTemp struct {
	baseInstruction
	expr Expression
}

// If more than one thread entered critical section,
// The player win
type CriticalSectionExpression struct {
	baseInstruction
}

func NewWhileStartIns(exp Expression, name string) *WhileInstruction {
	return &WhileInstruction{
		baseInstruction{
			Code:        WhileStart() + AddBraces(exp) + Then(),
			Description: "While statement",
			Name:        name,
		}, exp,
	}
}

func NewEndWhileIns(exp Expression, name string) *EndWhileInstruction {
	return &EndWhileInstruction{
		baseInstruction{
			Code:        End(),
			Description: "End of while",
			Name:        name,
		},
	}
}

func NewCriticalSectionExpression() *CriticalSectionExpression {
	return &CriticalSectionExpression{
		baseInstruction{
			Code:        InstructionExpr("criticalSection()"),
			Name:        "Critical section",
			Description: "Critical section",
		},
	}
}

func NewExpandableInstruction(name string, expandIns []Instruction) *ExpandableInstruction {
	return &ExpandableInstruction{baseInstruction{
		Code:               name,
		Name:               "Expandable ins",
		ExpandInstructions: expandIns,
	}}
}

func NewAtomicAssignFromTemp(exprLeft Expression) *AtomicAssignmentFromTemp {
	base := baseInstruction{
		Code:        InstructionExpr(AssignmentExpr(exprLeft, NewVariableExpression("temp"))),
		Description: "Atomic assign temp to left expr",
	}
	return &AtomicAssignmentFromTemp{base, exprLeft}
}

func NewAtomicAssignToTemp(exprRight Expression) *AtomicAssignmentToTemp {
	base := baseInstruction{
		Code:        InstructionExpr(AssignmentExpr(NewVariableExpression("temp"), exprRight)),
		Description: "Atomic assign right expr to temp",
	}
	return &AtomicAssignmentToTemp{base, exprRight}
}

func NewAssignmentInstruction(varName string, exp Expression) *ExpandableInstruction {
	expInsList := []Instruction{
		NewAtomicAssignToTemp(exp),
		NewAtomicAssignFromTemp(NewVariableExpression(varName)),
	}

	expIns := NewExpandableInstruction(
		AssignmentExpr(NewVariableExpression(varName), exp), expInsList)

	return expIns
}

func NewStartIfStatement(exp Expression, name string) *IfInstruction {
	base := baseInstruction{
		Code:        IfStart() + AddBraces(exp) + Then(),
		Description: "If statement",
		Name:        name,
	}
	return &IfInstruction{base, exp}
}

func NewEndIfStatement(name string) *EndIfInstruction {
	base := baseInstruction{
		Code:        End(),
		Description: "End if statement",
		Name:        name,
	}
	return &EndIfInstruction{base}
}

func (e *CriticalSectionExpression) Execute(gc *GlobalContext, tc *ThreadContext) {
	moveToNextInstruction(tc)
}

func (e *ExpandableInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	if !tc.Expanded {
		tc.Expanded = true
		for _, ins := range e.GetExpandInstructions() {
			ins.Execute(gc, tc)
		}
	}
}

func (e *AtomicAssignmentFromTemp) Execute(gc *GlobalContext, tc *ThreadContext) {
	gc.Values[fmt.Sprint(e.expr.GetName())] =
		GlobalStateType{Value: tc.TempVariable, Name: e.expr.GetName()}
	moveToNextInstruction(tc)
}

func (e *AtomicAssignmentToTemp) Execute(gc *GlobalContext, tc *ThreadContext) {
	// TODO: Do we really need to evaluate?
	tc.TempVariable = e.expr.Evaluate(gc, tc)
	moveToNextInstruction(tc)
}

func (e *EndIfInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	moveToNextInstruction(tc)
}

func (c *CommentInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	moveToNextInstruction(tc)
}

func (d *DummyInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	moveToNextInstruction(tc)
}

func (i *WhileInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	if (i.expr.Evaluate(gc, tc)).(bool) {
		moveToNextInstruction(tc)
	} else {
		matchingEndWhile := findMatchingInsIndex(tc, i.GetName(), reflect.TypeOf(&EndWhileInstruction{}))
		goToInstruction(tc, matchingEndWhile+1)
	}
}

func (i *EndWhileInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	// We jump back to while start
	matchingWhile := findMatchingInsIndex(tc, i.GetName(), reflect.TypeOf(&WhileInstruction{}))
	goToInstruction(tc, matchingWhile)
}

func (i *IfInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	if (i.expr.Evaluate(gc, tc)).(bool) {
		moveToNextInstruction(tc)
	} else {
		matchingEndIf := findMatchingInsIndex(tc, i.GetName(), reflect.TypeOf(&EndIfInstruction{}))
		goToInstruction(tc, matchingEndIf)
	}
}

func goToInstruction(context *ThreadContext, num int) {
	context.ProgramCounter = num
	context.ExpProgramCounter = 0
}

func findMatchingInsIndex(context *ThreadContext, name string, tp reflect.Type) int {
	for i, ins := range *context.Instructions {
		if reflect.TypeOf(ins) == tp && ins.GetName() == name {
			fmt.Println(reflect.TypeOf(ins), tp)
			return i
			//fmt.Println(i)
		}
	}
	panic("没找到对应的结束语句, 挂了吧")
	return -1
}

func IfStart() string {
	return "if"
}

func WhileStart() string {
	return "while"
}

func AddBraces(expCode Expression) string {
	return " (" + expCode.GetCode() + ") "
}

func Then() string {
	return "{"
}

func End() string {
	return "}"
}

func InstructionExpr(code string) string {
	return code + ";"
}

func AssignmentExpr(left Expression, right Expression) string {
	return BinaryOperationCode(left, right, "=")
}
