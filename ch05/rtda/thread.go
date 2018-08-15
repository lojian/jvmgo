package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{stack: newStack(1024)}
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(self, maxLocals, maxStack)
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}
func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) ThrowDivByZero() {
	panic("ThrowDivByZero")
}

func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) TopFrameN(n uint) *Frame {
	return self.stack.topN(n)
}
