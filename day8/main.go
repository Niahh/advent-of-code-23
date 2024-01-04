package main

import (
	"fmt"
	"strings"
	"utils"
)

type desertMap struct {
	itinerary string
	rows      []row
}

func (dm desertMap) String() string {
	strToDisplay := ""
	strToDisplay += dm.itinerary + "\n"
	strToDisplay += "\n"
	for _, row := range dm.rows {
		strToDisplay += row.startingNode + " = (" + row.leftNode + ", " + row.rightNode + ")\n"
	}
	return strToDisplay
}

type row struct {
	startingNode string
	leftNode     string
	rightNode    string
}

func parseMaps(input string) desertMap {
	file := utils.ReadInput(input)
	itinerary := file[0]
	var mapsRows = []row{}
	for _, line := range file[2:] {
		nodes := strings.Split(line, "=")
		outNodes := strings.Split(strings.TrimFunc(strings.Replace(nodes[1], " ", "", -1), func(r rune) bool {
			return r == '(' || r == ')'
		}), ",")
		mapsRows = append(mapsRows, row{nodes[0], outNodes[0], outNodes[1]})
	}
	return desertMap{itinerary, mapsRows}
}

func part1() {

}

func validateInput() {
	testInputStr := "test-input.txt"
	desertMap := parseMaps(testInputStr)
	fmt.Println(desertMap)
}

func main() {
	validateInput()
}
