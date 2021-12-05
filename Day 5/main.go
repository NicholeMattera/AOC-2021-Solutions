package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type Vent struct {
	start Coordinate
	end Coordinate
}


func min(a, b int) int {
	if a < b { return a }
	return b
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func toCoordinate(coordinatesString string) (Coordinate, error) {
	coordinates := strings.Split(coordinatesString, ",")
	x, err := strconv.Atoi(coordinates[0])
	if err != nil { return Coordinate{}, err }

	y, err := strconv.Atoi(coordinates[1])
	if err != nil { return Coordinate{}, err }

	return Coordinate{x,y}, nil
}

func partOne(vents []Vent, size Coordinate) {
	ventCount := make([][]int, size.x + 1)
	for i := range ventCount {
		ventCount[i] = make([]int, size.y + 1)
	}

	dangerousVentCount := 0
	for _, vent := range vents {
		if vent.start.x != vent.end.x && vent.start.y != vent.end.y {
			continue
		}

		if vent.start.x == vent.end.x && vent.start.y == vent.end.y { // Point
			ventCount[vent.start.x][vent.start.y] += 1
			if ventCount[vent.start.x][vent.start.y] == 2 {
				dangerousVentCount += 1
			}
			continue
		} else if vent.start.x == vent.end.x { // Vertical Line
			for y := min(vent.start.y, vent.end.y); y <= max(vent.start.y, vent.end.y); y++ {
				ventCount[vent.start.x][y] += 1
				if ventCount[vent.start.x][y] == 2 {
					dangerousVentCount += 1
				}
			}
		} else if vent.start.y == vent.end.y { // Horizontal Line
			for x := min(vent.start.x, vent.end.x); x <= max(vent.start.x, vent.end.x); x++ {
				ventCount[x][vent.start.y] += 1
				if ventCount[x][vent.start.y] == 2 {
					dangerousVentCount += 1
				}
			}
		}
	}

	fmt.Printf("Part One Answer: %d\n", dangerousVentCount)
}

func partTwo(vents []Vent, size Coordinate) {
	ventCount := make([][]int, size.x + 1)
	for i := range ventCount {
		ventCount[i] = make([]int, size.y + 1)
	}

	dangerousVentCount := 0
	for _, vent := range vents {
		if vent.start.x == vent.end.x && vent.start.y == vent.end.y { // Point
			ventCount[vent.start.x][vent.start.y] += 1
			if ventCount[vent.start.x][vent.start.y] == 2 {
				dangerousVentCount += 1
			}
			continue
		} else if vent.start.x == vent.end.x { // Vertical Line
			for y := min(vent.start.y, vent.end.y); y <= max(vent.start.y, vent.end.y); y++ {
				ventCount[vent.start.x][y] += 1
				if ventCount[vent.start.x][y] == 2 {
					dangerousVentCount += 1
				}
			}
		} else if vent.start.y == vent.end.y { // Horizontal Line
			for x := min(vent.start.x, vent.end.x); x <= max(vent.start.x, vent.end.x); x++ {
				ventCount[x][vent.start.y] += 1
				if ventCount[x][vent.start.y] == 2 {
					dangerousVentCount += 1
				}
			}
		} else { // Diagonal Lines
			slope := float64(vent.start.y - vent.end.y) / float64(vent.start.x - vent.end.x)
			if slope != 1 && slope != -1 { continue }

			yIntercept := vent.start.y - int(slope) * vent.start.x

			for x := min(vent.start.x, vent.end.x); x <= max(vent.start.x, vent.end.x); x++ {
				y := int(slope) * x + yIntercept

				ventCount[x][y] += 1
				if ventCount[x][y] == 2 {
					dangerousVentCount += 1
				}
			}
		}
	}

	fmt.Printf("Part Two Answer: %d\n", dangerousVentCount)
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("[Error] Unable to open file")
		os.Exit(1)
	}
	defer file.Close()

	vents := []Vent{}
	size := Coordinate{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coordinatesStrings := strings.Split(scanner.Text(), " -> ")
		if len(coordinatesStrings) != 2 {
			fmt.Printf("[Error] Invalid vent line parsed\n")
			continue
		}

		start, err := toCoordinate(coordinatesStrings[0])
		if err != nil {
			fmt.Printf("[Error] Invalid start coordinate parsed\n")
			continue
		}
		if start.x > size.x { size.x = start.x }
		if start.y > size.y { size.y = start.y }

		end, err := toCoordinate(coordinatesStrings[1])
		if err != nil {
			fmt.Printf("[Error] Invalid end coordinate parsed\n")
			continue
		}
		if end.x > size.x { size.x = end.x }
		if end.y > size.y { size.y = end.y }

		vents = append(vents, Vent{start,end})
	}

	partOne(vents, size)
	partTwo(vents, size)
}
