// Day 1: Calorie Counting
package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func partOne(calories []int) int {
	var maxCalorie int
	for _, calorie := range calories {
		if calorie > maxCalorie {
			maxCalorie = calorie
		}
	}
	return maxCalorie
}

func partTwo(calories []int) int {
	var maxThreeCalories [3]int
	for _, calorie := range calories {
		if calorie > maxThreeCalories[0] {
			maxThreeCalories[2] = maxThreeCalories[1]
			maxThreeCalories[1] = maxThreeCalories[0]
			maxThreeCalories[0] = calorie
		} else if calorie > maxThreeCalories[1] {
			maxThreeCalories[2] = maxThreeCalories[1]
			maxThreeCalories[1] = calorie
		} else if calorie > maxThreeCalories[2] {
			maxThreeCalories[2] = calorie
		}
	}
	return maxThreeCalories[0] + maxThreeCalories[1] + maxThreeCalories[2]
}

func Solution() {
	file, err := os.Open("day01/inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var calories []int
	var count int

	for scanner.Scan() {
		if scanner.Text() == "" {
			calories = append(calories, count)
			count = 0
		} else {
			calorie, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			count += calorie
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part one: ", partOne(calories))
	fmt.Println("Part two: ", partTwo(calories))
}
