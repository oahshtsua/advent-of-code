package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type passwordEntry struct {
	policyMinCount int
	policyMaxCount int
	policyChar     string
	password       string
}

func partOne(passwords []passwordEntry) int {
	validCount := 0
	for _, entry := range passwords {
		charCount := strings.Count(entry.password, entry.policyChar)
		if charCount >= entry.policyMinCount && charCount <= entry.policyMaxCount {
			validCount++
		}
	}
	return validCount
}

func partTwo(passwords []passwordEntry) int {
	validCount := 0
	for _, entry := range passwords {
		minIdxChar := entry.password[entry.policyMinCount-1]
		maxIdxChar := entry.password[entry.policyMaxCount-1]
		if (string(minIdxChar) == entry.policyChar) != (string(maxIdxChar) == entry.policyChar) {
			validCount++
		}
	}
	return validCount
}

func Solution() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var passwords []passwordEntry

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var minCount, maxCount int
		var char, password string

		_, err := fmt.Sscanf(scanner.Text(), "%d-%d %1s: %s", &minCount, &maxCount, &char, &password)
		if err != nil {
			log.Fatal(err)
		}

		pwEntry := passwordEntry{
			policyMinCount: minCount,
			policyMaxCount: maxCount,
			policyChar:     char,
			password:       password,
		}
		passwords = append(passwords, pwEntry)
	}

	fmt.Println("Part one: ", partOne(passwords))
	fmt.Println("Part two: ", partTwo(passwords))
}
