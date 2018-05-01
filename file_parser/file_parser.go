package file_parser

import (
	"io/ioutil"
	"strings"
	"regexp"
)

var newLineRegex = regexp.MustCompile(`(\r\n|\r|\n)`)

func ParseFile(filePath string) ([]string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var output = newLineRegex.Split(string(bytes), -1)
	output = trimSpaces(output)
	output = removeCommentLines(output)
	return output, nil
}

func trimSpaces(input []string) []string {
	for i := range input {
		input[i] = strings.TrimSpace(input[i])
	}
	return input
}

func removeCommentLines(input []string) (output []string) {
	for _, item := range input {
		if strings.HasPrefix(item, "#") || strings.HasPrefix(item, "//") || item == "" {
			continue
		}
		output = append(output, item)
	}
	return
}
