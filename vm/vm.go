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

	vm.printIfVerbose("Execution flow:\n")
	if err := vm.instrLoop(); err != nil {
		return err
	}
	return nil
}

func (vm *VM) instrLoop() error {
	for !vm.halt {
		if int(vm.pc)+2 == len(vm.memory) {
			return errors.New("reached end of memory")
		}
		var nextBytes = vm.memory[vm.pc : vm.pc+2]
		var nextNibbles = ByteArrayToNibbleArray(nextBytes)
		var instr, err = ParseFromNibbles(nextNibbles)
		if err != nil {
			return err
		}

		// don't add new line yet to append more information in later calls
		var logData = fmt.Sprintf("%-6s %-4x %-22s", PrintNibbles(nextNibbles), vm.pc, instr.ToString())
		vm.printIfVerbose(logData)

		// if there's a changed register or memory cell, log that in VV
		if instr.DestOperandIndex > -1 {
			var destName = instr.Operands[instr.DestOperandIndex].ToString()
			vm.printifVVerbose( fmt.Sprintf("%-5s: ", destName))
		}

		vm.Execute(instr)
		vm.pc += 2
		vm.printIfVerbose("\n")
	}
	return nil
}

func (vm *VM) loadInstructionsInMemory(instrs []*Instruction) {
	for _, instr := range instrs {
		vm.memory[vm.pc] = CombineNibblesToByte(instr.Nibbles[0], instr.Nibbles[1])
		vm.pc++
		vm.memory[vm.pc] = CombineNibblesToByte(instr.Nibbles[2], instr.Nibbles[3])
		vm.pc++
	}
	vm.pc = 0
}

func makeInstructionParseError(err error, instrStr string, i int)(error) {
	err = errors.New(fmt.Sprintf("%s\nInstruction: %s\nLine: %d", err.Error(), instrStr, i+1))
	return err
}