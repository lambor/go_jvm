package rtda
type Frame struct {
	lower *Frame
	localVars LocalVars
	operandStack *OperandStack
	thread *Thread
	nextPC int
	//method *heap.Method
}

func newFrame(thread *Thread,maxLocals,maxStack uint) *Frame {
	return &Frame {
		localVars:	newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
		thread: thread,
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) NextPC() int {
	return self.nextPC
}
