package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partOne(crabPositions []int) {
	sort.Ints(crabPositions)
	alignTo := crabPositions[int(len(crabPositions) / 2)]
	fuel := 0
	for _, crabPosition := range crabPositions {
		fuel += int(math.Abs(float64(alignTo - crabPosition)))
	}

	fmt.Printf("Part One Answer: %d\n", fuel)
}

func partTwo(crabPositions []int) {
	sum := 0
	for _, crabPosition := range crabPositions {
		sum += crabPosition
	}

	alignTo := int(math.Round(float64(sum / len(crabPositions))))
	fuel := 0
	for _, crabPosition := range crabPositions {
		for i := 1; i <= int(math.Abs(float64(alignTo - crabPosition))); i++ {
			fuel += i
		}
	}

	fmt.Printf("Part Two Answer: %d\n", fuel)
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("[Error] Unable to open file")
		os.Exit(1)
	}
	defer file.Close()

	crabPositions := []int{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	crabPositionsStrings := strings.Split(scanner.Text(), ",")
	if len(crabPositionsStrings) <= 1 {
		fmt.Printf("[Error] Invalid crab position list parsed\n")
	}

	for _, crabPositionsString := range crabPositionsStrings {
		crabPosition, err := strconv.Atoi(crabPositionsString)
		if err != nil {
			fmt.Printf("[Error] Invalid crab position parsed\n")
			continue
		}

		crabPositions = append(crabPositions, crabPosition)
	}

	partOne(crabPositions)
	partTwo(crabPositions)
}
