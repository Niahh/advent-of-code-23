package main

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"utils"
)

type hand struct {
	cards string
	bid   int
	aType int
	rank  int
}

var rankingDay1 = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

var rankingDay2 = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

func getHandsFromInput(input string) []*hand {
	var hands []*hand

	for _, line := range utils.ReadInput(input) {
		res := strings.Split(line, " ")
		cards := res[0]
		bid, err := strconv.Atoi(res[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, &hand{cards, bid, 0, 0})
	}

	return hands
}

func isCardsFiveOfAKind(occurences map[rune]int) bool {
	return len(reflect.ValueOf(occurences).MapKeys()) == 1
}

func isCardsFourOfAKind(occurences map[rune]int) bool {
	keys := reflect.ValueOf(occurences).MapKeys()
	if len(keys) == 2 {
		for _, v := range occurences {
			if v == 4 {
				return true
			}
		}

	}
	return false
}

func isCardsFullHouse(occurences map[rune]int) bool {
	keys := reflect.ValueOf(occurences).MapKeys()
	if len(keys) == 2 {
		for _, v := range occurences {
			if v == 3 {
				return true
			}
		}
	}
	return false
}

func isThreeOfAKind(occurences map[rune]int) bool {
	keys := reflect.ValueOf(occurences).MapKeys()
	if len(keys) == 3 {
		for _, v := range occurences {
			if v == 3 {
				return true
			}
		}
	}
	return false
}

func isTwoPair(occurences map[rune]int) bool {
	keys := reflect.ValueOf(occurences).MapKeys()
	if len(keys) == 3 {
		soloElement := 0
		for _, occurence := range occurences {
			if occurence == 1 {
				soloElement++
			}
		}
		return soloElement == 1
	}
	return false
}

func isOnePair(occurences map[rune]int) bool {
	keys := reflect.ValueOf(occurences).MapKeys()
	if len(keys) == 4 {
		soloPair := 0
		for _, occurence := range occurences {
			if occurence == 2 {
				soloPair++
			}
		}

		return soloPair == 1
	}
	return false
}

func getMaxInMap(inputMap map[rune]int) (int, rune) {
	max := 0
	var topCard rune
	for card, elem := range inputMap {
		if elem > max && card != 'J' {
			max = elem
			topCard = card
		}
	}
	return max, topCard
}

func getOccurences(cards string, part int) map[rune]int {
	occurences := make(map[rune]int)

	for _, card := range cards {
		occurences[card] += 1
	}

	if part == 2 {
		if wildcards, ok := occurences['J']; ok && wildcards != 5 {
			max, card := getMaxInMap(occurences)
			for _, elem := range occurences {
				if elem == max {
					occurences[card] += wildcards
					break
				}
			}
			delete(occurences, 'J')
		}
	}

	return occurences
}

func getHandType(hand *hand, part int) int {
	occurences := getOccurences(hand.cards, part)
	if isCardsFiveOfAKind(occurences) {
		return 6
	} else if isCardsFourOfAKind(occurences) {
		return 5
	} else if isCardsFullHouse(occurences) {
		return 4
	} else if isThreeOfAKind(occurences) {
		return 3
	} else if isTwoPair(occurences) {
		return 2
	} else if isOnePair(occurences) {
		return 1
	}
	return 0
}

func getStringIndexInSlice(slice []string, str string) (int, error) {
	for index, elem := range slice {
		if elem == str {
			return index, nil
		}
	}
	return -1, errors.New("string is not in the provided slice")
}

func getBetterHandSameType(hand1 hand, hand2 hand, part int) int {
	var rankingSystem = []string{}
	if part == 1 {
		rankingSystem = rankingDay1
	} else {
		rankingSystem = rankingDay2
	}
	for index, card := range hand1.cards {
		indexHand1, err := getStringIndexInSlice(rankingSystem, string(card))
		if err != nil {
			panic("nope")
		}
		indexHand2, err := getStringIndexInSlice(rankingSystem, string(hand2.cards[index]))
		if indexHand1 > indexHand2 {
			return 1
		} else if indexHand1 < indexHand2 {
			return -1
		}
	}
	return 0
}

func compareHands(hand1 hand, hand2 hand, part int) int {
	if hand1.aType == hand2.aType {
		return getBetterHandSameType(hand1, hand2, part)
	} else if hand1.aType > hand2.aType {
		return 1
	} else if hand1.aType < hand2.aType {
		return -1
	}
	return 0
}

func part1(hands []*hand) {
	for _, hand := range hands {
		hand.aType = getHandType(hand, 1)
	}
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(*hands[i], *hands[j], 1) == -1
	})

	sum := 0

	for index, elem := range hands {
		sum += elem.bid * (index + 1)
	}

	fmt.Println("Sum : " + strconv.Itoa(sum))
}

func part2(hands []*hand) {
	for _, hand := range hands {
		hand.aType = getHandType(hand, 2)
	}
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(*hands[i], *hands[j], 2) == -1
	})

	sum := 0

	for index, elem := range hands {
		sum += elem.bid * (index + 1)
	}

	fmt.Println("Sum : " + strconv.Itoa(sum))
}

func main() {
	inputFileName := "input.txt"
	hands := getHandsFromInput(inputFileName)
	// part1(hands)
	part2(hands)
}
