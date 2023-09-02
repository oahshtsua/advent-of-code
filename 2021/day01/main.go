package day01

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne(depths []int) int {
	var count int
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}
	return count
}

func partTwo(depths []int) int {
	var count int
	for i := 3; i < len(depths); i++ {
		if depths[i] > depths[i-3] {
			count++
		}
	}
	return count
}

func Solution() {
	file, err := os.ReadFile("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var depths []int
	for _, line := range strings.Split(string(file), "\n") {
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			depths = append(depths, num)
		}
	}

	fmt.Printf("Part one: %d\n", partOne(depths))
	fmt.Printf("Part two: %d\n", partTwo(depths))
}
