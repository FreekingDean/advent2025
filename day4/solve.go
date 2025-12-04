package day4

import (
	"fmt"
	"strings"
)

func Solve(input string) (string, string, error) {
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
	count, _ := remove(input)
	return fmt.Sprintf("%d", count), nil
}

func remove(input string) (int, string) {
	counts, parsedMap := parseMap(input)
	count := 0
	newMap := []string{}
	for i, row := range parsedMap {
		for j := range row {
			if counts[i][j] < 4 && parsedMap[i][j] == '@' {
				count += 1
				parsedMap[i][j] = '.'
			}
		}
		newMap = append(newMap, string(parsedMap[i]))
	}
	return count, strings.Join(newMap, "\n")
}

func Part2(input string) (string, error) {
	total := 0
	for {
		removed, newMap := remove(input)
		if removed == 0 {
			break
		}
		total += removed
		input = newMap
	}
	return fmt.Sprintf("%d", total), nil
}

func parseMap(input string) ([][]int, [][]rune) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make([][]int, len(lines))
	result[0] = make([]int, len(lines[0]))
	parsedMap := make([][]rune, len(lines))
	for row, line := range lines {
		parsedMap[row] = []rune(line)
		if line == "" {
			continue
		}
		if row != len(lines)-1 {
			result[row+1] = make([]int, len(line))
		}
		for col, char := range line {
			if char != '@' && char != '.' {
				continue
			}
			if char == '@' {
				incr(result, row-1, col-1)
				incr(result, row-1, col)
				incr(result, row-1, col+1)
				incr(result, row, col-1)
				incr(result, row, col+1)
				incr(result, row+1, col-1)
				incr(result, row+1, col)
				incr(result, row+1, col+1)
			}
		}
	}
	return result, parsedMap
}

func incr(arr [][]int, row, col int) {
	if row < 0 || row >= len(arr) {
		return
	}
	if col < 0 || col >= len(arr[row]) {
		return
	}
	arr[row][col] += 1
}
