package main

import (
	"aoc-2021/day01"
	"aoc-2021/day02"
	"flag"
	"fmt"
)

func main() {
	var day int

	flag.IntVar(&day, "day", 0, "AoC day number")
	flag.Parse()

	switch day {
	case 1:
		day01.Solution()
	case 2:
		day02.Solution()
	default:
		fmt.Println("Usage: go run main.go -day=n")
	}
}
