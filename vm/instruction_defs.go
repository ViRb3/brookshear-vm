package vm

import "fmt"

var Instructions = []*Instruction{makeNop(), makeMoveMemToReg(), makeMoveValToReg(), makeMoveRegToMem(),
	makeMoveRegToReg(), makeAddIRegToReg(), makeAddFRegToReg(), makeAndRegToReg(), makeOrRegToReg(), makeXorRegToReg(),
	makeRotReg(), makeJmpIfEq(), makeHalt()}

func NewInstr() *Instruction {
	var instr = &Instruction{}
	return instr
}

func makeNop() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "nop"
	instr.OpcodeNibble = 0x0
	instr.Operands = []Operand{}
	instr.Format = instr.Opcode
	instr.Execute = func(vm *VM) {
		vm.doNop()
	}
	instr.FromNibbles = func([]Nibble) string {
		return fmt.Sprintf(instr.Format)
	}
	instr.ToNibbles = func() []Nibble {
		return []Nibble{0x0, 0x0, 0x0, 0x0}
	}
	instr.NewInstance = func() *Instruction {
		return makeNop()
	}
	return instr
}

func makeMoveMemToReg() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "mov"
	instr.OpcodeNibble = 0x1
	instr.Operands = []Operand{NewOperandBlank(OperandMemory), NewOperandBlank(OperandRegister)}
	instr.Format = instr.Opcode + " [%x] -> r%x"
	instr.Execute = func(vm *VM) {
		vm.doMoveMemToReg(instr.Operands[0].Value, instr.Operands[1].Value)
	}
	instr.FromNibbles = func(nibbles []Nibble) string {
		return fmt.Sprintf(instr.Format, CombineNibblesToByte(nibbles[2], nibbles[3]), nibbles[1])
	}
	instr.ToNibbles = func() []Nibble {
		return append([]Nibble{instr.OpcodeNibble, Nibble(instr.GetOpValAt(1))}, SplitByteToNibbles(instr.GetOpValAt(0))...)
	}
	instr.NewInstance = func() *Instruction {
		return makeMoveMemToReg()
	}
	return instr
}

func makeMoveValToReg() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "mov"
	instr.OpcodeNibble = 0x2
	instr.Operands = []Operand{NewOperandBlank(OperandValue), NewOperandBlank(OperandRegister)}
	instr.Format = instr.Opcode + " %x -> r%x"
	instr.Execute = func(vm *VM) {
		vm.doMoveValToReg(instr.Operands[0].Value, instr.Operands[1].Value)
	}
	instr.FromNibbles = func(nibbles []Nibble) string {
		return fmt.Sprintf(instr.Format, CombineNibblesToByte(nibbles[2], nibbles[3]), nibbles[1])
	}
	instr.ToNibbles = func() []Nibble {
		return append([]Nibble{instr.OpcodeNibble, Nibble(instr.GetOpValAt(1))}, SplitByteToNibbles(instr.GetOpValAt(0))...)
	}
	instr.NewInstance = func() *Instruction {
		return makeMoveValToReg()
	}
	return instr
}

func makeMoveRegToMem() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "mov"
	instr.OpcodeNibble = 0x3
	instr.Operands = []Operand{NewOperandBlank(OperandRegister), NewOperandBlank(OperandMemory)}
	instr.Format = instr.Opcode + " r%x -> [%x]"
	instr.Execute = func(vm *VM) {
		vm.doMoveRegToMem(instr.Operands[0].Value, instr.Operands[1].Value)
	}
	instr.FromNibbles = func(nibbles []Nibble) string {
		return fmt.Sprintf(instr.Format, nibbles[1], CombineNibblesToByte(nibbles[2], nibbles[3]))
	}
	instr.ToNibbles = func() []Nibble {
		return append([]Nibble{instr.OpcodeNibble, Nibble(instr.GetOpValAt(0))}, SplitByteToNibbles(instr.GetOpValAt(1))...)
	}
	instr.NewInstance = func() *Instruction {
		return makeMoveRegToMem()
	}
	return instr
}

//TODO: Check of ignored nibble is properly handled
func makeMoveRegToReg() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "mov"
	instr.OpcodeNibble = 0x4
	instr.Format = instr.Opcode + " r%x -> r%x"
	instr.Operands = []Operand{NewOperandBlank(OperandRegister), NewOperandBlank(OperandRegister)}
	instr.Execute = func(vm *VM) {
		vm.doMoveRegToReg(instr.Operands[0].Value, instr.Operands[1].Value)
	}
	instr.FromNibbles = func(nibbles []Nibble) string {
		return fmt.Sprintf(instr.Format, nibbles[2], nibbles[3])
	}
	instr.ToNibbles = func() []Nibble {
		return []Nibble{instr.OpcodeNibble, 0x0, Nibble(instr.GetOpValAt(0)), Nibble(instr.GetOpValAt(1))}
	}
	instr.NewInstance = func() *Instruction {
		return makeMoveRegToReg()
	}
	return instr
}

func makeAddIRegToReg() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "addi"
	instr.OpcodeNibble = 0x5
	instr.Operands = []Operand{NewOperandBlank(OperandRegister), NewOperandBlank(OperandRegister), NewOperandBlank(OperandRegister)}
	instr.Format = instr.Opcode + " r%x, r%x -> r%x"
	instr.Execute = func(vm *VM) {
		vm.doAddIRegToReg(instr.Operands[0].Value, instr.Operands[1].Value, instr.Operands[2].Value)
	}
	instr.FromNibbles = func(nibbles []Nibble) string {
		return fmt.Sprintf(instr.Format, nibbles[2], nibbles[3], nibbles[1])
	}
	instr.ToNibbles = func() []Nibble {
		return []Nibble{instr.OpcodeNibble, Nibble(instr.GetOpValAt(2)), Nibble(instr.GetOpValAt(0)),
			Nibble(instr.GetOpValAt(1))}
	}
	instr.NewInstance = func() *Instruction {
		return makeAddIRegToReg()
	}
	return instr
}

func makeAddFRegToReg() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "addf"
	instr.OpcodeNibble = 0x6
	instr.Operands = []Operand{NewOperandBlank(OperandRegister), NewOperandBlank(OperandRegister), NewOperandBlank(OperandRegister)}
	instr.Format = instr.Opcode + " r%x, r%x -> r%x"
	instr.Execute = func(vm *VM) {
		vm.doAddFRegToReg(instr.Operands[0].Value, instr.Operands[1].Value, instr.Operands[2].Value)
	}
	instr.FromNibbles = func(nibbles []Nibble) string {
		return fmt.Sprintf(instr.Format, nibbles[2], nibbles[3], nibbles[1])
	}
	instr.ToNibbles = func() []Nibble {
		return []Nibble{instr.OpcodeNibble, Nibble(instr.GetOpValAt(2)), Nibble(instr.GetOpValAt(0)),
			Nibble(instr.GetOpValAt(1))}
	}
	instr.NewInstance = func() *Instruction {
		return makeAddFRegToReg()
	}
	return instr
}

func makeOrRegToReg() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "or"
	instr.OpcodeNibble = 0x7
	instr.Operands = []Operand{NewOperandBlank(OperandRegister), NewOperandBlank(OperandRegister), NewOperandBlank(OperandRegister)}
	instr.Format = instr.Opcode + " r%x, r%x -> r%x"
	instr.Execute = func(vm *VM) {
		vm.doOrRegWithReg(instr.Operands[0].Value, instr.Operands[1].Value, instr.Operands[2].Value)
	}
	instr.FromNibbles = func(nibbles []Nibble) string {
		return fmt.Sprintf(instr.Format, nibbles[2], nibbles[3], nibbles[1])
	}
	instr.ToNibbles = func() []Nibble {
		return []Nibble{instr.OpcodeNibble, Nibble(instr.GetOpValAt(2)), Nibble(instr.GetOpValAt(0)),
			Nibble(instr.GetOpValAt(1))}
	}
	instr.NewInstance = func() *Instruction {
		return makeOrRegToReg()
	}
	return instr
}

func makeAndRegToReg() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "and"
	instr.OpcodeNibble = 0x8
	instr.Operands = []Operand{NewOperandBlank(OperandRegister), NewOperandBlank(OperandRegister), NewOperandBlank(OperandRegister)}
	instr.Format = instr.Opcode + " r%x, r%x -> r%x"
	instr.Execute = func(vm *VM) {
		vm.doAndRegWithReg(instr.Operands[0].Value, instr.Operands[1].Value, instr.Operands[2].Value)
	}
	instr.FromNibbles = func(nibbles []Nibble) string {
		return fmt.Sprintf(instr.Format, nibbles[2], nibbles[3], nibbles[1])
	}
	instr.ToNibbles = func() []Nibble {
		return []Nibble{instr.OpcodeNibble, Nibble(instr.GetOpValAt(2)), Nibble(instr.GetOpValAt(0)),
			Nibble(instr.GetOpValAt(1))}
	}
	instr.NewInstance = func() *Instruction {
		return makeAndRegToReg()
	}
	return instr
}

func makeXorRegToReg() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "xor"
	instr.OpcodeNibble = 0x9
	instr.Operands = []Operand{NewOperandBlank(OperandRegister), NewOperandBlank(OperandRegister), NewOperandBlank(OperandRegister)}
	instr.Format = instr.Opcode + " r%x, r%x -> r%x"
	instr.Execute = func(vm *VM) {
		vm.doXorRegWithReg(instr.Operands[0].Value, instr.Operands[1].Value, instr.Operands[2].Value)
	}
	instr.FromNibbles = func(nibbles []Nibble) string {
		return fmt.Sprintf(instr.Format, nibbles[2], nibbles[3], nibbles[1])
	}
	instr.ToNibbles = func() []Nibble {
		return []Nibble{instr.OpcodeNibble, Nibble(instr.GetOpValAt(2)), Nibble(instr.GetOpValAt(0)),
			Nibble(instr.GetOpValAt(1))}
	}
	instr.NewInstance = func() *Instruction {
		return makeXorRegToReg()
	}
	return instr
}

//TODO: Rot should only support 1 nibble of value
func makeRotReg() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "rot"
	instr.OpcodeNibble = 0xa
	instr.Operands = []Operand{NewOperandBlank(OperandRegister), NewOperandBlank(OperandValue)}
	instr.Format = instr.Opcode + " r%x, %x"
	instr.Execute = func(vm *VM) {
		vm.doRotReg(instr.Operands[0].Value, instr.Operands[1].Value)
	}
	instr.FromNibbles = func(nibbles []Nibble) string {
		return fmt.Sprintf(instr.Format, nibbles[1], nibbles[3])
	}
	instr.ToNibbles = func() []Nibble {
		return append([]Nibble{instr.OpcodeNibble, Nibble(instr.GetOpValAt(0)),}, SplitByteToNibbles(instr.GetOpValAt(1))...)
	}
	instr.NewInstance = func() *Instruction {
		return makeRotReg()
	}
	return instr
}

func makeJmpIfEq() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "jmpeq"
	instr.OpcodeNibble = 0xb
	instr.Operands = []Operand{NewOperandBlank(OperandValue), NewOperandBlank(OperandRegister)}
	instr.Format = instr.Opcode + " %x, r%x"
	instr.Execute = func(vm *VM) {
		vm.doJmpIfEq(instr.Operands[0].Value, instr.Operands[1].Value)
	}
	instr.FromNibbles = func(nibbles []Nibble) string {
		return fmt.Sprintf(instr.Format, CombineNibblesToByte(nibbles[2], nibbles[3]), nibbles[1])
	}
	instr.ToNibbles = func() []Nibble {
		return append([]Nibble{instr.OpcodeNibble, Nibble(instr.GetOpValAt(1))}, SplitByteToNibbles(instr.GetOpValAt(0))...)
	}
	instr.NewInstance = func() *Instruction {
		return makeJmpIfEq()
	}
	return instr
}

func makeHalt() (instr *Instruction) {
	instr = NewInstr()
	instr.Opcode = "halt"
	instr.OpcodeNibble = 0xc
	instr.Operands = []Operand{}
	instr.Format = instr.Opcode
	instr.Execute = func(vm *VM) {
		vm.doHalt()
	}
	instr.FromNibbles = func([]Nibble) string {
		return fmt.Sprintf(instr.Format)
	}
	instr.ToNibbles = func() []Nibble {
		return []Nibble{0xc, 0x0, 0x0, 0x0}
	}
	instr.NewInstance = func() *Instruction {
		return makeHalt()
	}
	return instr
}
