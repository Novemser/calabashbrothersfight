package execution

import "instructions"

type thread_context struct{
	pc int32 // program context
	instructions []instructions.Instruction
}
