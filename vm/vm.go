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

func New(verboseLvl int) *VM {
	var vm = VM{}
	vm.memory = make([]byte, 256)
	vm.registers = make([]byte, 16)
	vm.verboseLvl = verboseLvl
	return &vm
}

func ParseInstructions(instrStrs []string) ([]*Instruction, error) {
	var instrs []*Instruction

	for i, instrStr := range instrStrs {
		if IsIgnoredLine(instrStr) {
			continue
		}
		instrStr = RemoveTrailingComment(instrStr)

		instr, err := Parse(instrStr, i)
		if err != nil {
			return instrs, makeInstructionParseError(err, instrStr, i)
		}
		instrs = append(instrs, instr)
	}
	return instrs, nil
}

func (vm *VM) Execute(instr *Instruction) {
	instr.Execute(vm)
}

func (vm *VM) Run(instrs[]*Instruction) error {
	vm.loadInstructionsInMemory(instrs)

	vm.printIfVerbose("Execution flow:")
	if err := vm.instrLoop(); err != nil {
		return err
	}
	return nil
}

func (vm *VM) instrLoop() error {
	for !vm.halt {
		var nextBytes = vm.memory[vm.pc : vm.pc+4]
		var instr, err = ParseFromNibbles(BytesToNibbles(nextBytes))
		vm.printIfVerbose(fmt.Sprintf("%-9s %-4x %s", PrettyPrintNibbles(BytesToNibbles(nextBytes)), vm.pc, instr.GetText()))
		if err != nil {
			return err
		}
		vm.Execute(instr)
		vm.pc += 4
	}
	return nil
}

func (vm *VM) loadInstructionsInMemory(instrs []*Instruction) {
	for _, instr := range instrs {
		var nibbles = instr.Nibbles
		for i := 0; i < 4; i++ {
			vm.memory[vm.pc] = byte(nibbles[i])
			vm.pc++
		}
	}
	vm.pc = 0
}

func makeInstructionParseError(err error, instrStr string, i int)(error) {
	err = errors.New(fmt.Sprintf("%s\nInstruction: %s\nLine: %d", err.Error(), instrStr, i+1))
	return err
}