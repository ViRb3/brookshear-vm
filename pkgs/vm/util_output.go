package vm

import "fmt"

func PrintNibbles(nibbles []Nibble) (output string) {
	for _, nibble := range nibbles {
		output += fmt.Sprintf("%X", nibble)
	}
	return output
}

func PrintByteVerboseIfNonZero(val byte) {
	if val == 0 {
		fmt.Printf("%2X ", val)
	} else {
		fmt.Printf("%02X ", val)
	}
}
