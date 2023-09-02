package main

import (
	"flag"
	"fmt"
)

func main() {
	var day int

	flag.IntVar(&day, "day", 0, "AoC day number")
	flag.Parse()

	switch day {
	default:
		fmt.Println("Usage: go run main.go -day=n")
	}
}
