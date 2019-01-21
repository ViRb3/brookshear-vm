package cmd

import (
	"github.com/spf13/cobra"
	"brookshear-vm/pkgs/vm"
	"fmt"
	"brookshear-vm/pkgs/io"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an assembly or binary file",
	Example: `run -f sample.asm.txt -t asm
run -f sample.asm.bin -t bin -v 2`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	switch InputType {
	case "asm":
		runAsm()
	case "bin":
		runBin()
	default:
		fmt.Println("Invalid input type")
		return
	}
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().IntVarP(&Verbosity, "verbosity", "v", 0, `Level of verbosity. One of:
0 - print memory and register dump on completion (DEFAULT)
1 - print each executed instruction
2 - print data changes on each executed instruction`)
	runCmd.Flags().StringVarP(&Input, "file", "f", "", `Input file path. For assembly, supports comments: '//', '#' and ';'`)
	runCmd.Flags().StringVarP(&InputType, "input-type", "t", "", `Input file type. One of:
asm - assembly instructions
bin - binary bytecode`)

	runCmd.MarkFlagRequired("file")
	runCmd.MarkFlagRequired("input-type")
}

func runBin() {
	err := doRunBin()
	if err != nil {
		fmt.Printf("Error during binary execution:\n%+v\n", err)
		return
	}
}

func doRunBin() error {
	instrs, err := io.Decompile(Input)
	if err != nil {
		return err
	}
	return startVM(instrs)
}

func runAsm() {
	var err = doRunAsm()
	if err != nil {
		fmt.Printf("Error during assembly execution:\n%+v\n", err)
	}
}

func doRunAsm() error {
	instrStrs, err := io.ReadAsmFile(Input)
	if err != nil {
		return err
	}
	instrs, err := vm.ParseInstructions(instrStrs)
	if err != nil {
		return err
	}
	return startVM(instrs)
}

func startVM(instrs []*vm.Instruction) error {
	var vm = vm.New(Verbosity)
	if err := vm.Run(instrs); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Memory dump (16x16):")
	vm.PrintMemory()
	fmt.Println()
	fmt.Println("Registers (16):")
	vm.PrintRegisters()
	return nil
}


