package day11

import (
	"AOC2024/utils"
	"log"
	"strconv"
	"strings"
)

func Part1() int {
	lines, err := utils.ReadLines("input/day11.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	stones := parseInput(lines[0])

	for i := 0; i < 25; i++ {
		stones = transformStones(stones)
	}

	return len(stones)
}

func Part2() int {
	lines, err := utils.ReadLines("input/day11.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	stones := parseInput(lines[0])

	steps := 75
	memo := make(map[[2]int]int)

	totalCount := 0
	for _, s := range stones {
		totalCount += countStonesAfterN(s, steps, memo)
	}

	return totalCount
}

func parseInput(line string) []int {
	splitLines := strings.Split(line, " ")
	intLine := make([]int, len(splitLines))
	for i, s := range splitLines {
		intLine[i], _ = strconv.Atoi(s)
	}
	return intLine
}

func transformStones(stones []int) []int {
	newStones := make([]int, 0)
	for _, s := range stones {
		switch {
		case s == 0:
			newStones = append(newStones, 1)

		case len(strconv.Itoa(s))%2 == 0:
			sStr := strconv.Itoa(s)
			half := len(sStr) / 2
			firstHalf := sStr[:half]
			secondHalf := sStr[half:]

			f1, _ := strconv.Atoi(firstHalf)
			f2, _ := strconv.Atoi(secondHalf)
			newStones = append(newStones, f1)
			newStones = append(newStones, f2)
		default:
			newStones = append(newStones, s*2024)
		}
	}
	return newStones
}

func countStonesAfterN(stone int, steps int, memo map[[2]int]int) int {
	if steps == 0 {
		return 1
	}

	key := [2]int{stone, steps}
	if val, found := memo[key]; found {
		return val
	}

	count := 0

	switch {
	case stone == 0:
		count += countStonesAfterN(1, steps-1, memo)

	case len(strconv.Itoa(stone))%2 == 0:
		sStr := strconv.Itoa(stone)
		half := len(sStr) / 2
		firstHalf := sStr[:half]
		secondHalf := sStr[half:]

		f1, _ := strconv.Atoi(firstHalf)
		f2, _ := strconv.Atoi(secondHalf)

		count += countStonesAfterN(f1, steps-1, memo)
		count += countStonesAfterN(f2, steps-1, memo)

	default:
		newStone := stone * 2024
		count += countStonesAfterN(newStone, steps-1, memo)
	}

	memo[key] = count
	return count
}
