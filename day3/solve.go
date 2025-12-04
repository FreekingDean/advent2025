package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(input string) (string, string, error) {
	// Placeholder implementation
	part1, err := Part1(input)
	if err != nil {
		return "", "", err
	}
	part2, err := Part2(input)
	if err != nil {
		return "", "", err
	}
	return part1, part2, nil
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		num, err := findLargest(line, 2)
		if err != nil {
			return "", fmt.Errorf("error finding largest in line '%s': %w", line, err)
		}
		sum += num
	}
	return fmt.Sprintf("%d", sum), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		num, err := findLargest(line, 12)
		if err != nil {
			return "", fmt.Errorf("error finding largest in line '%s': %w", line, err)
		}
		sum += num
	}
	return fmt.Sprintf("%d", sum), nil
}

func findLargest(input string, count int) (int, error) {
	if input == "" {
		return 0, nil
	}
	numStr := make([]rune, count)
	for i, char := range input {
		for j := 0; j < count; j++ {
			if char > numStr[j] && i < len(input)-(count-j-1) {
				numStr[j] = char
				for k := j + 1; k < count; k++ {
					numStr[k] = 0
				}
				break
			}
		}
	}
	num, err := strconv.Atoi(string(numStr))
	if err != nil {
		return 0, fmt.Errorf("invalid number '%s': %w", string(numStr), err)
	}
	return num, nil
}
