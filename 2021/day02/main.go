package day02

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne(commands [][]string) int {
	var horizontalPos, depth int

	for _, command := range commands {
		direction := command[0]
		step, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			horizontalPos += step
		case "down":
			depth += step
		case "up":
			depth -= step
		}
	}
	return horizontalPos * depth
}

func partTwo(commands [][]string) int {
	var horizontalPos, depth, aim int

	for _, command := range commands {
		direction := command[0]
		step, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			horizontalPos += step
			depth += aim * step
		case "down":
			aim += step
		case "up":
			aim -= step
		}
	}
	return horizontalPos * depth
}

func Solution() {
	file, err := os.ReadFile("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var commands [][]string
	for _, line := range strings.Split(string(file), "\n") {
		if line != "" {
			command := strings.Split(line, " ")
			if err != nil {
				log.Fatal(err)
			}
			commands = append(commands, command)
		}
	}
	fmt.Printf("Part one: %d\n", partOne(commands))
	fmt.Printf("Part two: %d\n", partTwo(commands))
}
