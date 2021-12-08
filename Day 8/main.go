package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

/* --- Start of UniqueSignalPattern type --- */

type UniqueSignalPatterns [10]SignalPattern

func (u UniqueSignalPatterns) Len() int { return len(u) }
func (u UniqueSignalPatterns) Swap(i, j int) { u[i], u[j] = u[j], u[i] }
func (u UniqueSignalPatterns) Less(i, j int) bool { return u[i].numberOfSignals < u[j].numberOfSignals }

func MakeUniqueSignalPatterns(data string) UniqueSignalPatterns {
	uniqueSignalPatterns := UniqueSignalPatterns{}

	uniqueSignalPatternsStrings := strings.Split(data, " ")
	for uniqueSignalPatternsIndex, uniqueSignalPatternsString := range uniqueSignalPatternsStrings {
		uniqueSignalPatterns[uniqueSignalPatternsIndex] = MakeSignalPattern([]byte(uniqueSignalPatternsString))
	}

	sort.Sort(uniqueSignalPatterns)

	return uniqueSignalPatterns
}

/* --- End of UniqueSignalPattern type --- */

/* --- Start of SignalPattern struct --- */

type SignalPattern struct {
	signals [7]bool
	numberOfSignals int
}

func MakeSignalPattern(data []byte) SignalPattern {
	signals := [7]bool{}
	for _, character := range data {
		signals[int(character) - 0x61] = true
	}
	return SignalPattern{signals,len(data)}
}

/* --- End of SignalPattern struct --- */

/* --- Start of InputLine struct --- */

type InputLine struct {
	key	map[int]int
	uniqueSignalPatterns UniqueSignalPatterns
	outputValues [4]SignalPattern
}

func MakeInputLine(data string) (InputLine, error) {
	inputLine := InputLine{}

	inputLineStrings := strings.Split(data, " | ")
	if len(inputLineStrings) != 2 {
		return inputLine, errors.New("Invalid input line parsed")
	}

	inputLine.uniqueSignalPatterns = MakeUniqueSignalPatterns(inputLineStrings[0])

	outputValuesStrings := strings.Split(inputLineStrings[1], " ")
	for outputValueIndex, outputValueString := range outputValuesStrings {
		inputLine.outputValues[outputValueIndex] = MakeSignalPattern([]byte(outputValueString))
	}

	return inputLine, nil
}

/* --- End of InputLine struct --- */

func partOne(inputLines []InputLine) {
	numberOfUniqueSegments := 0
	for _, inputLine := range inputLines {
		for _, outputValue := range inputLine.outputValues {
			if outputValue.numberOfSignals == 2 || outputValue.numberOfSignals == 3 || outputValue.numberOfSignals == 4 || outputValue.numberOfSignals == 7 {
				numberOfUniqueSegments++
			}
		}
	}
	fmt.Printf("Part One Answer: %d\n", numberOfUniqueSegments)
}

func partTwo(inputLines []InputLine) {
	fmt.Printf("Part Two Answer: %d\n", 0)
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("[Error] Unable to open file")
		os.Exit(1)
	}
	defer file.Close()

	inputLines := []InputLine{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine, err := MakeInputLine(scanner.Text())
		if err != nil {
			fmt.Printf("[Error] %s\n", err)
			continue
		}

		inputLines = append(inputLines, inputLine)
	}

	partOne(inputLines)
	partTwo(inputLines)
}
