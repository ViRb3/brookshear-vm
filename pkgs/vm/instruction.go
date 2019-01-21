package vm

import (
	"github.com/pkg/errors"
	"strings"
)

type Instruction struct {
	Opcode           string // e.g. mov
	OpcodeNibble     Nibble
	Format           string // e.g. mov [%x] -> r%x
	DestOperandIndex int    // operand index the instruction modifies (register or memory cell); -1 for none
	InstructionData
	InstructionMethods
}

type InstructionMethods struct {
	Execute     func(vm *VM)
	FromNibbles func([]Nibble) string
	ToNibbles   func() []Nibble
	NewInstance func() *Instruction
}

type InstructionData struct {
	SourceLine int
	Operands   []Operand
	Nibbles    []Nibble
}

func NewInstr() *Instruction {
	var instr = &Instruction{}
	instr.DestOperandIndex = -1
	return instr
}

func (instr *Instruction) ToString() string {
	return instr.FromNibbles(instr.Nibbles)
}

// Get operand value at
func (instr *InstructionData) GetOpValAt(i int) byte {
	return instr.Operands[i].Value
}

// Get operand extra value at
func (instr *InstructionData) GetOpExtraAt(i int) byte {
	return instr.Operands[i].Extra
}

func (instr *InstructionData) GetOperandsInOrder(order []int) []Operand {
	var result []Operand
	for _, i := range order {
		result = append(result, instr.Operands[i])
	}
	return result
}

func Parse(instrStr string, sourceIndex int) (instr *Instruction, err error) {
	instrStr = strings.ToLower(instrStr)
	var instrData InstructionData

	if err := instrData.parseOperands(instrStr); err != nil {
		return instr, err
	}

	instr, err = matchWithTemplate(instrStr, instrData)
	if err != nil {
		return instr, errors.New("unable to parse instruction")
	}

	instr.Nibbles = instr.ToNibbles()
	instr.SourceLine = sourceIndex
	return instr, nil
}

func ParseFromNibbles(nibbles []Nibble) (*Instruction, error) {
	for i, instrTemplate := range Instructions {
		if instrTemplate.OpcodeNibble == nibbles[0] {
			return Parse(instrTemplate.FromNibbles(nibbles), i)
		}
	}
	return &Instruction{}, errors.New("unable to parse nibbles")
}

func matchWithTemplate(instrStr string, instrData InstructionData) (*Instruction, error) {
	// match parsed data with defined instruction
	for _, instrTemplate := range Instructions {
		var instrOpcode = strings.SplitN(instrStr, " ", 2)[0]
		if instrOpcode != instrTemplate.Opcode {
			continue
		}
		if len(instrData.Operands) != len(instrTemplate.Operands) {
			continue
		}
		var equal = true
		for i := range instrData.Operands {
			if instrData.Operands[i].Type != instrTemplate.Operands[i].Type {
				equal = false
				break
			}
		}
		if equal {
			// create new copy of the proper instruction
			instrTemplate = instrTemplate.NewInstance()
			// inject parsed data
			instrTemplate.InstructionData = instrData
			return instrTemplate, nil
		}
	}
	return &Instruction{}, errors.New("unable to match instruction")
}
