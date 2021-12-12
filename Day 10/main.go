package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getOppositeRune(character rune) rune {
	if character == '(' { return ')' }
	if character == '[' { return ']' }
	if character == '{' { return '}' }

	return '>'
}

func partOne(lines []string) {
	score := 0
	for _, line := range lines {
		openedChunks := []rune{}
		for _, character := range line {
			if character == '(' || character == '[' || character == '{' || character == '<' {
				openedChunks = append(openedChunks, character)
			} else if character == getOppositeRune(openedChunks[len(openedChunks) - 1]) {
				openedChunks = openedChunks[:len(openedChunks) - 1]
			} else {
				if character == ')' {score += 3 }
				if character == ']' { score += 57 }
				if character == '}' { score += 1197 }
				if character == '>' { score += 25137 }
				break
			}
		}
	}

	fmt.Printf("Part One Answer: %d\n", score)
}

func partTwo(lines []string) {
	scores := []int{}
	for _, line := range lines {
		openedChunks := []rune{}
		corrupted := false
		for _, character := range line {
			if character == '(' || character == '[' || character == '{' || character == '<' {
				openedChunks = append(openedChunks, character)
			} else if character == getOppositeRune(openedChunks[len(openedChunks) - 1]) {
				openedChunks = openedChunks[:len(openedChunks) - 1]
			} else {
				corrupted = true
				break
			}
		}

		if corrupted {
			continue
		}

		score := 0
		for i := len(openedChunks) - 1; i >= 0; i-- {
			character := openedChunks[i]
			if character == '(' { score = score * 5 + 1 }
			if character == '[' { score = score * 5 + 2 }
			if character == '{' { score = score * 5 + 3 }
			if character == '<' { score = score * 5 + 4 }
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)

	fmt.Printf("Part Two Answer: %d\n", scores[(len(scores) - 1) / 2])
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("[Error] Unable to open file")
		os.Exit(1)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		lines = append(lines, line)
	}

	partOne(lines)
	partTwo(lines)
}
