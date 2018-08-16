package interpreter

import (
	"fmt"

	"github.com/lojian/jvmgo/ch05/instructions/base"

	"github.com/lojian/jvmgo/ch05/classfile"

	"github.com/lojian/jvmgo/ch05/instructions"
	"github.com/lojian/jvmgo/ch05/rtda"
)

func Interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		fmt.Printf("pc: %2d  inst: %T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
