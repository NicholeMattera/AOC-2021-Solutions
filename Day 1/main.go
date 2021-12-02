package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partOne(depths []int) {
	previousDepth := 0
	increases := 0

	for i, depth := range depths {
		if i > 0 && depth > previousDepth {
			increases += 1
		}

		previousDepth = depth
	}

	fmt.Printf("Part One Answer: %d\n", increases)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func partTwo(depths []int) {
	slidingDepths := []int {}
	increases := 0

	for i, depth := range depths {
		slidingDepths = append(slidingDepths, depth)
		for slide := i - 1; slide >= max(i - 2, 0); slide-- {
			slidingDepths[slide] += depth
		}

		if i > 2 && slidingDepths[i - 2] > slidingDepths[i - 3] {
			increases += 1
		}
	}

	fmt.Printf("Part Two Answer: %d\n", increases)
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Println("[Error] Unable to open file")
		os.Exit(1)
	}
	defer file.Close()

	depths := []int {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("[Error] Unable to parse depth")
			continue
		}

		depths = append(depths, depth)
	}

	partOne(depths);
	partTwo(depths);
}
