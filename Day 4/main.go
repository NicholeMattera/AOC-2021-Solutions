package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type BingoBoardSpace struct {
	marked bool
	value int
}

func calculateScore(board []BingoBoardSpace, callout int, invert bool) int {
	score := 0
	for _, space := range board {
		if !invert && !space.marked { score += space.value }
		if invert && space.marked { score += space.value }
	}
	
	return score * callout
}

func checkForWinner(board []BingoBoardSpace) bool {
	for y := 0; y < 5; y++ {
		// Horizontal
		if board[5 * y].marked && board[5 * y + 1].marked && board[5 * y + 2].marked && board[5 * y + 3].marked && board[5 * y + 4].marked {
			return true
		}

		// Vertical
		if board[y].marked && board[5 + y].marked && board[10 + y].marked && board[15 + y].marked && board[20 + y].marked {
			return true
		}
	}

	return false
}

func partOne(callouts []int, boards [][]BingoBoardSpace) {
	winningBoard := []BingoBoardSpace{}
	winningCallout := 0

	for _, callout := range callouts {
		winnerFound := false

		for boardIndex := range boards {
			for spaceIndex := range boards[boardIndex] {
				if boards[boardIndex][spaceIndex].value == callout {
					boards[boardIndex][spaceIndex].marked = true

					if checkForWinner(boards[boardIndex]) {
						winnerFound = true
						winningBoard = boards[boardIndex]
						winningCallout = callout
					}

					break
				}
			}

			if winnerFound {
				break
			}
		}

		if winnerFound {
			break
		}
	}

	score := calculateScore(winningBoard, winningCallout, false)

	fmt.Printf("Part One Answer: %d\n", score)
}

func checkForLoser(board []BingoBoardSpace) bool {
	incompleteLines := 0

	for y := 0; y < 5; y++ {
		// Horizontal
		if board[5 * y].marked || board[5 * y + 1].marked || board[5 * y + 2].marked || board[5 * y + 3].marked || board[5 * y + 4].marked {
			incompleteLines += 1 
		}

		// Vertical
		if board[y].marked || board[5 + y].marked || board[10 + y].marked || board[15 + y].marked || board[20 + y].marked {
			incompleteLines += 1 
		}
	}

	return incompleteLines == 10
}

func partTwo(callouts []int, boards [][]BingoBoardSpace) {
	losingBoard := []BingoBoardSpace{}
	winningCallout := 0

	for calloutIndex := len(callouts) - 1; calloutIndex >= 0; calloutIndex-- {
		callout := callouts[calloutIndex]
		loserFound := false

		for boardIndex := range boards {
			for spaceIndex := range boards[boardIndex] {
				if boards[boardIndex][spaceIndex].value == callout {
					boards[boardIndex][spaceIndex].marked = true


					if checkForLoser(boards[boardIndex]) {
						loserFound = true
						boards[boardIndex][spaceIndex].marked = false
						losingBoard = boards[boardIndex]
						winningCallout = callout
					}

					break
				}
			}

			if loserFound {
				break
			}
		}

		if loserFound {
			break
		}
	}

	score := calculateScore(losingBoard, winningCallout, true)

	fmt.Printf("Part Two Answer: %d\n", score)
}

func copyBoards(boards [][]BingoBoardSpace) [][]BingoBoardSpace {
	duplicate := make([][]BingoBoardSpace, len(boards))
	for i := range boards {
		duplicate[i] = make([]BingoBoardSpace, len(boards[i]))
		copy(duplicate[i], boards[i])
	}
	return duplicate
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("[Error] Unable to open file")
		os.Exit(1)
	}
	defer file.Close()

	callouts := []int{}
	boards := [][]BingoBoardSpace{}
	lineNum := 0
	scanner := bufio.NewScanner(file)
	cellIndex := 0
	for scanner.Scan() {
		if lineNum == 0 {
			calloutStrings := strings.Split(scanner.Text(), ",")

			for _, calloutString := range calloutStrings {
				callout, err := strconv.Atoi(calloutString)
				if err != nil {
					fmt.Printf("[Error] Invalid callout parsed\n")
					continue
				}

				callouts = append(callouts, callout)
			}
		} else if (lineNum - 1) % 6 == 0 {
			boards = append(boards, make([]BingoBoardSpace, 25))
			cellIndex = 0
		} else {
			cellStrings := strings.Split(scanner.Text(), " ")

			for _, cellString := range cellStrings {
				if len(cellString) == 0 {
					continue
				}

				cell, err := strconv.Atoi(cellString)
				if err != nil {
					fmt.Printf("[Error] Invalid cell parsed\n")
					continue
				}

				boardNumber := int(math.Floor(float64(lineNum - 1) / 6))
				boards[boardNumber][cellIndex] = BingoBoardSpace{false,cell}
				cellIndex += 1
			}
		}

		lineNum += 1
	}

	partOne(callouts, copyBoards(boards))
	partTwo(callouts, copyBoards(boards))
}
