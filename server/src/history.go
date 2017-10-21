package main

import "execution"

type History struct {
	threadContext  []execution.ThreadContext
	globalContext execution.GlobalContext
}
