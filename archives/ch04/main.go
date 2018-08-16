package main

import "fmt"
import "github.com/lojian/jvmgo/ch04/rtda"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushLong(2997924580)
	println(ops.PopInt())
	println(ops.PopLong())
}
