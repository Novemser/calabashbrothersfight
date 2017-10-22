package content

import (
	e "execution"
)

type Level struct {
	Label            string
	Title            string
	Description      string
	VictoryCondition string
	ThreadContexts   []*e.ThreadContext
	GlobalContext    *e.GlobalContext
}

func GetLevel(id int) *Level {
	switch id {
	case 0:
		return &Level{
			"初级教程1",
			"你渴望力量吗",
			"这是一个能模拟多线程并发执行的在线教程，在我们的故事中，变身葫芦娃，拒绝平庸，挑战极限。相信你一定能通过考验。",
			"同时让所有协程进入临界区",
			[]*e.ThreadContext{
				e.DefaultThreadContext(
					0, &[]e.Instruction{
						e.NewCommentInstruction("下面这一行是与多线程无关的业务代码"),
						e.NewDummyInstruction("doDummyThings"),
						e.NewCommentInstruction("下面是\"临界区\"代码块,多个线程不能同时访问"),
						e.NewCommentInstruction("你的目标是让所有的代码都一起进入临界区"),
						e.NewCriticalSectionExpression(),
						e.NewCommentInstruction("又是一行没用的业务代码"),
						e.NewDummyInstruction("doDummyThings"),
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
			"同时让所有协程进入临界区",
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
			"让协程-0执行panic语句",
			[]*e.ThreadContext{
				e.DefaultThreadContext(
					0, &[]e.Instruction{
						e.NewForStartIns(e.NewEqualityExpression(e.NewLiteralExpression(1), e.NewLiteralExpression(1)), "for"),
						e.NewMutexLockIns("mutex"),
						e.NewAssignmentInstruction("i", e.NewAdditionExpression(e.NewVariableExpression("i"), e.NewLiteralExpression(2))),
						e.NewStartIfStatement(
							e.NewEqualityExpression(e.NewVariableExpression("i"), e.NewLiteralExpression(5)), "if",
						),
						e.NewPanicIns("exitHere"),
						e.NewEndIfStatement("if"),
						e.NewMutexUnLockIns("mutex"),
						e.NewEndForIns("for"),
					},
				),
				e.DefaultThreadContext(
					1, &[]e.Instruction{
						e.NewForStartIns(e.NewEqualityExpression(e.NewLiteralExpression(1), e.NewLiteralExpression(1)), "for"),
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
				Key: "i", Value: e.GlobalStateType{Name: "i", Value: 0},
			}),
		}
	case 3:
		return &Level{
			"初级教程4",
			"武功再高也怕菜刀",
			"葫芦兄弟正在和白骨精正面战斗，老爷爷这时却遇到了麻烦，蛇精化身为3头巨蛇怪来吃他，没有葫芦兄弟那样本领的爷爷只好手握菜刀，伺机而动。他发现这三头巨蛇虽然强大，但是并不灵活，爷爷想到用死锁的方法，然这3个头互相攻击，循环等待，击败巨蛇。",
			"让三个协程发生死锁",
			[]*e.ThreadContext{
				e.DefaultThreadContext(0, &[]e.Instruction{
					e.NewForStartIns(e.NewEqualityExpression(e.NewLiteralExpression(1), e.NewLiteralExpression(1)), "for"),
					e.NewMutexLockIns("MutexA"),
					e.NewMutexLockIns("MutexB"),
					e.NewMutexUnLockIns("MutexA"),
					e.NewMutexUnLockIns("MutexB"),
					e.NewEndForIns("for"),
				}),
				e.DefaultThreadContext(1, &[]e.Instruction{
					e.NewForStartIns(e.NewEqualityExpression(e.NewLiteralExpression(1), e.NewLiteralExpression(1)), "for"),
					e.NewMutexLockIns("MutexB"),
					e.NewMutexLockIns("MutexC"),
					e.NewMutexUnLockIns("MutexB"),
					e.NewMutexUnLockIns("MutexC"),
					e.NewEndForIns("for"),
				}),
				e.DefaultThreadContext(2, &[]e.Instruction{
					e.NewForStartIns(e.NewEqualityExpression(e.NewLiteralExpression(1), e.NewLiteralExpression(1)), "for"),
					e.NewMutexLockIns("MutexC"),
					e.NewMutexLockIns("MutexA"),
					e.NewMutexUnLockIns("MutexA"),
					e.NewMutexUnLockIns("MutexC"),
					e.NewEndForIns("for"),
				}),
			},
			e.NewGlobalContext(e.Pair{
				Key: "MutexA", Value: e.GlobalStateType{Name: "MutexA", Value: &e.Lock{LastLockedThreadID: -1, LockCount: 0}},
			}, e.Pair{
				Key: "MutexB", Value: e.GlobalStateType{Name: "MutexB", Value: &e.Lock{LastLockedThreadID: -1, LockCount: 0}},
			}, e.Pair{
				Key: "MutexC", Value: e.GlobalStateType{Name: "MutexC", Value: &e.Lock{LastLockedThreadID: -1, LockCount: 0}},
			}),
		}
	case 4:
		return &Level{
			"初级教程4",
			"穿越时空去见你",
			"经过激烈的交战，葫芦娃和白骨精两败俱伤。就在千钧一刻，老爷爷发现了一个时空隧道，可以通过这个隧道在交战双方的场地来回。机智的老爷爷钻入了隧道，悄悄地来到了已经虚弱的白骨精背后。<br/>至于老爷爷是否能够成功袭击白骨精，就看你了！",
			"同时让所有协程进入临界区，注意channel的执行顺序",
			[]*e.ThreadContext{
				e.DefaultThreadContext(0, &[]e.Instruction{
					e.NewCommentInstruction("初始化了一个channel"),
					e.NewCommentInstruction("messages := make(chan string)"),
					e.NewDummyInstruction("otherLogic"),
					e.NewCommentInstruction("从channel读取一个数据"),
					e.NewChanReadIns("messages"),
					e.NewDummyInstruction("otherLogic"),
					e.NewCriticalSectionExpression(),
				}),
				e.DefaultThreadContext(1, &[]e.Instruction{
					e.NewCommentInstruction("初始化了一个channel"),
					e.NewCommentInstruction("messages := make(chan string)"),
					e.NewDummyInstruction("otherLogic"),
					e.NewCommentInstruction("向channel写入一个数据"),
					e.NewCommentInstruction("若无goroutine请求此数据会被阻塞"),
					e.NewChanWriteIns("messages", "Hello Golang!"),
					e.NewDummyInstruction("otherLogic"),
					e.NewCriticalSectionExpression(),
				}),
			},
			e.NewGlobalContext(e.Pair{
				Key: "messages", Value: e.GlobalStateType{Name: "messages", Value: &e.Channel{ReaderReady: false, WriterReady: false, DataReady: false}},
			}),
		}
	default:
		return nil
	}
}
