package execution

type GlobalContext struct {
	values map[string]GlobalStateType
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
		values: mapCtx,
	}
}


