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
			"这是一个能模拟多线程并发执行的在线教程，在我们的故事中，变身葫芦娃，拒绝平庸，挑战极限。相信你一定能通过考验。",
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
						e.NewForStartIns(e.NewEqualityExpression(e.NewLiteralExpression(1), e.NewLiteralExpression(1)), "for"),
						e.NewAssignmentInstruction("a", e.NewLiteralExpression(23)),
						e.NewCriticalSectionExpression(),
						e.NewEndForIns("for"),
					},
				),
			},
			e.NewGlobalContext(e.Pair{
				Key: "mutex", Value: e.GlobalStateType{Name: "mutex", Value: &e.Lock{LastLockedThreadID: -1, LockCount: 0}},
			}),
		}
	case 1:
		return &Level{
			"初级教程2",
			"魔鬼般的赋值语句",
			"欢迎来到葫芦娃的王国。在这里,7只英勇的葫芦娃将与诡计多端的白骨精展开一场计算机科学的较量。并且，似乎还隐藏着更多的恶势力……",
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
	case 2:
		return &Level{
			"初级教程3",
			"装甲再厚也能打破",
			"葫芦兄弟的火娃被白骨精打伤，落单了，这时候蜘蛛精穿着她的蛛丝做的新盔甲要抓他，尽管她的盔甲用上了互斥锁，但是还是有弱点，抓住时机，利用两个线程，一举击败她。",
			[]*e.ThreadContext{
				e.DefaultThreadContext(
					0, &[]e.Instruction{
						e.NewForStartIns(e.NewEqualityExpression(e.NewLiteralExpression(1), e.NewLiteralExpression(1)), "for"),
						e.NewMutexLockIns("mutex"),
						e.NewAssignmentInstruction("i", e.NewAdditionExpression(e.NewVariableExpression("i"), e.NewLiteralExpression(2))),
						e.NewStartIfStatement(
							e.NewEqualityExpression(e.NewVariableExpression("i"), e.NewLiteralExpression(5)), "if",
						),
						e.NewCriticalSectionExpression(),
						e.NewEndIfStatement("if"),
						e.NewMutexUnLockIns("mutex"),
						e.NewEndForIns("for"),
					},
				),
				e.DefaultThreadContext(
					0, &[]e.Instruction{
						e.NewForStartIns(e.NewEqualityExpression(e.NewLiteralExpression(1), e.NewLiteralExpression(1)), "for"),
						e.NewCriticalSectionExpression(),
						e.NewMutexLockIns("mutex"),
						e.NewAssignmentInstruction("i", e.NewAdditionExpression(e.NewVariableExpression("i"), e.NewLiteralExpression(-1))),
						e.NewMutexUnLockIns("mutex"),
						e.NewEndForIns("for"),
					},
				),
			},
			e.NewGlobalContext(e.Pair{
				Key: "mutex", Value: e.GlobalStateType{Name: "mutex", Value: &e.Lock{LastLockedThreadID: -1, LockCount: 0}},
			}, e.Pair{
				Key: "i", Value: e.GlobalStateType{Name: "i", Value: 1},
			}),
		}
	//case 3:
		//return &Level{
		//	"",
		//	"",
		//	"",
		//
		//}
	default:
		return nil
	}
}
