package vm

import (
	"testing"
	"fmt"
	"github.com/pkg/errors"
)

func TestBytesToInstr(t *testing.T) {
	testInstrBytes([]Nibble{0x0, 0x0, 0x0, 0x0}, t)
	testInstrBytes([]Nibble{0x1, 0x2, 0x3, 0x4}, t)
	testInstrBytes([]Nibble{0x2, 0x2, 0x3, 0x4}, t)
	testInstrBytes([]Nibble{0x3, 0x2, 0x3, 0x4}, t)
	testInstrBytes([]Nibble{0x4, 0x0, 0x3, 0x4}, t)
	testInstrBytes([]Nibble{0x5, 0x2, 0x3, 0x4}, t)
	testInstrBytes([]Nibble{0x6, 0x2, 0x3, 0x4}, t)
	testInstrBytes([]Nibble{0x7, 0x2, 0x3, 0x4}, t)
	testInstrBytes([]Nibble{0x8, 0x2, 0x3, 0x4}, t)
	testInstrBytes([]Nibble{0x9, 0x2, 0x3, 0x4}, t)
	testInstrBytes([]Nibble{0xa, 0x2, 0x0, 0x4}, t)
	testInstrBytes([]Nibble{0xb, 0x2, 0x3, 0x4}, t)
	testInstrBytes([]Nibble{0xc, 0x0, 0x0, 0x0}, t)
}

func TestBytesToInstrFail(t *testing.T) {
	testInstrBytesFail([]Nibble{0x0, 0x2, 0x3, 0x4}, t)
	testInstrBytesFail([]Nibble{0x4, 0x2, 0x3, 0x4}, t)
	testInstrBytesFail([]Nibble{0xa, 0x2, 0x3, 0x4}, t)
	testInstrBytesFail([]Nibble{0xc, 0x2, 0x3, 0x4}, t)
}

// must fail, since some bytes must not match
func testInstrBytesFail(instrBytes []Nibble, t *testing.T) {
	instr, err := ParseFromNibbles(instrBytes)
	if err != nil {
		t.Errorf("%+v", err)
	}
	var perfectMatch = true
	for i := range instrBytes {
		if instrBytes[i] != instr.Nibbles[i] {
			perfectMatch = false
		}
	}
	if perfectMatch {
		t.Errorf("%+v", errors.New(fmt.Sprintf("bytes match")))
	}
}

func testInstrBytes(instrBytes []Nibble, t *testing.T) {
	instr, err := ParseFromNibbles(instrBytes)
	if err != nil {
		t.Errorf("%+v", err)
	}
	for i := range instrBytes {
		if instrBytes[i] != instr.Nibbles[i] {
			t.Errorf("%+v", errors.New(fmt.Sprintf("bytes don't match")))
		}
	}
}
