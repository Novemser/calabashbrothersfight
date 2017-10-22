package execution

import "fmt"

type Lock struct {
	LastLockedThreadID int
	LockCount          int
}

type Channel struct {
	ReaderReady bool
	WriterReady bool
	DataReady   bool
}

type MutexLockInstruction struct {
	baseInstruction
	lockName string
}

type MutexUnlockInstruction struct {
	baseInstruction
	lockName string
}

type ChannelReadInstruction struct {
	baseInstruction
	chanName string
}

type ChannelWriteInstruction struct {
	baseInstruction
	chanName string
}

func (i *ChannelWriteInstruction) IsBlocking(gc *GlobalContext, tc *ThreadContext) bool {
	chanObj := gc.Values[i.chanName].Value.(*Channel)

	if chanObj.ReaderReady {
		return false
	} else {
		// If blocked, writer is ready.
		chanObj.WriterReady = true
		return true
	}
}

func (i *ChannelReadInstruction) IsBlocking(gc *GlobalContext, tc *ThreadContext) bool {
	chanObj := gc.Values[i.chanName].Value.(*Channel)

	if !chanObj.DataReady {
		chanObj.ReaderReady = true
		return true
	} else {
		return false
	}
}

func (i *MutexLockInstruction) IsBlocking(gc *GlobalContext, tc *ThreadContext) bool {
	lockObj := gc.Values[i.lockName].Value.(*Lock)
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

func NewChanReadIns(name string) *ChannelReadInstruction {
	return &ChannelReadInstruction{
		baseInstruction{
			Code:        "<-\"" + name + "\"",
			Description: "Channel read",
			Name:        "chan read",
		}, name,
	}
}

func NewChanWriteIns(name string, value string) *ChannelWriteInstruction {
	return &ChannelWriteInstruction{
		baseInstruction{
			Code:        name + " <- \"" + value + "\"",
			Description: "Channel write",
			Name:        "chan write",
		}, name,
	}
}

func NewMutexLockIns(name string) *MutexLockInstruction {
	return &MutexLockInstruction{
		baseInstruction{
			Code: InstructionExpr(
				MemberCall(name, MethodCall("Lock")),
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
				MemberCall(name, MethodCall("UnLock")),
			),
			Name:        name,
			Description: "A mutex unlock",
		}, name,
	}
}

func (i *ChannelWriteInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	chanObj := gc.Values[i.chanName].Value.(*Channel)
	chanObj.DataReady = true
	moveToNextInstruction(tc)
}

func (i *ChannelReadInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	//chanObj := gc.Values[i.chanName].Value.(*Channel)
	moveToNextInstruction(tc)
}

func (i *MutexLockInstruction) Execute(gc *GlobalContext, tc *ThreadContext) {
	lockObj := gc.Values[i.lockName].Value.(*Lock)
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
	lockObj := gc.Values[i.lockName].Value.(*Lock)
	if lockObj.LastLockedThreadID == tc.Id {
		lockObj.LockCount--
		if lockObj.LockCount <= 0 {
			lockObj.LastLockedThreadID = -1
		}
		moveToNextInstruction(tc)
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
