package main

import (
	"brookshear-vm/vm"
	"fmt"
	"flag"
	"strings"
	"path/filepath"
	"brookshear-vm/compiler"
	"brookshear-vm/file_parser"
)

func main() {
	var filePath = flag.String("f", "", "(REQUIRED) Input file. Ignores blank lines or lines starting with '//' or '#'")

	var mode = flag.String("a", "", `(REQUIRED) Action to perform, one of:
compile - compiles an input asm file to binary
execute - executes an input asm file`)
	//TODO: annotate - annotates an input asm file with byte comments
	//TODO: execute_b - executes an input binary file

	var verboseLvl = flag.Int("v", 0, `Set level of verbosity, one of:
0 - print memory dump on end
1 - print each executed instruction
2 - print data changes on each executed instruction`)

	var origUsage = flag.Usage
	flag.Usage = func() {
		preHelp()
		origUsage()
		postHelp()
	}
	flag.Parse()

	if strings.TrimSpace(*filePath) == "" || strings.TrimSpace(*mode) == "" {
		flag.Usage()
		return
	}

	instrsStr, err := file_parser.ParseFile(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %+v\n", err)
		return
	}

	switch *mode {
	case "execute":
		runInstrs(instrsStr, *verboseLvl)
	case "compile":
		compile(instrsStr, *filePath)
	default:
		fmt.Println("Unknown option:", *mode)
		return
	}

}

func preHelp() {
	println("A brookshear virtual machine implementation capable of executing and compiling brookshear assembly instructions.")
	println()
}

func postHelp() {
}

func compile(instrsStr []string, filePath string) {
	if err := compiler.Compile(instrsStr, filepath.Clean(filePath)); err != nil {
		fmt.Printf("Error during compilation: %+v\n", err)
		return
	}
}
func runInstrs(instrsStr []string, verboseLvl int) {
	var vm = vm.New(verboseLvl)
	if err := vm.Run(instrsStr); err != nil {
		fmt.Printf("Error during execution: %+v\n", err)
		return
	}

	fmt.Println()
	fmt.Println("Memory dump (16x16):")
	vm.PrintMemory()
	fmt.Println()
	fmt.Println("Registers (16):")
	vm.PrintRegisters()
}
