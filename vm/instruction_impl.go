package vm

import "fmt"

//0iii NOP
func (vm *VM) doNop() {
	vm.executeInstruction(func() {
		// NOP
	}, nil)
}

//1rxy MOV [xy] -> Rr
func (vm *VM) doMoveMemToReg(srcMem byte, dstReg byte) {
	vm.executeInstruction(func() {
		vm.registers[dstReg] = vm.memory[srcMem]
	}, &vm.registers[dstReg])
}

//2rxy MOV xy -> Rr
func (vm *VM) doMoveValToReg(val byte, dstReg byte) {
	vm.executeInstruction(func() {
		vm.registers[dstReg] = val
	}, &vm.registers[dstReg])
}

//3rxy MOV Rr -> [xy]
func (vm *VM) doMoveRegToMem(srcReg byte, dstMem byte) {
	vm.executeInstruction(func() {
		vm.memory[dstMem] = vm.registers[srcReg]
	}, &vm.memory[dstMem])
}

//4irs MOV Rr -> Rs
func (vm *VM) doMoveRegToReg(srcReg byte, dstReg byte) {
	vm.executeInstruction(func() {
		vm.registers[srcReg] = vm.registers[dstReg]
	}, &vm.registers[srcReg])
}

//5rst ADDI Rs, Rt -> Rr
func (vm *VM) doAddIRegToReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.executeInstruction(func() {
		vm.registers[dstReg] = vm.registers[srcReg1] + vm.registers[srcReg2]
	}, &vm.registers[dstReg])
}

//6rst ADDF Rs, Rt -> Rr
func (vm *VM) doAddFRegToReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	//TODO: Implement ADDF
	fmt.Println("UNSUPPORTED OPCODE 'ADDF'")
}

//7rst OR Rs, Rt -> Rr
func (vm *VM) doOrRegWithReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.executeInstruction(func() {
		vm.registers[dstReg] = vm.registers[srcReg1] + vm.registers[srcReg2]
	}, &vm.registers[dstReg])
}

//8rst AND Rs, Rt -> Rr
func (vm *VM) doAndRegWithReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.executeInstruction(func() {
		vm.registers[dstReg] = vm.registers[srcReg1] + vm.registers[srcReg2]
	}, &vm.registers[dstReg])
}

//9rst XOR Rs, Rt -> Rr
func (vm *VM) doXorRegWithReg(srcReg1 byte, srcReg2 byte, dstReg byte) {
	vm.executeInstruction(func() {
		vm.registers[dstReg] = vm.registers[srcReg1] + vm.registers[srcReg2]
	}, &vm.registers[dstReg])
}

//Ariz ROT Rr, z ; circular shift RIGHT
func (vm *VM) doRotReg(dstReg byte, times byte) {
	var val = ConvByteToBitStringArray(vm.registers[dstReg])
	RotateRCircular(&val, int(times))

	vm.executeInstruction(func() {
		vm.registers[dstReg] = ConvBitStringArrayToByte(val)
	}, &vm.registers[dstReg])

}

//Brxy JMPEQ xy, Rr
func (vm *VM) doJmpIfEq(val byte, reg byte) {
	vm.executeInstruction(func() {
		// account for the pc+4 every cycle
		// TODO: Cleaner approach for PC fix?
		if vm.registers[reg] == vm.registers[0] {
			vm.pc = val - 2
		}
	}, nil)

}

//Ciii HALT
func (vm *VM) doHalt() {
	vm.halt = true
}

func (vm *VM) executeInstruction(action func(), dstValue *byte) {
	if dstValue == nil {
		action()
		return
	} else {
		var oldVal = *dstValue
		action()
		var newVal = *dstValue
		vm.printifVVerbose(fmt.Sprintf("%X -> %X", oldVal, newVal))
	}
}
