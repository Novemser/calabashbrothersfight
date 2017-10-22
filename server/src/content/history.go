package content

import "execution"

type History struct {
	ThreadContext []execution.ThreadContext
	GlobalContext execution.GlobalContext
}
