package day5

import (
	"AOC2024/utils"
	"log"
	"slices"
	"strconv"
	"strings"
)

func Part1() int {
	// Read input file
	lines, err := utils.ReadLines("input/day5.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// Parse input
	rules_map := ParseRulesInput(lines)

	// Read input file
	lines, err = utils.ReadLines("input/day5_2.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	input := ParseTextInput(lines)

	counter := 0
	for _, line := range input {
		line_ok, _, _ := is_ok(line, rules_map)
		if line_ok {
			counter += line[len(line)/2]
		}

	}
	return counter
}

func Part2() int {
	// Read input file
	lines, err := utils.ReadLines("input/day5.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// Parse input
	rules_map := ParseRulesInput(lines)

	// Read input file
	lines, err = utils.ReadLines("input/day5_2.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	input := ParseTextInput(lines)

	counter := 0
	for _, line := range input {
		line_ok, index, index_2 := is_ok(line, rules_map)
		if line_ok {
			continue
		}

		for !line_ok {
			// Debugging output
			// fmt.Printf("line_ok: %v, index: %d, index_2: %d, line: %v\n", line_ok, index, index_2, line)

			// Ensure indices are valid
			if index < 0 || index >= len(line) || index_2 < 0 || index_2 >= len(line) {
				log.Fatalf("Invalid indices: index=%d, index_2=%d, len(line)=%d", index, index_2, len(line))
			}

			// Remove the value at index_2
			value := line[index_2]
			line = append(line[:index_2], line[index_2+1:]...)

			// Adjust `index` if it becomes out of bounds
			if index >= len(line) {
				index = len(line) - 1
			}

			// Insert the value directly after `index`
			line = append(line[:index+1], append([]int{value}, line[index+1:]...)...)

			// Recheck the line
			line_ok, index, index_2 = is_ok(line, rules_map)
		}

		counter += line[len(line)/2]
	}
	return counter
}

func ParseRulesInput(lines []string) map[int][]int {
	rules_map := make(map[int][]int)
	for _, line := range lines {
		parts := strings.Split(line, "|")
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Failed to parse input: %v", err)
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Failed to parse input: %v", err)
		}
		rules_map[num] = append(rules_map[num], num2)
	}
	return rules_map
}

func ParseTextInput(lines []string) [][]int {
	text_input := make([][]int, len(lines))
	for index, line := range lines {
		split_line := strings.Split(line, ",")
		numbers := make([]int, 0, len(split_line))
		for _, s := range split_line {
			s_number, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("Failed to parse input: %v", err)
			}
			numbers = append(numbers, s_number)
		}
		text_input[index] = numbers
	}
	return text_input
}

func is_ok(line []int, rules_map map[int][]int) (bool, int, int) {
	for i := len(line) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if slices.Contains(rules_map[line[i]], line[j]) {
				return false, i, j
			}
		}
	}
	return true, -1, -1
}
