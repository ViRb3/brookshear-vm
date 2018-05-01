package vm

import (
	"fmt"
	"github.com/pkg/errors"
)

type VM struct {
	memory     []byte
	registers  []byte
	pc         byte
	halt       bool
	verboseLvl int
}

func New(verboseLvl int) VM {
	var vm = VM{}
	vm.memory = make([]byte, 256)
	vm.registers = make([]byte, 16)
	vm.verboseLvl = verboseLvl
	return vm
}

func ParseInstructions(instrStrs []string) ([]Instruction, error) {
	var instrs []Instruction

	for i, instrStr := range instrStrs {
		instr, err := Parse(instrStr)
		if err != nil {
			err = errors.New(fmt.Sprintf("%s\nInstruction: %s\nLine: %d", err.Error(), instrStr, i))
			return instrs, err
		}
		instrs = append(instrs, instr)
	}
	return instrs, nil
}

func (vm *VM) Run(instrStrs []string) error {
	instrs, err := ParseInstructions(instrStrs)
	if err != nil {
		return err
	}

	vm.loadInstructionsToMemory(instrs)

	if err = vm.instrLoop(); err != nil {
		return err
	}
	return nil
}

func (vm *VM) Execute(instr Instruction) {
	instr.Execute(vm)
}

func (vm *VM) instrLoop() error {
	for !vm.halt {
		var nextBytes = vm.memory[vm.pc : vm.pc+4]
		var instr, err = ParseFromNibbles(ByteArrayToNibbleArray(nextBytes))
		vm.printIfVerbose(fmt.Sprintf("%- 13x %-4x %s", nextBytes, vm.pc, instr.Text))
		if err != nil {
			return err
		}
		vm.Execute(instr)
		vm.pc += 4
	}
	return nil
}

func (vm *VM) loadInstructionsToMemory(instrs []Instruction) {
	for _, instr := range instrs {
		var nibbles = instr.Nibbles
		for i := 0; i < 4; i++ {
			vm.memory[vm.pc] = byte(nibbles[i])
			vm.pc++
		}
	}
	vm.pc = 0
}

func (vm *VM) PrintMemory() {
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			fmt.Printf("%2X", vm.memory[16*y+x])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func (vm *VM) PrintRegisters() {
	for _, regVal := range vm.registers {
		fmt.Printf("%2X", regVal)
		fmt.Print(" ")
	}
}

func (vm *VM) printIfVerbose(data string) {
	if vm.verboseLvl > 0 {
		fmt.Println(data)
	}
}

func (vm *VM) printifVVerbose(data string) {
	if vm.verboseLvl > 1 {
		fmt.Println(data)
	}
}
