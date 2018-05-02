package io

import (
	"io/ioutil"
	"regexp"
)

var newLineRegex = regexp.MustCompile(`(\r\n|\r|\n)`)

func ReadAsmFile(filePath string) ([]string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var output = newLineRegex.Split(string(bytes), -1)
	return output, nil
}