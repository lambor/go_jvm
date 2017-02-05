package constants

import (
	"ch05/rtda"
	"ch05/instructions/base"
)

// Do nothing
type NOP struct{
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
