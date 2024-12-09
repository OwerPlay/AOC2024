package day7

import (
	"AOC2024/utils"
	"log"
	"math"
	"strconv"
	"strings"
)

func Part1() int {
	lines, err := utils.ReadLines("input/day7.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	left, right := ParseInput(lines)

	sum := 0
	for i := range left {
		totalCombinations := 1 << (len(right[i]) - 1)

		for combo := 0; combo < totalCombinations; combo++ {
			binary := make([]int, len(right[i])-1)
			for bit := 0; bit < len(binary); bit++ {
				if (combo & (1 << bit)) != 0 {
					binary[bit] = 1
				}
			}

			calc := right[i][0]
			for k := 0; k < len(binary); k++ {
				if binary[k] == 1 {
					calc *= right[i][k+1]
				} else {
					calc += right[i][k+1]
				}
			}

			// fmt.Println("Binary:", binary)
			// fmt.Println("Calc:", calc)
			// fmt.Println("Left:", left[i])
			if calc == left[i] {
				sum += left[i]
				break
			}
		}
	}
	return sum
}

func Part2() int {
	lines, err := utils.ReadLines("input/day7.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	left, right := ParseInput(lines)

	sum := 0
	for i := range left {
		if len(right[i]) == 1 {
			if right[i][0] == left[i] {
				sum += left[i]
			}
			continue
		}

		length := len(right[i]) - 1
		totalCombinations := 1
		for j := 0; j < length; j++ {
			totalCombinations *= 3
		}

	combinationLoop:
		for combo := 0; combo < totalCombinations; combo++ {
			ops := make([]int, length)
			tmp := combo
			for k := 0; k < length; k++ {
				ops[k] = tmp % 3
				tmp /= 3
			}

			calc := right[i][0]
			for k := 0; k < len(ops); k++ {
				nextVal := right[i][k+1]
				switch ops[k] {
				case 0:
					calc += nextVal
				case 1:
					calc *= nextVal
				case 2:
					digits := numDigits(nextVal)
					if nextVal >= 0 {
						calc = calc*int(math.Pow10(digits)) + nextVal
					} else {
						posVal := -nextVal
						dpos := numDigits(posVal)
						calc = calc*int(math.Pow10(dpos+1)) - posVal
					}
				}
			}

			if calc == left[i] {
				sum += left[i]
				break combinationLoop
			}
		}
	}
	return sum
}

func ParseInput(lines []string) ([]int, [][]int) {

	left := []int{}
	right := [][]int{}
	for _, line := range lines {
		parts := strings.Split(line, ":")
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}
		left = append(left, num)

		parts = strings.Split(strings.TrimLeft(parts[1], " "), " ")
		right_nums := []int{}
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Fatalf("Failed to convert string to int: %v", err)
			}
			right_nums = append(right_nums, num)
		}
		right = append(right, right_nums)
	}

	return left, right
}

func numDigits(x int) int {
	if x == 0 {
		return 1
	}
	count := 0
	n := x
	if n < 0 {
		n = -n
	}
	for n > 0 {
		n /= 10
		count++
	}
	return count
}
