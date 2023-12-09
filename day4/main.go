package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"utils"
)

type Scratchcard struct {
	Id             int
	WinningNumbers []string
	Numbers        []string
}

func (card Scratchcard) displayScratchcards() {
	fmt.Print(strconv.Itoa(card.Id) + " ")
	fmt.Print(card.WinningNumbers)
	fmt.Print(" ")
	fmt.Println(card.Numbers)
}

func (card Scratchcard) isNumberAWiningOne(number string) bool {
	for _, winningNumber := range card.WinningNumbers {
		if winningNumber == number {
			return true
		}
	}
	return false
}

func (card Scratchcard) getNumbersOfWinningCards() int {
	nbNumbersThatWins := 0
	for _, number := range card.Numbers {
		if card.isNumberAWiningOne(number) {
			nbNumbersThatWins += 1
		}
	}
	return nbNumbersThatWins
}

func (card Scratchcard) getPoints() int {
	return int(math.Pow(2, float64(card.getNumbersOfWinningCards()-1)))
}

func getScratchcards(inputFile string) []Scratchcard {
	input := utils.ReadInput(inputFile)

	var scratchcards []Scratchcard

	for cardIndex, card := range input {
		cardInfo := strings.Split(card, ":")[1] // we retrieve only the numbers part as the name and the index of the cards are not usefull information.
		numbers := strings.Split(cardInfo, "|")
		scratchcards = append(scratchcards, Scratchcard{
			Id: cardIndex,
			// We trim to remove whitespace from the start and end of the string.
			// Then we replace the double space that can happen if a number is a single digit, then we split it.
			WinningNumbers: strings.Split(strings.Replace(strings.Trim(numbers[0], " "), "  ", " ", -1), " "),
			Numbers:        strings.Split(strings.Replace(strings.Trim(numbers[1], " "), "  ", " ", -1), " "),
		})
	}

	return scratchcards
}

func getTotalPointsFromScratchCards(scratchcards []Scratchcard) int {
	ourPoints := 0
	for _, scratchCard := range scratchcards {
		pts := scratchCard.getPoints()
		ourPoints += pts
	}
	return ourPoints
}

func part1(inputFile string) {
	scratchcards := getScratchcards(inputFile)

	fmt.Println("Sum part 1 : " + strconv.Itoa(getTotalPointsFromScratchCards(scratchcards)))
}

func part2(inputFile string) {
	scratchcards := getScratchcards(inputFile)
	copies := make(map[int]int, 0)

	for _, scratchcard := range scratchcards {
		copies[scratchcard.Id] += 1

		for j := 0; j < copies[scratchcard.Id]; j++ {
			for i := 0; i < scratchcard.getNumbersOfWinningCards(); i++ {
				copies[scratchcard.Id+1+i] += 1
			}
		}
	}

	totalSum := 0
	for index := range scratchcards {
		totalSum += copies[index]
	}
	fmt.Println("Sum part 2 : " + strconv.Itoa(totalSum))
}

func main() {
	inputFile := "input.txt"
	part1(inputFile)
	part2(inputFile)
}
