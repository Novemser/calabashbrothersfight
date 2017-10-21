package execution

type Thread struct {
	Context ThreadContext
	Name    string
}

func NewThread(context ThreadContext, name string) Thread {
	return Thread{Context: context, Name: name}
}
