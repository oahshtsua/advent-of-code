// Day 1: Sonar Sweep
package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateIncrements(depths []int, step int) int {
	var increments int
	for i := 0; i < len(depths)-step; i++ {
		if depths[i] < depths[i+step] {
			increments++
		}
	}
	return increments

}
func partOne(depths []int) int {
	return calculateIncrements(depths, 1)
}

func partTwo(depths []int) int {
	return calculateIncrements(depths, 3)
}

func Solution() {
	file, err := os.Open("day01/inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var depths []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part one: ", partOne(depths))
	fmt.Println("Part two: ", partTwo(depths))
}
