package compiler

import (
	"fmt"
	"brookshear-vm/vm"
	"os"
)

func Compile(instrStr []string, srcFilePath string) error {
	var dstFilePath = srcFilePath + ".bin"
	fmt.Println("Compiling to file:", dstFilePath)
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
			vm.CombineNibblesToByte(instr.Nibbles[0], instr.Nibbles[1])}
		file.Write(bytes)
	}
	fmt.Println("Done!")
	return nil
}
