package day3

import (
	"AOC2024/utils"
	"regexp"
	"strconv"
)

func Part1() int {
	lines, _ := utils.ReadLines("input/day3.txt")
	pattern := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(pattern)

	sum := 0
	numberPattern := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			numbers := numberPattern.FindAllString(match, -1)
			number1, _ := strconv.Atoi(numbers[0])
			number2, _ := strconv.Atoi(numbers[1])
			sum += number1 * number2
		}
	}
	return sum
}

func Part2() int {
	lines, _ := utils.ReadLines("input/day3.txt")
	pattern := `mul\(\d+,\d+\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	sum := 0
	numberPattern := regexp.MustCompile(`\d+`)
	enabled := true
	for _, line := range lines {
		matches := re.FindAllString(line, -1)

		for _, match := range matches {
			if match == "do()" {
				enabled = true
				continue
			} else if match == "don't()" {
				enabled = false
				continue
			}
			if enabled {
				numbers := numberPattern.FindAllString(match, -1)
				number1, _ := strconv.Atoi(numbers[0])
				number2, _ := strconv.Atoi(numbers[1])
				sum += number1 * number2
			}
		}
	}
	return sum
}
