package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(path string) []string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func main() {

	input := getInput("input.txt")
	// part1
	// targetRedCubes := 12
	// targetGreenCubes := 13
	// targetBlueCubes := 14
	sumOfPossibleGames := 0

	for _, game := range input {
		rounds := strings.Split(strings.Split(game, ":")[1], ";")
		maxCubeColor := make(map[string]int)
		for _, round := range rounds {
			for _, colorInfo := range strings.Split(round, ",") {
				colorInfo = strings.Trim(colorInfo, " ")
				colorName := strings.Split(colorInfo, " ")[1]
				nbOfCube, err := strconv.Atoi(strings.Split(colorInfo, " ")[0])

				if err != nil {
					panic(err)
				}

				if maxCubeColor[colorName] < nbOfCube {
					maxCubeColor[colorName] = nbOfCube
				}
			}
		}

		// part 1
		// if maxCubeColor["red"] <= targetRedCubes && maxCubeColor["green"] <= targetGreenCubes && maxCubeColor["blue"] <= targetBlueCubes {
		// 	fmt.Println(maxCubeColor)
		// 	fmt.Println(index + 1)
		// 	sumOfPossibleGames += index + 1
		// }

		gamePower := 1
		for _, amount := range maxCubeColor {
			if amount > 0 {
				gamePower *= amount
			}
		}
		sumOfPossibleGames += gamePower
	}
	fmt.Println(sumOfPossibleGames)
}
