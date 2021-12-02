package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction string

const(
	Forward Direction = "forward"
	Up = "up"
	Down = "down"
)

type Command struct {
	direction Direction
	amount int
}

func partOne(commands []Command) {
	horiPos := 0
	depPos := 0
	for _, command := range commands { 
		switch command.direction {
			case Forward:
				horiPos += command.amount

			case Up:
				depPos -= command.amount

			case Down:
				depPos += command.amount
		}
	}

	fmt.Printf("Part One Answer: %d\n", horiPos * depPos)
}

func partTwo(commands []Command) {
	horiPos := 0
	depPos := 0
	aimPos := 0
	for _, command := range commands { 
		switch command.direction {
			case Forward:
				horiPos += command.amount
				depPos += aimPos * command.amount

			case Up:
				aimPos -= command.amount

			case Down:
				aimPos += command.amount
		}
	}

	fmt.Printf("Part Two Answer: %d\n", horiPos * depPos)
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("[Error] Unable to open file")
		os.Exit(1)
	}
	defer file.Close()

	commands := []Command {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		if len(command) != 2 {
			fmt.Printf("[Error] Invalid command parsed")
			continue
		}

		amount, err := strconv.Atoi(command[1])
		if err != nil {
			fmt.Printf("[Error] Invalid amount parsed")
			continue
		}

		switch command[0] {
			case "forward":
				commands = append(commands, Command{Forward,amount})

			case "up":
				commands = append(commands, Command{Up,amount})

			case "down":
				commands = append(commands, Command{Down,amount})

			default:
				fmt.Printf("[Error] Invalid command parsed")
				continue
		}
	}

	partOne(commands)
	partTwo(commands)
}
