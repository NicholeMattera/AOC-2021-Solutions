package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func filter(slice []int, test func(int, int) bool, bitPosition int) (results []int) {
	for _, element := range slice {
		if test(element, bitPosition) {
			results = append(results, element)
		}
	}
	return
}

func partOne(reports []int, numberOfBits int) {
	bitCounts := []int {}
	gammaRate := 0
	epsilonRate := 0

	for i := 0; i < numberOfBits; i++ {
		for _, report := range reports { 
			if len(bitCounts) <= i {
				bitCounts = append(bitCounts, 0)
			}

			power := powInt(2, i)
			if power & report == power {
				bitCounts[i] += 1
			}
		}

		if bitCounts[i] > len(reports) / 2 {
			gammaRate += powInt(2, i)
		} else {
			epsilonRate += powInt(2, i)
		}
	}

	fmt.Printf("Part One Answer: %d\n", gammaRate * epsilonRate)
}

func partTwo(reports []int, numberOfBits int) {
	fmt.Printf("Part Two Answer: %d\n", 0)
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("[Error] Unable to open file")
		os.Exit(1)
	}
	defer file.Close()

	reports := []int {}
	numberOfBits := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report, err := strconv.ParseInt(scanner.Text(), 2, 16)
		if err != nil {
			fmt.Printf("[Error] Invalid binary number parsed")
			continue
		}

		reports = append(reports, int(report))
		numberOfBits = len(scanner.Text())
	}

	partOne(reports, numberOfBits)
	partTwo(reports, numberOfBits)
}
