package io

import (
	"brookshear-vm/pkgs/vm"
	"os"
)

func Compile(instrStr []string, dstFilePath string) error {
	instrs, err := vm.ParseInstructions(instrStr)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(dstFilePath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, instr := range instrs {
		var bytes = []byte{vm.CombineNibblesToByte(instr.Nibbles[0], instr.Nibbles[1]),
			vm.CombineNibblesToByte(instr.Nibbles[2], instr.Nibbles[3])}
		file.Write(bytes)
	}
	return nil
}
