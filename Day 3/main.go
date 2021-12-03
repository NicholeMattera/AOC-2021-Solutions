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

func getCommonBitValues(reports []int, bitPosition int) (int, int) {
	bitCounts := 0
	for _, report := range reports { 
		power := powInt(2, bitPosition)
		if power & report == power {
			bitCounts += 1
		}
	}

	mcv := 0
	lcv := 0

	// Most Common Value
	if bitCounts >= len(reports) / 2 {
		mcv = 1
	} else {
		mcv = 0
	}
	
	// Least Common Value
	if bitCounts >= len(reports) / 2 {
		lcv = 0
	} else {
		lcv = 1
	}

	return mcv, lcv
}

func filter(slice []int, test func(int, int, int) bool, bitPosition int, bitValue int) (results []int) {
	for _, element := range slice {
		if test(element, bitPosition, bitValue) {
			results = append(results, element)
		}
	}
	return
}

func partOne(reports []int, numberOfBits int) {
	gammaRate := 0
	epsilonRate := 0

	for i := 0; i < numberOfBits; i++ {
		if mcv, _ := getCommonBitValues(reports, i); mcv == 1 {
			gammaRate += powInt(2, i)
		} else {
			epsilonRate += powInt(2, i)
		}
	}

	fmt.Printf("Part One Answer: %d\n", gammaRate * epsilonRate)
}

func partTwo(reports []int, numberOfBits int) {
	ogrFilteredReports := reports
	co2srFilteredReports := reports

	filterTest := func (report, bitPosition, bitValue int) bool {
		power := powInt(2, bitPosition)
		return (bitValue == 1 && power & report == power) || (bitValue == 0 && power & report != power)
	}

	for i := numberOfBits - 1; i >= 0; i-- {
		mcv, _ := getCommonBitValues(ogrFilteredReports, i)

		if len(ogrFilteredReports) > 1 {
			ogrFilteredReports = filter(ogrFilteredReports, filterTest, i, mcv)
		}

		_, lcv := getCommonBitValues(co2srFilteredReports, i)

		if len(co2srFilteredReports) > 1 {
			co2srFilteredReports = filter(co2srFilteredReports, filterTest, i, lcv)
		}

		if len(ogrFilteredReports) <= 1 && len(co2srFilteredReports) <= 1 {
			break
		}
	}

	fmt.Printf("Part Two Answer: %d\n", ogrFilteredReports[0] * co2srFilteredReports[0])
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
