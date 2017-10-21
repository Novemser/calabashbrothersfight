package execution

import (
	"testing"
)

func TestNewThread(t *testing.T) {
	assign := NewAssignmentStatment("a", NewLiteralExpression(2))
	ifUp := NewStartIfStatement(
		NewEqualityExpression(NewVariableExpression("a"),
			NewLiteralExpression(2)), "if1")

	context := NewThreadContext(0, 0, 0, []Instruction{assign, ifUp})
	gc := NewGlobalContext()
	for i := 0; i < 2; i++ {
		ins := context.Instructions[i]
		ins.Execute(gc, &context)

	}
}
