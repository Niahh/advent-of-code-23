package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"utils"
)

type Race struct {
	Time     int
	Distance int
}

func (r Race) displayRace() {
	fmt.Println("Time : " + strconv.Itoa(r.Time) + " - Distance : " + strconv.Itoa(r.Distance))
}

func (r Race) calculateWinPossibilities() int {
	nbPossibilitiesToWin := 0
	for i := 1; i < r.Time; i++ {
		timeRemaining := r.Time - i
		if timeRemaining*i > r.Distance {
			nbPossibilitiesToWin += 1
		}

	}

	return nbPossibilitiesToWin
}

func parseInput(input string, part int) []Race {
	var races []Race

	inputLines := utils.ReadInput(input)

	var patternNumbers string = "[0-9]+"
	numbersRegex := regexp.MustCompile(patternNumbers)

	times := numbersRegex.FindAll([]byte(inputLines[0]), -1)
	distances := numbersRegex.FindAll([]byte(inputLines[1]), -1)

	if part == 1 {
		for i := 0; i < len(times); i++ {
			time, err := strconv.Atoi(string(times[i]))
			if err != nil {
				panic("Error while parsing a time value - Part 1")
			}

			distance, err := strconv.Atoi(string(distances[i]))
			if err != nil {
				panic("Error while parsing a distance value - Part 1")
			}

			races = append(races, Race{
				Time:     time,
				Distance: distance,
			})
		}
	} else {
		time, err := strconv.Atoi(strings.ReplaceAll(string(bytes.Join(times, []byte{'|'})), "|", ""))
		if err != nil {
			panic("Error while parsing a time value - Part 2")
		}

		distance, err := strconv.Atoi(strings.ReplaceAll(string(bytes.Join(distances, []byte{'|'})), "|", ""))
		if err != nil {
			panic("Error while parsing a distance value - Part 2")
		}
		races = append(races, Race{
			Time:     time,
			Distance: distance,
		})
	}

	return races
}

func calculateAnswer(result []int) int {
	answer := 1
	for _, waysToWin := range result {
		answer *= waysToWin
	}
	return answer
}

func part1(input string) {
	races := parseInput(input, 1)
	var results []int
	for _, race := range races {
		results = append(results, race.calculateWinPossibilities())
	}

	fmt.Println(calculateAnswer(results))
}

func part2(input string) {
	race := parseInput(input, 2)[0]

	fmt.Println(calculateAnswer([]int{race.calculateWinPossibilities()}))
}

func main() {
	input := "input.txt"
	part1(input)
	part2(input)
}
