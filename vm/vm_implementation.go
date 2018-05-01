package vm

import "fmt"

//0iii NOP
func (vm *VM) doNop() {
	// NOP
}

//1rxy MOV [xy] -> Rr
func (vm *VM) doMoveMemToReg(srcMem byte, dstReg byte) {
	vm.printifVVerbose(fmt.Sprintf("Before: %x", vm.registers[dstReg]))
	vm.registers[dstReg] = vm.memory[srcMem]
	vm.printifVVerbose(fmt.Sprintf("After: %x", vm.registers[dstReg]))
}

//2rxy MOV xy -> Rr
func (vm *VM) doMoveValToReg(val byte, dstReg byte) {
	vm.printifVVerbose(fmt.Sprintf("Before: %x", vm.registers[dstReg]))
	vm.registers[dstReg] = val
	vm.printifVVerbose(fmt.Sprintf("After: %x", vm.registers[dstReg]))
}

//3rxy MOV Rr -> [xy]
func (vm *VM) doMoveRegToMem(srcReg byte, dstMem byte) {
	vm.printifVVerbose(fmt.Sprintf("Before: %x", vm.memory[dstMem]))
	vm.memory[dstMem] = vm.registers[srcReg]
	vm.printifVVerbose(fmt.Sprintf("After: %x", vm.memory[dstMem]))
}

//4irs MOV Rr -> Rs
func (vm *VM) doMoveRegToReg(srcReg byte, dstReg byte) {
	vm.printifVVerbose(fmt.Sprintf("Before: %x", vm.registers[srcReg]))
	vm.registers[srcReg] = vm.registers[dstReg]
	vm.printifVVerbose(fmt.Sprintf("After: %x", vm.registers[srcReg]))
}

//5rst ADDI Rs, Rt -> Rr
func (vm *VM) doAddIRegToReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.printifVVerbose(fmt.Sprintf("Before: %x", vm.registers[dstReg]))
	vm.registers[dstReg] = vm.registers[srcReg1] + vm.registers[srcReg2]
	vm.printifVVerbose(fmt.Sprintf("After: %x", vm.registers[dstReg]))
}

//6rst ADDF Rs, Rt -> Rr
func (vm *VM) doAddFRegToReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	//TODO: Implement
	vm.doAddIRegToReg(srcReg1, srcReg2, dstReg)
}

//7rst OR Rs, Rt -> Rr
func (vm *VM) doOrRegWithReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.printifVVerbose(fmt.Sprintf("Before: %x", vm.registers[dstReg]))
	vm.registers[dstReg] = vm.registers[srcReg1] | vm.registers[srcReg2]
	vm.printifVVerbose(fmt.Sprintf("After: %x", vm.registers[dstReg]))
}

//8rst AND Rs, Rt -> Rr
func (vm *VM) doAndRegWithReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.printifVVerbose(fmt.Sprintf("Before: %x", vm.registers[dstReg]))
	vm.registers[dstReg] = vm.registers[srcReg1] & vm.registers[srcReg2]
	vm.printifVVerbose(fmt.Sprintf("After: %x", vm.registers[dstReg]))
}

//9rst XOR Rs, Rt -> Rr
func (vm *VM) doXorRegWithReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.printifVVerbose(fmt.Sprintf("Before: %x", vm.registers[dstReg]))
	vm.registers[dstReg] = vm.registers[srcReg1] ^ vm.registers[srcReg2]
	vm.printifVVerbose(fmt.Sprintf("After: %x", vm.registers[dstReg]))
}

//Ariz ROT Rr, z ; circular shift RIGHT
func (vm *VM) doRotReg(dstReg byte, times byte) {
	vm.printifVVerbose(fmt.Sprintf("Before: %x", vm.registers[dstReg]))
	var val = ConvByteToBitStringArray(vm.registers[dstReg])
	RotateRCircular(&val, int(times))
	vm.registers[dstReg] = ConvBitStringArrayToByte(val)
	vm.printifVVerbose(fmt.Sprintf("After: %x", vm.registers[dstReg]))
}

//Brxy JMPEQ xy, Rr
func (vm *VM) doJmpIfEq(val byte, reg byte) {
	vm.printifVVerbose(fmt.Sprintf("Before: %x", vm.pc))

	// account for the pc+4 every cycle
	// TODO: Cleaner approach?

	if vm.registers[reg] == vm.registers[0] {
		vm.pc = val - 4
	}
	vm.printifVVerbose(fmt.Sprintf("After: %x", vm.pc+4))
}

//Ciii HALT
func (vm *VM) doHalt() {
	vm.halt = true
}
