package vm

import "fmt"

func (vm *VM) PrintMemory() {
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			fmt.Printf("%2X ", vm.memory[16*y+x])
		}
		fmt.Println()
	}
}

func (vm *VM) PrintRegisters() {
	for _, regVal := range vm.registers {
		fmt.Printf("%2X ", regVal)
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

func printBeforeData(data byte) {
	fmt.Sprintf("Before: %x", data)
}

func printAfterData(data byte) {
	fmt.Sprintf("After: %x", data)
}