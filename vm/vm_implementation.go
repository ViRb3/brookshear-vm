package vm

//0iii NOP
func (vm *VM) doNop() {
	// NOP
}

//1rxy MOV [xy] -> Rr
func (vm *VM) doMoveMemToReg(srcMem byte, dstReg byte) {
	vm.registers[dstReg] = vm.memory[srcMem]
}

//2rxy MOV xy -> Rr
func (vm *VM) doMoveValToReg(val byte, dstReg byte) {
	vm.registers[dstReg] = val
}

//3rxy MOV Rr -> [xy]
func (vm *VM) doMoveRegToMem(srcReg byte, dstMem byte) {
	vm.memory[dstMem] = vm.registers[srcReg]
}

//4irs MOV Rr -> Rs
func (vm *VM) doMoveRegToReg(srcReg byte, dstReg byte) {
	vm.registers[srcReg] = vm.registers[dstReg]
}

//5rst ADDI Rs, Rt -> Rr
func (vm *VM) doAddIRegToReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.registers[dstReg] = vm.registers[srcReg1] + vm.registers[srcReg2]
}

//6rst ADDF Rs, Rt -> Rr
func (vm *VM) doAddFRegToReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	//TODO: Implement
	vm.doAddIRegToReg(srcReg1, srcReg2, dstReg)
}

//7rst OR Rs, Rt -> Rr
func (vm *VM) doOrRegWithReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.registers[dstReg] = vm.registers[srcReg1] | vm.registers[srcReg2]
}

//8rst AND Rs, Rt -> Rr
func (vm *VM) doAndRegWithReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.registers[dstReg] = vm.registers[srcReg1] & vm.registers[srcReg2]
}

//9rst XOR Rs, Rt -> Rr
func (vm *VM) doXorRegWithReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.registers[dstReg] = vm.registers[srcReg1] ^ vm.registers[srcReg2]
}

//Ariz ROT Rr, z ; circular shift RIGHT
func (vm *VM) doRotReg(dstReg byte, times byte) {
	var val = ConvByteToBitStringArray(vm.registers[dstReg])
	RotateRCircular(&val, int(times))
	vm.registers[dstReg] = ConvBitStringArrayToByte(val)
}

//Brxy JMPEQ xy, Rr
func (vm *VM) doJmpIfEq(val byte, reg byte) {
	// account for the pc+4 every cycle
	// TODO: Cleaner approach?
	if vm.registers[reg] == vm.registers[0] {
		vm.pc = val - 4
	}
}

//Ciii HALT
func (vm *VM) doHalt() {
	vm.halt = true
}

func (vm *VM) executeInstruction(action func(), dstValue *byte) {
	printBeforeData(*dstValue)
	action()
	printAfterData(*dstValue)
}
