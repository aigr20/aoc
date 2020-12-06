package util

import (
	"io/ioutil"
	"strings"
)

// ReadInput takes the path to an input file and returns the content
func ReadInput(filepath string) (input string, err error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	input = strings.TrimSuffix(string(file), "\n")

	return
}
