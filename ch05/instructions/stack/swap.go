package stack

import "github.com/lojian/jvmgo/ch05/instructions/base"
import "github.com/lojian/jvmgo/ch05/rtda"

type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
