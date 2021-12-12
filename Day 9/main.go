package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func partOne(heightmap [][]uint8) {
	riskLevel := 0
	for y, _ := range heightmap {
		for x, _ := range heightmap[y] {
			if (y-1 >= 0 && heightmap[y][x] >= heightmap[y-1][x]) ||
				(y+1 <= len(heightmap)-1 && heightmap[y][x] >= heightmap[y+1][x]) ||
				(x-1 >= 0 && heightmap[y][x] >= heightmap[y][x-1]) ||
				(x+1 <= len(heightmap[y])-1 && heightmap[y][x] >= heightmap[y][x+1]) {
				continue
			}

			riskLevel += int(heightmap[y][x]) + 1
		}
	}

	fmt.Printf("Part One Answer: %d\n", riskLevel)
}

func getBasinSize(heightmap [][]uint8, basinMap [][]bool, x int, y int) int {
	basinMap[y][x] = true
	basinSize := 1

	if x-1 >= 0 && heightmap[y][x-1] != 9 && !basinMap[y][x-1] {
		basinSize += getBasinSize(heightmap, basinMap, x-1, y)
	}

	if x+1 <= len(heightmap[y]) - 1 && heightmap[y][x+1] != 9 && !basinMap[y][x+1] {
		basinSize += getBasinSize(heightmap, basinMap, x+1, y)
	}

	if y-1 >= 0 && heightmap[y-1][x] != 9 && !basinMap[y-1][x] {
		basinSize += getBasinSize(heightmap, basinMap, x, y-1)
	}

	if y+1 <= len(heightmap) - 1 && heightmap[y+1][x] != 9 && !basinMap[y+1][x] {
		basinSize += getBasinSize(heightmap, basinMap, x, y+1)
	}

	return basinSize
}

func partTwo(heightmap [][]uint8) {
	basinMap := make([][]bool, len(heightmap))
	for basinIndex, _ := range basinMap {
		basinMap[basinIndex] = make([]bool, len(heightmap[0]))
	}

	basins := []int{}
	for y, _ := range heightmap {
		for x, _ := range heightmap[y] {
			if heightmap[y][x] == 9 || basinMap[y][x] {
				continue
			}

			basinSize := getBasinSize(heightmap, basinMap, x, y)

			basins = append(basins, basinSize)
		}
	}

	sort.Ints(basins)
	
	fmt.Printf("Part Two Answer: %d\n", basins[len(basins) - 1] * basins[len(basins) - 2] * basins[len(basins) - 3])
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("[Error] Unable to open file")
		os.Exit(1)
	}
	defer file.Close()

	heightmap := [][]uint8{}
	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowString := scanner.Text()
		if len(rowString) == 0 {
			continue
		}

		heightmap = append(heightmap, make([]uint8, len(rowString)))
		for x, heightCharacter := range rowString {
			height, err := strconv.Atoi(string(heightCharacter))
			if err != nil {
				fmt.Printf("[Error] Invalid height")
				continue
			}
			heightmap[y][x] = uint8(height)
		}

		y += 1
	}

	partOne(heightmap)
	partTwo(heightmap)
}
