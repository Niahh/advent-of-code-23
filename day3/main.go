package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput(testFile string) []string {
	data, err := os.ReadFile(testFile)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func parseInput(testFile string) [][][][]int {
	var parsedInput [][][][]int

	var patternNumbers string = "([0-9]+)"
	numbersRegex := regexp.MustCompile(patternNumbers)

	var patternSymbols string = "[^\\d\\.]"
	symbolRegex := regexp.MustCompile(patternSymbols)

	for _, line := range readInput(testFile) {
		// Find the index of the numbers
		numbersMatches := numbersRegex.FindAllStringSubmatchIndex(line, -1)
		// find the indexes of the symbols
		symbolMatches := symbolRegex.FindAllStringSubmatchIndex(line, -1)
		parsedInput = append(parsedInput, [][][]int{numbersMatches, symbolMatches})
	}
	return parsedInput
}

func getNumbersFromInput(testFile string) [][][]string {
	var patternNumbers string = `([0-9]+)`
	numbersRegex := regexp.MustCompile(patternNumbers)

	var parsedNumbers [][][]string
	for _, line := range readInput(testFile) {
		numbersMatches := numbersRegex.FindAllStringSubmatch(line, -1)
		parsedNumbers = append(parsedNumbers, numbersMatches)
	}

	return parsedNumbers
}

func getSymbolsLitterals(testFile string) map[string]int {
	var patternSymbols string = "[=%@&\\*\\#\\+\\$\\-\\/]"
	symbolRegex := regexp.MustCompile(patternSymbols)

	var parsedSymbol = make(map[string]int)

	for _, line := range readInput(testFile) {
		symbolsMatches := symbolRegex.FindAllStringSubmatch(line, -1)
		if len(symbolsMatches) > 0 {
			parsedSymbol[symbolsMatches[0][0]] += 1
		}
	}

	return parsedSymbol
}

func part1(testFile string) int {

	data := parseInput(testFile)
	sum := 0

	numbers := getNumbersFromInput(testFile)

	for index, row := range data {
		for indNum, positions := range row[0] {
			alreadyPicked := false
			// check top row
			if index > 0 && len(data[index-1][1]) > 0 && !alreadyPicked {
				for _, symbIndexData := range data[index-1][1] {
					if symbIndexData[0] >= positions[0]-1 && symbIndexData[1] <= positions[1]+1 {
						tmp, err := strconv.Atoi(numbers[index][indNum][0])
						if err != nil {
							panic(err)
						}
						sum += tmp
						alreadyPicked = true
						break
					}
				}
			}
			// check bottom row
			if index < len(parseInput(testFile))-1 && !alreadyPicked {
				for _, symbIndexData := range data[index+1][1] {
					if symbIndexData[0] >= positions[0]-1 && symbIndexData[1] <= positions[1]+1 {
						tmp, err := strconv.Atoi(numbers[index][indNum][0])
						if err != nil {
							panic(err)
						}
						sum += tmp
						alreadyPicked = true
						break
					}
				}
			}

			// check left
			if positions[0] > 0 && !alreadyPicked {
				for _, symbIndexData := range row[1] {
					if symbIndexData[0] == positions[0]-1 {
						tmp, err := strconv.Atoi(numbers[index][indNum][0])
						if err != nil {
							panic(err)
						}
						sum += tmp
						alreadyPicked = true
						break
					}
				}
			}

			// check right
			if !alreadyPicked {
				for _, symbIndexData := range row[1] {
					if symbIndexData[0] == positions[1] {
						tmp, err := strconv.Atoi(numbers[index][indNum][0])
						if err != nil {
							panic(err)
						}
						sum += tmp
						alreadyPicked = true
						break
					}
				}
			}
		}

	}
	return sum
}

func main() {
	testFile := "input.txt"
	// testFileTestBelow := "tests/below.txt"
	// testFileTestUp := "tests/up.txt"
	// testFileTestLeft := "tests/left.txt"
	// testFileTestRight := "tests/right.txt"
	// testFileTestSymbols := "tests/symbols.txt"
	fmt.Println(part1(testFile))
	// fmt.Println(part1(testFileTestBelow))   // must be equals to 666
	// fmt.Println(part1(testFileTestUp))      // must be equals to 555
	// fmt.Println(part1(testFileTestLeft))    // must be equals to 111
	// fmt.Println(part1(testFileTestRight)) // must be equals to 111
	// fmt.Println(part1(testFileTestSymbols)) // must be equals to 1110
}
