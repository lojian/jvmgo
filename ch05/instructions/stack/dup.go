package stack

import "github.com/lojian/jvmgo/ch05/instructions/base"
import "github.com/lojian/jvmgo/ch05/rtda"

type DUP struct{ base.NoOperandsInstruction }

func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

type DUP_X1 struct{ base.NoOperandsInstruction }

func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

type DUP_X2 struct{ base.NoOperandsInstruction }

func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

type DUP2 struct{ base.NoOperandsInstruction }

func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

type DUP2_X1 struct{ base.NoOperandsInstruction }

func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

type DUP2_X2 struct{ base.NoOperandsInstruction }

func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	val4 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val4)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}
