package execution

type GlobalContext struct {
	values map[string]GlobalStateType
}

func NewGlobalContext() *GlobalContext {
	return &GlobalContext{
		values: make(map[string]GlobalStateType),
	}
}
