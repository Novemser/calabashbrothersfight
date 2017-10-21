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

func NewMutexLockIns(name string) *MutexLockInstruction {
	return &MutexLockInstruction{
		baseInstruction{
			Code: InstructionExpr(
				MemberCall("Mutex", MethodCall("Lock")),
			),
			Name:        name,
			Description: "A mutex lock",
		}, name,
	}
}

func NewMutexUnLockIns(name string) *MutexUnlockInstruction {
	return &MutexUnlockInstruction{
		baseInstruction{
			Code: InstructionExpr(
				MemberCall("Mutex", MethodCall("UnLock")),
			),
			Name:        name,
			Description: "A mutex unlock",
		}, name,
	}
}

func (i *MutexLockInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	lockObj := gc.Values[i.lockName].Value.(Lock)
	if lockObj.lastLockedThreadID != -1 && lockObj.lastLockedThreadID != tc.Id {
		panic("WTF R U doing???")
	} else {
		lockObj.lastLockedThreadID = tc.Id
		if lockObj.lockCount >= 0 {
			lockObj.lockCount++
		} else {
			lockObj.lockCount = 1
		}
		moveToNextInstruction(tc)
	}
}

func (i *MutexUnlockInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	lockObj := gc.Values[i.lockName].Value.(Lock)
	if lockObj.lastLockedThreadID == tc.Id {
		lockObj.lockCount--
		if lockObj.lockCount <= 0 {
			lockObj.lastLockedThreadID = -1
		}
	} else {
		panic("WTF R UU doing??")
	}

}

func MemberCall(memberName string, methodCall string) string {
	return memberName + "." + methodCall
}

func MethodCall(methodName string, args ...string) string {
	argList := ""
	for i, str := range args {
		if i != len(args)-1 {
			argList += str + ","
		} else {
			argList += str
		}
	}

	return methodName + "(" + argList + ")"
}
