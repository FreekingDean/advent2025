package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/FreekingDean/advent2025/day1"

	"github.com/FreekingDean/advent2025/utils"
)

type SolveFunc func(input string) string

var days = map[string]SolveFunc{
	"1": day1.Solve,
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 || args[0] == "" {
		log.Fatal(
			fmt.Sprintf("Please provide one day to solve, got: '%+v'", argsWithoutProg),
		)
	}
	day := args[0]

	inputFile := path.Join(".", day, "input.txt")
	input, err := utils.LoadInput(inputFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("error opening '%s', got: %w", err))
	}
}
