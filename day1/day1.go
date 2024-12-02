package day1

import (
	"AOC2024/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Part1() int {
	// Read input file
	lines, err := utils.ReadLines("input/day1.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// Parse input
	left, right, err := ParseInput(lines)
	if err != nil {
		log.Fatalf("Failed to parse input: %v", err)
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		totalDistance += utils.Abs(left[i] - right[i])
	}
	return totalDistance
}

func Part2() int {
	// Read input file
	lines, err := utils.ReadLines("input/day1.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// Parse input
	left, right, err := ParseInput(lines)
	if err != nil {
		log.Fatalf("Failed to parse input: %v", err)
	}

	rightCounts := make(map[int]int)
	for _, num := range right {
		rightCounts[num]++
	}

	similarityScore := 0
	for _, num := range left {
		similarityScore += num * rightCounts[num]
	}
	return similarityScore
}

func ParseInput(lines []string) ([]int, []int, error) {
	var left, right []int
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}
		left = append(left, leftNum)
		right = append(right, rightNum)
	}
	return left, right, nil
}
