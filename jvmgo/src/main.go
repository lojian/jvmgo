package main

import (
	"classfile"
	"classpath"
	"flag"
	"fmt"
	"interpreter"
	"os"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.xjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpreter.Interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s \n", cmd.class)
	}

}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		fmt.Printf("found method:%v,  %v\n", m.Name(), m.Descriptor())

		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	xjreOption  string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help messages")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help messages")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.xjreOption, "Xjre", "", "path to JRE")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
