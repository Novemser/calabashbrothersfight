package main

import "execution"

type ThreadState struct{
	id int
	tc execution.ThreadContext
}
