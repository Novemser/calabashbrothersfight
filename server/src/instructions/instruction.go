package instructions

import "execution"

type Instruction struct {
	Code               string
	Description        string
	Name               string
	ExpandInstructions []Instruction
}

func moveToNextInstruction(tc *execution.ThreadContext) {
	if tc.Expanded {
		tc.ExpProgramCounter++

		expandInsLen := len(tc.Instructions[tc.ProgramCounter].ExpandInstructions)
		// End of expandInstruction
		if tc.ExpProgramCounter >= expandInsLen {
			tc.Expanded = false
			tc.ProgramCounter++
			tc.ExpProgramCounter = 0
		}
	} else {
		tc.ProgramCounter++
		tc.ExpProgramCounter = 0
	}
}

type CommentInstruction struct {
	Instruction
}

func (c *CommentInstruction) execute(gc *execution.GlobalContext, tc *execution.ThreadContext) {

}
