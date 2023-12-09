package utils

import (
	"os"
	"strings"
)

func ReadInput(testFile string) []string {
	data, err := os.ReadFile(testFile)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}
