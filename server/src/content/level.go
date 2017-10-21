package content

import (
	e "execution"
)

type Level struct {
	Label          string
	Title          string
	Description    string
	ThreadContexts []*e.ThreadContext
	GlobalContext  *e.GlobalContext
}

func GetLevel(id int) *Level {
	switch id {
	case 0:
		return &Level{
			"初级教程1",
			"你渴望力量吗",
			"这是一个能模拟多线程并发执行的在线教程～",
			[]*e.ThreadContext{
				e.DefaultThreadContext(
					0, &[]e.Instruction{
						e.NewCommentInstruction("下面这一行是与多线程无关的业务代码"),
						e.NewDummyInstruction("do_dummy_things"),
						e.NewCommentInstruction("下面是\"临界区\"代码块,多个线程不能同时访问"),
						e.NewCommentInstruction("你的目标是让所有的代码都一起进入临界区"),
						e.NewCriticalSectionExpression(),
						e.NewCommentInstruction("又是一行没用的业务代码"),
						e.NewDummyInstruction("do_dummy_things"),
					},
				),
				e.DefaultThreadContext(
					1, &[]e.Instruction{
						e.NewWhileStartIns(e.NewEqualityExpression(e.NewLiteralExpression(1), e.NewLiteralExpression(1)), "while"),
						e.NewAssignmentInstruction("a", e.NewLiteralExpression(23)),
						e.NewCriticalSectionExpression(),
						e.NewEndWhileIns("while"),
					},
				),
			},
			e.NewGlobalContext(e.Pair{
				Key: "mutex", Value: e.GlobalStateType{Name: "mutex", Value: e.Lock{LastLockedThreadID: -1, LockCount: 0}},
			}),
		}
	case 1:
		return &Level{
			"初级教程2",
			"魔鬼般的赋值语句",
			"欢迎来到葫芦娃的王国。在这里,7只英勇的葫芦娃将与诡计多端的白骨精展开一场计算机科学的较量……",
			[]*e.ThreadContext{
				e.DefaultThreadContext(
					0, &[]e.Instruction{
						e.NewAssignmentInstruction("a",
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
					1, &[]e.Instruction{
						e.NewAssignmentInstruction("a",
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
				Key: "a", Value: e.GlobalStateType{"a", "a", 0},
			}),
		}
	default:
		return nil
	}
}

var Level1 = GetLevel(1)
