package extended

import (
	"ch06/rtda"
	"ch06/instructions/base"
)

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
