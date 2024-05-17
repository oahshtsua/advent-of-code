package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func partOne(expenses []int, tgt int) (int, int) {
	lookup := make(map[int]bool)

	var x, y int
	for _, entry := range expenses {
		if exists := lookup[tgt-entry]; exists {
			x, y = entry, tgt-entry
			break
		}
		lookup[entry] = true
	}
	return x, y
}

func partTwo(expenses []int, tgt int) (int, int, int) {
	var x, y, z int
	for _, entry := range expenses {
		x, y = partOne(expenses, tgt-entry)
		if x+y == (tgt - entry) {
			z = entry
			break
		}
	}
	return x, y, z
}

func Solution() {
	file, err := os.Open("day01/inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var expenses []int
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		expenses = append(expenses, num)
	}

	x, y := partOne(expenses, 2020)
	fmt.Println("Part one: ", x*y)

	x, y, z := partTwo(expenses, 2020)
	fmt.Println("Part one: ", x*y*z)
}
