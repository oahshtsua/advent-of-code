package day01

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne(expenses []int, target int) (int, int) {
	lookup := make(map[int]int)

	var x, y int
	for _, entry := range expenses {
		if val, ok := lookup[entry]; ok {
			x, y = val, entry
			break
		} else {
			lookup[target-entry] = entry
		}
	}
	return x, y
}

func partTwo(expenses []int, target int) (int, int, int) {
	var x, y, z int

	for _, entry := range expenses {
		tempTarget := target - entry
		x, y = partOne(expenses, tempTarget)
		if x+y == tempTarget {
			z = entry
			break
		}
	}
	return x, y, z
}

func Solution() {
	file, err := os.ReadFile("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var expenses []int
	for _, line := range strings.Split(string(file), "\n") {
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			expenses = append(expenses, num)
		}
	}

	target := 2020

	x, y := partOne(expenses, target)
	fmt.Printf("Part one: %d\n", x*y)

	x, y, z := partTwo(expenses, target)
	fmt.Printf("Part two: %d\n", x*y*z)
}
