package day1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve(input string) (string, error) {
	part1, err := part1(input)
	if err != nil {
		return "", fmt.Errorf("error solving part 1: %w", err)
	}

	part2, err := part2(input)
	if err != nil {
		return "", fmt.Errorf("error solving part 2: %w", err)
	}
	return fmt.Sprintf("Part 1:\n%s\n\nPart 2:\n%s", part1, part2), nil
}

func part2(input string) (string, error) {
	cur := 50
	lines := strings.Split(input, "\n")
	zeroes := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		mult := 1
		if line[0] == 'L' {
			mult = -1
		}

		rot, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", fmt.Errorf("invalid rotation in line '%s': %w", line, err)
		}
		if rot == 0 {
			continue
		}
		rot = (mult * rot) + cur
		if rot == 0 || (rot < 0 && cur != 0) {
			zeroes += 1
		}

		zeroes += int(math.Floor(math.Abs(float64(rot)) / 100))
		cur = ((rot % 100) + 100) % 100
		fmt.Println("rot=", rot, " zeroes=", zeroes, " line=", line, " cur=", cur)
	}

	// Placeholder implementation
	return strconv.Itoa(zeroes), nil
}

func part1(input string) (string, error) {
	cur := 50
	lines := strings.Split(input, "\n")
	zeroes := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		mult := 1
		if line[0] == 'L' {
			mult = -1
		}

		rot, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", fmt.Errorf("invalid rotation in line '%s': %w", line, err)
		}
		rot = (mult * rot) + cur
		cur = ((rot % 100) + 100) % 100
		if cur == 0 {
			zeroes++
		}
	}

	// Placeholder implementation
	return strconv.Itoa(zeroes), nil
}
