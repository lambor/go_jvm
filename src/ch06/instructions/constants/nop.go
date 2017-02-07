package constants

import (
	"ch06/rtda"
	"ch06/instructions/base"
)

// Do nothing
type NOP struct{
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
