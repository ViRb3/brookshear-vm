package vm

import "fmt"

func (vm *VM) PrintMemory() {
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			PrintByteVerboseIfNonZero(vm.memory[16*y+x])
		}
		fmt.Println()
	}
}

func (vm *VM) PrintRegisters() {
	for _, regVal := range vm.registers {
		PrintByteVerboseIfNonZero(regVal)
	}
}

func (vm *VM) printIfVerbose(data string) {
	if vm.verboseLvl > 0 {
		fmt.Print(data)
	}
}

func (vm *VM) printifVVerbose(data string) {
	if vm.verboseLvl > 1 {
		fmt.Print(data)
	}
}