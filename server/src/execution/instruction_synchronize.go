package execution

type Lock struct {
	lastLockedThreadID int
	lockCount          int
}

type MutexLockInstruction struct {
	baseInstruction
	lockName string
}

type MutexUnlockInstruction struct {
	baseInstruction
	lockName string
}

//func NewMutexLockIns(name string) *MutexLockInstruction {
//	return &MutexLockInstruction{
//		baseInstruction{
//			Code: InstructionExpr()
//		}
//	}
//}

//func NewMutexUnLockIns() *MutexUnlockInstruction {
//
//}

func (i *MutexLockInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {

}

//func MemberCall(memberName string) string {
//}


func MethodCall(methodName string, args...string) string {
	argList := ""
	for i, str := range args {
		if i != len(args) - 1 {
			argList += str + ","
		} else {
			argList += str
		}
	}

	return methodName + "(" + argList + ")"
}

func LockMutex() string {
	return "Mutex.Lock"
}