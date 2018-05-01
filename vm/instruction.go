package vm

import (
	"github.com/pkg/errors"
	"fmt"
)

type Instruction struct {
	Text              string
	Opcode            string
	LeftOperands      []Operand
	RightOperands     []Operand
	CallingConvention string
	Execute           func(vm *VM)
	Nibbles           []Nibble
}

func Parse(instrStr string) (Instruction, error) {
	var instr Instruction
	instr.Text = instrStr

	if err := instr.parseOperands(instrStr); err != nil {
		return instr, err
	}
	instr.parseCallingConvention()
	if err := instr.parseExecuteMethodsAndBytes(); err != nil {
		return instr, err
	}
	return instr, nil
}

func ParseFromNibbles(nibbles []Nibble) (Instruction, error) {
	var instr Instruction
	if len(nibbles) != 4 {
		return instr, errors.New("unable to parse instruction (bad nibble length)")
	}

	switch nibbles[0] {
	case 0x0:
		return Parse(fmt.Sprintf("nop"))
	case 0x1:
		return Parse(fmt.Sprintf("mov [%x] -> r%x", CombineNibblesToByte(nibbles[1], nibbles[2]), nibbles[3]))
	case 0x2:
		return Parse(fmt.Sprintf("mov %x -> r%x", CombineNibblesToByte(nibbles[1], nibbles[2]), nibbles[3]))
	case 0x3:
		return Parse(fmt.Sprintf("mov r%x -> [%x]", nibbles[1], CombineNibblesToByte(nibbles[2], nibbles[3])))
	case 0x4:
		return Parse(fmt.Sprintf("mov r%x -> r%x", nibbles[2], nibbles[3]))
	case 0x5:
		return Parse(fmt.Sprintf("addi r%x, r%x -> r%x", nibbles[1], nibbles[2], nibbles[3]))
	case 0x6:
		return Parse(fmt.Sprintf("addf r%x, r%x -> r%x", nibbles[1], nibbles[2], nibbles[3]))
	case 0x7:
		return Parse(fmt.Sprintf("or r%x, r%x -> r%x", nibbles[1], nibbles[2], nibbles[3]))
	case 0x8:
		return Parse(fmt.Sprintf("and r%x, r%x -> r%x", nibbles[1], nibbles[2], nibbles[3]))
	case 0x9:
		return Parse(fmt.Sprintf("xor r%x, r%x -> r%x", nibbles[1], nibbles[2], nibbles[3]))
	case 0xa:
		return Parse(fmt.Sprintf("rot r%x, %x", nibbles[1], CombineNibblesToByte(nibbles[2], nibbles[3])))
	case 0xb:
		return Parse(fmt.Sprintf("jmpeq %x, r%x", CombineNibblesToByte(nibbles[1], nibbles[2]), nibbles[3]))
	case 0xc:
		return Parse(fmt.Sprintf("halt"))
	default:
		return instr, errors.New("unable to parse instruction (bad opcode)")
	}
}

func (instr *Instruction) parseExecuteMethodsAndBytes() error {
	var execMethod func(vm *VM)
	var nibbles []Nibble

	switch instr.CallingConvention {
	case "nop":
		execMethod = func(vm *VM) { vm.doNop() }
		nibbles = []Nibble{0x00, 0x00, 0x00, 0x00}
	case "mov:[->r":
		execMethod = func(vm *VM) { vm.doMoveMemToReg(instr.LeftOperands[0].Value, instr.RightOperands[0].Value) }
		var splitOp = SplitByteToNibbles(instr.LeftOperands[0].Value)
		nibbles = []Nibble{0x01, splitOp[0], splitOp[1], Nibble(instr.RightOperands[0].Value)}
	case "mov:v->r":
		execMethod = func(vm *VM) { vm.doMoveValToReg(instr.LeftOperands[0].Value, instr.RightOperands[0].Value) }
		var splitOp = SplitByteToNibbles(instr.LeftOperands[0].Value)
		nibbles = []Nibble{0x02, splitOp[0], splitOp[1], Nibble(instr.RightOperands[0].Value)}
	case "mov:r->[":
		execMethod = func(vm *VM) { vm.doMoveRegToMem(instr.LeftOperands[0].Value, instr.RightOperands[0].Value) }
		var splitOp = SplitByteToNibbles(instr.RightOperands[0].Value)
		nibbles = []Nibble{0x03, Nibble(instr.LeftOperands[0].Value), splitOp[0], splitOp[1]}
	case "mov:r->r":
		execMethod = func(vm *VM) { vm.doMoveRegToReg(instr.LeftOperands[0].Value, instr.RightOperands[0].Value) }
		nibbles = []Nibble{0x04, 0x00, Nibble(instr.LeftOperands[0].Value), Nibble(instr.RightOperands[0].Value)}
	case "addi:rr->r":
		execMethod = func(vm *VM) { vm.doAddIRegToReg(instr.LeftOperands[0].Value, instr.LeftOperands[1].Value, instr.RightOperands[0].Value) }
		nibbles = []Nibble{0x05, Nibble(instr.LeftOperands[0].Value), Nibble(instr.LeftOperands[1].Value), Nibble(instr.RightOperands[0].Value)}
	case "addf:rr->r":
		execMethod = func(vm *VM) { vm.doAddFRegToReg(instr.LeftOperands[0].Value, instr.LeftOperands[1].Value, instr.RightOperands[0].Value) }
		nibbles = []Nibble{0x06, Nibble(instr.LeftOperands[0].Value), Nibble(instr.LeftOperands[1].Value), Nibble(instr.RightOperands[0].Value)}
	case "or:rr->r":
		execMethod = func(vm *VM) { vm.doOrRegWithReg(instr.LeftOperands[0].Value, instr.LeftOperands[1].Value, instr.RightOperands[0].Value) }
		nibbles = []Nibble{0x07, Nibble(instr.LeftOperands[0].Value), Nibble(instr.LeftOperands[1].Value), Nibble(instr.RightOperands[0].Value)}
	case "and:rr->r":
		execMethod = func(vm *VM) { vm.doAndRegWithReg(instr.LeftOperands[0].Value, instr.LeftOperands[1].Value, instr.RightOperands[0].Value) }
		nibbles = []Nibble{0x08, Nibble(instr.LeftOperands[0].Value), Nibble(instr.LeftOperands[1].Value), Nibble(instr.RightOperands[0].Value)}
	case "xor:rr->r":
		execMethod = func(vm *VM) { vm.doXorRegWithReg(instr.LeftOperands[0].Value, instr.LeftOperands[1].Value, instr.RightOperands[0].Value) }
		nibbles = []Nibble{0x09, Nibble(instr.LeftOperands[0].Value), Nibble(instr.LeftOperands[1].Value), Nibble(instr.RightOperands[0].Value)}
	case "rot:rv":
		execMethod = func(vm *VM) { vm.doRotReg(instr.LeftOperands[0].Value, instr.LeftOperands[1].Value) }
		var splitOp = SplitByteToNibbles(instr.LeftOperands[1].Value)
		nibbles = []Nibble{0x0a, Nibble(instr.LeftOperands[0].Value), 0x0, splitOp[1]}
	case "jmpeq:vr":
		execMethod = func(vm *VM) { vm.doJmpIfEq(instr.LeftOperands[0].Value, instr.LeftOperands[1].Value) }
		var splitOp = SplitByteToNibbles(instr.LeftOperands[0].Value)
		nibbles = []Nibble{0x0b, splitOp[0], splitOp[1], Nibble(instr.LeftOperands[1].Value)}
	case "halt":
		execMethod = func(vm *VM) { vm.doHalt() }
		nibbles = []Nibble{0x0c, 0x00, 0x00, 0x00}
	default:
		return errors.New("unable to parse instruction (bad opcode+operand combination)")
	}

	instr.Execute = execMethod
	instr.Nibbles = nibbles
	return nil
}

func (instr *Instruction) parseCallingConvention() {
	instr.CallingConvention = instr.Opcode
	if len(instr.LeftOperands) > 0 {
		instr.CallingConvention += ":"
	}
	for _, op := range instr.LeftOperands {
		instr.parseOperandCallingConvention(op)
	}
	if len(instr.RightOperands) > 0 {
		instr.CallingConvention += "->"
	}
	for _, op := range instr.RightOperands {
		instr.parseOperandCallingConvention(op)
	}
}

func (instr *Instruction) parseOperandCallingConvention(op Operand) {
	if op.Type == OperandAddress {
		instr.CallingConvention += "["
	} else if op.Type == OperandValue {
		instr.CallingConvention += "v"
	} else if op.Type == OperandRegister {
		instr.CallingConvention += "r"
	}
}
