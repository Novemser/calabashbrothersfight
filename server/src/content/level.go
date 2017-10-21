package content

import (
	e "execution"
)

type Level struct {
	Label          string
	Title          string
	Description    string
	ThreadContexts []e.ThreadContext
	GlobalContext  *e.GlobalContext
}

var Level1 = &Level{
	"教程",
	"魔鬼般的赋值语句",
	"欢迎来到葫芦娃的王国。在这里,7只英勇的葫芦娃将与诡计多端的白骨精展开一场计算机科学的较量……",
	[]e.ThreadContext{
		e.DefaultThreadContext(
			0, []e.Instruction{
				e.NewAssignmentStatment("a",
					e.NewAdditionExpression(
						e.NewVariableExpression("a"),
						e.NewLiteralExpression(1),
					),
				),
				e.NewStartIfStatement(
					e.NewEqualityExpression(
						e.NewVariableExpression("a"),
						e.NewLiteralExpression(1),
					), "if",
				),
				e.NewCriticalSectionExpression(),
				e.NewEndIfStatement("if"),
			},
		),
		e.DefaultThreadContext(
			1, []e.Instruction{
				e.NewAssignmentStatment("a",
					e.NewAdditionExpression(
						e.NewVariableExpression("a"),
						e.NewLiteralExpression(1),
					),
				),
				e.NewStartIfStatement(
					e.NewEqualityExpression(
						e.NewVariableExpression("a"),
						e.NewLiteralExpression(1),
					), "if",
				),
				e.NewCriticalSectionExpression(),
				e.NewEndIfStatement("if"),
			},
		),
	},
	e.NewGlobalContext(e.Pair{
		Key:"a", Value:e.GlobalStateType{"a", "a", 0},
	}),
}
