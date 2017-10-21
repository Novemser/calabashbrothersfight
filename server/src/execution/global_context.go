package execution

type GlobalContext struct {
	Values map[string]GlobalStateType
	LockMsg string
	IsPanic bool
}

type Pair struct {
	Key   string
	Value GlobalStateType
}

func NewGlobalContext(args...Pair) *GlobalContext {
	mapCtx := make(map[string]GlobalStateType)
	for _, pair := range args {
		mapCtx[pair.Key] = pair.Value
	}

	return &GlobalContext{
		Values: mapCtx,
		LockMsg:"",
		IsPanic:false,
	}
}


