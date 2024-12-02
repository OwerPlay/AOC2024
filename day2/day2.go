package day2

import (
	"AOC2024/utils"
	"strconv"
	"strings"
)

func Part1() int {
	lines, err := utils.ReadLines("input/day2.txt")
	if err != nil {
		panic(err)
	}

	divided := parseReports(lines)
	counter := 0
	for _, line := range divided {
		if isSafe(line) {
			counter++
		}
	}
	return counter
}

func Part2() int {
	lines, err := utils.ReadLines("input/day2.txt")
	if err != nil {
		panic(err)
	}

	divided := parseReports(lines)
	counter := 0

	for _, line := range divided {
		if isSafe(line) {
			counter++
		} else {
			isOK := false
			for i := range line {
				modified := append([]int{}, line[:i]...)
				modified = append(modified, line[i+1:]...)
				if isSafe(modified) {
					isOK = true
					break
				}
			}
			if isOK {
				counter++
			}
		}
	}
	return counter
}

func isSafe(line []int) bool {
	if isSortedAscending(line) || isSortedDescending(line) {
		for i := 1; i < len(line); i++ {
			diff := utils.Abs(line[i] - line[i-1])
			if diff == 0 || diff > 3 {
				return false
			}
		}
		return true
	}
	return false
}

func isSortedAscending(line []int) bool {
	for i := 1; i < len(line); i++ {
		if line[i] < line[i-1] {
			return false
		}
	}
	return true
}

func isSortedDescending(line []int) bool {
	for i := 1; i < len(line); i++ {
		if line[i] > line[i-1] {
			return false
		}
	}
	return true
}

func parseReports(lines []string) [][]int {
	var reports [][]int
	for _, line := range lines {
		fields := strings.Fields(line)
		report := make([]int, len(fields))
		for i, field := range fields {
			num, _ := strconv.Atoi(field)
			report[i] = num
		}
		reports = append(reports, report)
	}
	return reports
}
