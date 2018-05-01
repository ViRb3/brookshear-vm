package vm

import (
	"strings"
	"regexp"
	"strconv"
	"github.com/pkg/errors"
)

var operandRegex = regexp.MustCompile(`( +| *, +)`)

var addressRegex = regexp.MustCompile(`^\[([0-9a-f]{1,2})\]$`)
var valueRegex = regexp.MustCompile(`^([0-9a-f]{1,2})$`)
var registerRegex = regexp.MustCompile(`^r([0-9a-f])$`)

type Operand struct {
	Type  OperandType
	Value byte
}

type OperandType int

const (
	OperandAddress  = iota
	OperandRegister
	OperandValue
)

func (instr *Instruction) parseOperands(instrStr string) (error) {
	var operandSides = strings.Split(instrStr, "->")

	var leftOperands = operandRegex.Split(strings.TrimSpace(operandSides[0]), -1)
	instr.Opcode = leftOperands[0]
	// remove opcode
	leftOperands = leftOperands[1:]

	if err := parseAndAddOperandSide(leftOperands, &instr.LeftOperands); err != nil {
		return err
	}

	// right operand side
	if len(operandSides) == 2 {
		var rightOperands = operandRegex.Split(strings.TrimSpace(operandSides[1]), -1)
		if err := parseAndAddOperandSide(rightOperands, &instr.RightOperands); err != nil {
			return err
		}
	}
	return nil
}

func parseAndAddOperandSide(operands []string, operandList *[]Operand) error {
	for i := range operands {
		parsedOperand, err := parseOperand(operands[i])
		if err != nil {
			return err
		}
		*operandList = append(*operandList, parsedOperand)
	}
	return nil
}

func parseOperand(operandStr string) (Operand, error) {
	operandStr = strings.ToLower(operandStr)

	if operand, err := tryParseOperand(operandStr, addressRegex, OperandAddress); err == nil {
		return operand, nil
	}
	if operand, err := tryParseOperand(operandStr, valueRegex, OperandValue); err == nil {
		return operand, nil
	}
	if operand, err := tryParseOperand(operandStr, registerRegex, OperandRegister); err == nil {
		return operand, nil
	}
	return Operand{}, errors.New("unable to parse operand")
}

func tryParseOperand(operandStr string, parseRegex *regexp.Regexp, operandType OperandType) (Operand, error) {
	var operand Operand
	var result = parseRegex.FindStringSubmatch(operandStr)
	if len(result) < 2 {
		return operand, errors.New("unable to parse operand as current type")
	}
	operand.Type = operandType
	setOperandValue(&operand, result[1])
	return operand, nil
}

func setOperandValue(operand *Operand, value string) {
	var valueInt, _ = strconv.ParseInt(value, 16, 0)
	operand.Value = byte(valueInt)
}
