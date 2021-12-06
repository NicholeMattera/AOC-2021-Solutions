package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne(lanternfishes []uint8, numberOfDays int) {
	for i := 0; i < numberOfDays; i++ {
		newLanternFishes := []uint8{}

		for lanternfishIndex := range lanternfishes {
			if lanternfishes[lanternfishIndex] == 0 {
				lanternfishes[lanternfishIndex] = 6
				newLanternFishes = append(newLanternFishes, uint8(8))
			} else {
				lanternfishes[lanternfishIndex] -= 1
			}
		}

		lanternfishes = append(lanternfishes, newLanternFishes[:]...)
	}

	fmt.Printf("Part One Answer: %d\n", len(lanternfishes))
}

func partTwo(lanternfishes []uint8, numberOfDays int) {
	days := [9]uint64{}

	for _, lanternfish := range lanternfishes {
		days[lanternfish] += 1
	}

	totalFish := uint64(0)
	for i := 0; i < numberOfDays; i++ {
		newDays := [9]uint64{}

		for day, numberOfFish := range days {
			if i == numberOfDays - 1 {
				if day == 0 {
					totalFish += numberOfFish * 2
				} else {
					totalFish += numberOfFish
				}
			} else {
				if day == 0 {
					newDays[8] += numberOfFish
					newDays[6] += numberOfFish
				} else {
					newDays[day - 1] += numberOfFish
				}
			}
		}

		days = newDays
	}

	fmt.Printf("Part Two Answer: %d\n", totalFish)
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("[Error] Unable to open file")
		os.Exit(1)
	}
	defer file.Close()

	partOneLanternfishes := []uint8{}
	partTwoLanternfishes := []uint8{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lanternfishesStrings := strings.Split(scanner.Text(), ",")
	if len(lanternfishesStrings) <= 1 {
		fmt.Printf("[Error] Invalid lanternfish list parsed\n")
	}

	for _, lanternfishesString := range lanternfishesStrings {
		lanternfish, err := strconv.Atoi(lanternfishesString)
		if err != nil {
			fmt.Printf("[Error] Invalid lanternfish parsed\n")
			continue
		}

		partOneLanternfishes = append(partOneLanternfishes, uint8(lanternfish))
		partTwoLanternfishes = append(partTwoLanternfishes, uint8(lanternfish))
	}

	partOne(partOneLanternfishes, 80)
	partTwo(partTwoLanternfishes, 256)
}
