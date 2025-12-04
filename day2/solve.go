package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(input string) (string, string, error) {
	part1, err := Part1(input)
	if err != nil {
		return "", "", fmt.Errorf("error solving part 1: %w", err)
	}

	part2, err := Part2(input)
	if err != nil {
		return "", "", fmt.Errorf("error solving part 2: %w", err)
	}

	// Placeholder implementation
	return part1, part2, nil
}

func Part2(input string) (string, error) {
	ranges := strings.Split(input, ",")
	sum := 0
	for _, r := range ranges {
		numbers := strings.Split(r, "-")
		if len(numbers) != 2 {
			return "", fmt.Errorf("invalid range '%s'", r)
		}

		num1, err := strconv.Atoi(strings.TrimSpace(numbers[0]))
		if err != nil {
			return "", fmt.Errorf("invalid number in range '%s': %w", r, err)
		}
		num2, err := strconv.Atoi(strings.TrimSpace(numbers[1]))
		if err != nil {
			return "", fmt.Errorf("invalid number in range '%s': %w", r, err)
		}

		seen := map[int]struct{}{}
		for x := num1; x <= num2; x++ {
			xstr := fmt.Sprintf("%d", x)
			for i := 1; i <= len(xstr)/2; i++ {
				prefix := xstr[0:i]
				newStr := ""
				for j := 1; j <= len(xstr)/i; j++ {
					newStr += prefix
				}
				if _, ok := seen[x]; !ok && xstr == newStr {
					sum += x
					seen[x] = struct{}{}
				}
			}
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func Part1(input string) (string, error) {
	ranges := strings.Split(input, ",")
	sum := 0
	for _, r := range ranges {
		numbers := strings.Split(r, "-")
		if len(numbers) != 2 {
			return "", fmt.Errorf("invalid range '%s'", r)
		}

		num1, err := strconv.Atoi(strings.TrimSpace(numbers[0]))
		if err != nil {
			return "", fmt.Errorf("invalid number in range '%s': %w", r, err)
		}
		num2, err := strconv.Atoi(strings.TrimSpace(numbers[1]))
		if err != nil {
			return "", fmt.Errorf("invalid number in range '%s': %w", r, err)
		}

		toDouble, _ := strconv.Atoi(numbers[0][0 : len(numbers[0])/2])
		n := 0
		for n < num1 {
			if num1 == 3 {
			}
			n, err = strconv.Atoi(fmt.Sprintf("%d%d", toDouble, toDouble))
			if num1 == 3 {
				fmt.Printf("n=%d num1=%d num2=%d toDouble=%d\n", n, num1, num2, toDouble)
			}
			if err != nil {
				return "", fmt.Errorf("invalid number during doubling '%d': %w", toDouble, err)
			}
			toDouble++
		}
		for n <= num2 {
			sum += n
			//fmt.Printf("n=%d num1=%d num2=%d toDouble=%d\n", n, num1, num2, toDouble)
			n, err = strconv.Atoi(fmt.Sprintf("%d%d", toDouble, toDouble))
			if err != nil {
				return "", fmt.Errorf("invalid number during doubling '%d': %w", toDouble, err)
			}
			toDouble++
		}
	}
	// Placeholder implementation
	return fmt.Sprintf("%d", sum), nil
}
