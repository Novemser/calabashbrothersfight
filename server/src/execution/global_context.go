package execution

type GlobalContext struct {
	values map[string]GlobalStateType
}

type Pair struct {
	key string
	value GlobalStateType
}

func NewGlobalContext(args...Pair) *GlobalContext {
	mapCtx := make(map[string]GlobalStateType)
	for _, pair := range args {
		mapCtx[pair.key] = pair.value
	}

	return &GlobalContext{
		values: mapCtx,
	}
}


