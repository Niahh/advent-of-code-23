package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Indexes struct {
	Numbers [][]int
	Symbols [][]int
}

func readInput(testFile string) []string {
	data, err := os.ReadFile(testFile)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func parseInput(testFile string, part int) []Indexes {
	var parsedInput []Indexes

	var patternNumbers string = "[0-9]+"
	numbersRegex := regexp.MustCompile(patternNumbers)

	var patternSymbols string

	if part == 1 {
		// Get everything that is not a digit or a dot
		patternSymbols = "[^\\d\\.]"
	} else if part == 2 {
		// Need to spot only the stars
		patternSymbols = "[\\*]"
	} else {
		panic("Please enter the correct number of the part you want to solve")
	}
	symbolRegex := regexp.MustCompile(patternSymbols)

	for _, line := range readInput(testFile) {
		// Find the index of the numbers
		numbersMatches := numbersRegex.FindAllStringSubmatchIndex(line, -1)
		// find the indexes of the symbols
		symbolMatches := symbolRegex.FindAllStringSubmatchIndex(line, -1)
		parsedInput = append(parsedInput, Indexes{numbersMatches, symbolMatches})
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

func part1(testFile string) int {

	data := parseInput(testFile, 1)
	sum := 0

	numbers := getNumbersFromInput(testFile)

	for index, row := range data {
		for indNum, positions := range row.Numbers {
			alreadyPicked := false
			// check top row
			if index > 0 && len(data[index-1].Symbols) > 0 && !alreadyPicked {
				for _, symbIndexData := range data[index-1].Symbols {
					if symbIndexData[0] >= positions[0]-1 && symbIndexData[1] <= positions[1]+1 {
						tmp, err := strconv.Atoi(numbers[index][indNum][0])
						if err != nil {
							panic(err)
						}
						sum += tmp
						alreadyPicked = true
					}
				}
			}
			// check bottom row
			if index < len(data)-1 && !alreadyPicked {
				for _, symbIndexData := range data[index+1].Symbols {
					if symbIndexData[0] >= positions[0]-1 && symbIndexData[1] <= positions[1]+1 {
						tmp, err := strconv.Atoi(numbers[index][indNum][0])
						if err != nil {
							panic(err)
						}
						sum += tmp
						alreadyPicked = true
					}
				}
			}

			// check left
			if positions[0] > 0 && !alreadyPicked {
				for _, symbIndexData := range row.Symbols {
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
				for _, symbIndexData := range row.Symbols {
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

func part2(testFile string) int {
	data := parseInput(testFile, 2)
	sum := 0

	numbers := getNumbersFromInput(testFile)

	for index, row := range data {
		for _, positions := range row.Symbols {

			var numbersToMultiply []int

			// check top row
			if index > 0 && len(data[index-1].Numbers) > 0 {
				for loopedIndex, numbIndexData := range data[index-1].Numbers {
					tmpString := numbers[index-1][loopedIndex][0]
					tmp, err := strconv.Atoi(tmpString)
					if err != nil {
						panic(err)
					}
					if numbIndexData[0] >= (positions[0]-(len(tmpString))) && numbIndexData[0] <= positions[0]+1 {

						numbersToMultiply = append(numbersToMultiply, tmp)
					}
				}
			}
			// check bottom row
			if index < len(data)-1 {
				for loopedIndex, numbIndexData := range data[index+1].Numbers {
					tmpString := numbers[index+1][loopedIndex][0]
					tmp, err := strconv.Atoi(numbers[index+1][loopedIndex][0])
					if err != nil {
						panic(err)
					}
					if numbIndexData[0] >= (positions[0]-(len(tmpString))) && numbIndexData[0] <= positions[0]+1 {

						numbersToMultiply = append(numbersToMultiply, tmp)
					}
				}
			}

			// check left
			if positions[0] > 0 {
				for loopedIndex, numbIndexData := range row.Numbers {
					if numbIndexData[1] == positions[0] {
						tmp, err := strconv.Atoi(numbers[index][loopedIndex][0])
						if err != nil {
							panic(err)
						}
						numbersToMultiply = append(numbersToMultiply, tmp)
						break
					}
				}
			}

			// check right

			for loopedIndex, numbIndexData := range row.Numbers {
				if numbIndexData[0] == positions[1] {
					tmp, err := strconv.Atoi(numbers[index][loopedIndex][0])
					if err != nil {
						panic(err)
					}
					numbersToMultiply = append(numbersToMultiply, tmp)
					break
				}
			}

			if len(numbersToMultiply) == 2 {
				sum += numbersToMultiply[0] * numbersToMultiply[1]
			}
		}

	}
	return sum
}

func main() {
	testFile := "input.txt"
	fmt.Println(part1(testFile))
	fmt.Println(part2(testFile))
}
