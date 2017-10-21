package main

import "execution"

type History struct {
	threadState   []ThreadState
	globalContext execution.GlobalContext
}
