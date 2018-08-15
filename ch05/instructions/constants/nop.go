package constants

import "github.com/lojian/jvmgo/ch05/instructions/base"
import "github.com/lojian/jvmgo/ch05/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
