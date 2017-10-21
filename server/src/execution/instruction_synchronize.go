package execution

import "fmt"

type Lock struct {
	LastLockedThreadID int
	LockCount          int
}

type MutexLockInstruction struct {
	baseInstruction
	lockName string
}

type MutexUnlockInstruction struct {
	baseInstruction
	lockName string
}

func (i *MutexLockInstruction) IsBlocking(gc *GlobalContext, tc *ThreadContext) bool {
	lockObj := gc.Values[i.lockName].Value.(Lock)
	if lockObj.LastLockedThreadID != -1 &&
		lockObj.LastLockedThreadID != tc.Id {
		// Waiting for other thread
		gc.LockMsg = "Thread " + fmt.Sprint(tc.Id) + " is waiting for thread " + fmt.Sprint(lockObj.LastLockedThreadID)
		return true
	} else {
		gc.LockMsg = "Thread " + fmt.Sprint(tc.Id) + " occupied lock " + i.lockName
		return false
	}
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
	if lockObj.LastLockedThreadID != -1 && lockObj.LastLockedThreadID != tc.Id {
		panic("WTF R U doing???")
	} else {
		lockObj.LastLockedThreadID = tc.Id
		if lockObj.LockCount >= 0 {
			lockObj.LockCount++
		} else {
			lockObj.LockCount = 1
		}
		moveToNextInstruction(tc)
	}
}

func (i *MutexUnlockInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	lockObj := gc.Values[i.lockName].Value.(Lock)
	if lockObj.LastLockedThreadID == tc.Id {
		lockObj.LockCount--
		if lockObj.LockCount <= 0 {
			lockObj.LastLockedThreadID = -1
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
