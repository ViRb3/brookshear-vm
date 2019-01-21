package io

import (
	"os"
	"github.com/pkg/errors"
	"brookshear-vm/pkgs/vm"
	"io"
)


func Decompile(srcFilePath string) (instrStr []*vm.Instruction, err error) {
	file, err := os.OpenFile(srcFilePath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//TODO: Optimize buffer length
	var buffer = make([]byte, 2)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if n != len(buffer) {
			return nil, errors.New("invalid instruction length")
		}
		var nibbles = vm.ByteArrayToNibbleArray(buffer)
		instruction, err := vm.ParseFromNibbles(nibbles)
		if err != nil {
			return nil, err
		}
		instrStr = append(instrStr, instruction)
	}
	return instrStr, nil
}
