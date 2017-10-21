package execution

import (
	"testing"
	"fmt"
)

func TestNewThread(t *testing.T) {
	gc := NewGlobalContext(Pair{
		"a",
		GlobalStateType{value:0, name:"a"},
	})
	//gc.values["a"] = GlobalStateType{value: 0, name: "a"}
	/**
	a = 0;
	if (a == 3) {
		a = 10;
	}
	print(a);
	 */
	assign := NewAssignmentStatment("a", NewLiteralExpression(3))
	assignIn := NewAssignmentStatment("a", NewLiteralExpression(10))
	ifUp := NewStartIfStatement(
		NewEqualityExpression(NewVariableExpression("a"), NewLiteralExpression(3)),
			"if1")
	ifEnd := NewEndIfStatement("if1")

	context := NewThreadContext(0, 0, 0, []Instruction{assign, ifUp, assignIn, ifEnd})
	for i := 0; i < 3; {
		ins := context.Instructions[i]
		ins.Execute(gc, context)
		i = context.ProgramCounter
	}
	fmt.Println(gc.values["a"])
}

func TestLevel(t *testing.T) {

}