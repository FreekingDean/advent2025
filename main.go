package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/FreekingDean/advent2025/day1"
	"github.com/FreekingDean/advent2025/day2"
	"github.com/FreekingDean/advent2025/day3"

	"github.com/FreekingDean/advent2025/utils"
)

type SolveFunc func(input string) (string, string, error)

var days = map[string]SolveFunc{
	"day1": day1.Solve,
	"day2": day2.Solve,
	"day3": day3.Solve,
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 || args[0] == "" {
		log.Fatalf("Please provide one day to solve, got: '%+v'", args)
	}
	day := args[0]

	inputFile := path.Join(".", day, "input.txt")
	input, err := utils.LoadInput(inputFile)
	if err != nil {
		log.Fatalf("error opening '%s', got: %s", inputFile, err)
	}

	solveFunc, ok := days[day]
	if !ok {
		log.Fatalf("no solve function for '%s'", day)
	}

	part1, part2, err := solveFunc(input)
	if err != nil {
		log.Fatalf("error solving '%s': %s", day, err)
	}
	fmt.Printf("Solution for %s part1:\n====================\n%s\n\n", day, part1)
	fmt.Printf("Solution for %s part2:\n====================\n%s\n\n", day, part2)
}
