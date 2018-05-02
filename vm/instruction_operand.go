package vm

import (
	"strings"
	"regexp"
	"github.com/pkg/errors"
	"strconv"
)

//TODO: Support for parsing with stacked intervals
var operandRegex = regexp.MustCompile(` -> |, | `)

var addressRegex = regexp.MustCompile(`^\[([0-9a-f]{1,2})\]$`)
var valueRegex = regexp.MustCompile(`^([0-9a-f]{1,2})$`)
var registerRegex = regexp.MustCompile(`^r([0-9a-f])$`)

type Operand struct {
	Type  OperandType
	Value byte
}

type OperandType int

const (
	OperandMemory   = iota
	OperandRegister
	OperandValue
)

func NewOperand(operandType OperandType, value string) Operand {
	var valueInt, _ = strconv.ParseInt(value, 16, 0)
	return Operand{operandType, byte(valueInt)}
}

func NewOperandBlank(operandType OperandType) Operand {
	return Operand{operandType, 0}
}

func (instr *InstructionData) parseOperands(instrStr string) (error) {
	// also remove opcode at index 0
	var operands = operandRegex.Split(strings.TrimSpace(instrStr), -1)[1:]

	for i := range operands {
		parsedOperand, err := parseOperand(operands[i])
		if err != nil {
			return err
		}
		instr.Operands = append(instr.Operands, parsedOperand)
	}

	return nil
}

func parseOperand(operandStr string) (Operand, error) {
	operandStr = strings.ToLower(operandStr)

	if operand, err := tryParseAsOperandType(operandStr, addressRegex, OperandMemory); err == nil {
		return operand, nil
	}
	if operand, err := tryParseAsOperandType(operandStr, valueRegex, OperandValue); err == nil {
		return operand, nil
	}
	if operand, err := tryParseAsOperandType(operandStr, registerRegex, OperandRegister); err == nil {
		return operand, nil
	}
	return Operand{}, errors.New("unable to parse operand")
}

func tryParseAsOperandType(operandStr string, parseRegex *regexp.Regexp, operandType OperandType) (Operand, error) {
	var result = parseRegex.FindStringSubmatch(operandStr)
	if len(result) < 2 {
		return Operand{}, errors.New("unable to parse operand as current type")
	}
	return NewOperand(operandType, result[1]), nil
}
