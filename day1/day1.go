package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("test")
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	sum := 0

	rows := strings.Split(string(data), "\n")
	// rows = []string{"a1ocucousevenine", "sevend95"}
	for _, coor := range rows {
		var digits []int
		fmt.Print(coor)
		fmt.Print(" ")
		numStr := ""

		firstDigitsFound := false
		lastDigitFound := false

		for _, letter := range coor {

			if unicode.IsDigit(letter) {
				digits = append(digits, int(letter)-'0')
				firstDigitsFound = true
				break
			} else {
				numStr = numStr + string(letter)
				for i, n := range numbers {
					if strings.Contains(numStr, n) {
						digits = append(digits, i+1)
						numStr = ""
						firstDigitsFound = true
						break
					}
				}
			}
		}

		for j := len(coor) - 1; j >= 0; j-- {

			if unicode.IsDigit(rune(coor[j])) {
				digits = append(digits, int(rune(coor[j]))-'0')
				lastDigitFound = true
				break
			} else {
				numStr = string(rune(coor[j])) + numStr
				for i, n := range numbers {
					if strings.Contains(numStr, n) {
						digits = append(digits, i+1)
						numStr = ""
						lastDigitFound = true
						break
					}
				}
			}

		}

		fmt.Print(" ")
		fmt.Print(digits)
		fmt.Print(" ")

		if firstDigitsFound && lastDigitFound {
			pair := strconv.Itoa(digits[0]) + strconv.Itoa(digits[len(digits)-1])
			fmt.Print(pair)
			result, err := strconv.Atoi(pair)

			if err != nil {
				panic(err)
			}

			sum += result
		} else if firstDigitsFound {
			result, err := strconv.Atoi(strconv.Itoa(digits[0]) + strconv.Itoa(digits[0]))
			fmt.Print(result)

			if err != nil {
				panic(err)
			}

			sum += result
		}
		println(" ")
	}
	println(sum)
}
