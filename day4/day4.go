package day4

import (
	"AOC2024/utils"
	"log"
	"strings"
)

func Part1() int {
	lines, err := utils.ReadLines("input/day4.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	total := 0
	for _, line := range lines {
		matches := strings.Count(line, "XMAS")
		total += matches
		matches = strings.Count(line, "SAMX")
		total += matches
	}

	for col := 0; col < len(lines[0]); col++ {
		column := ""
		for row := 0; row < len(lines); row++ {
			column += string(lines[row][col])
		}
		matches := strings.Count(column, "XMAS")
		total += matches
		matches = strings.Count(column, "SAMX")
		total += matches
	}

	// Check diagonals from top-left to bottom-right
	for startCol := 0; startCol < len(lines[0])-3; startCol++ {
		for startRow := 0; startRow < len(lines)-3; startRow++ {
			diagonal := ""
			for i := 0; i < 4; i++ {
				diagonal += string(lines[startRow+i][startCol+i])
			}
			if diagonal == "XMAS" || diagonal == "SAMX" {
				total++
			}
		}
	}

	// Check diagonals from top-right to bottom-left
	for startCol := len(lines[0]) - 1; startCol >= 3; startCol-- {
		for startRow := 0; startRow < len(lines)-3; startRow++ {
			diagonal := ""
			for i := 0; i < 4; i++ {
				diagonal += string(lines[startRow+i][startCol-i])
			}
			if diagonal == "XMAS" || diagonal == "SAMX" {
				total++
			}
		}
	}

	return total
}

func Part2() int {
	lines, err := utils.ReadLines("input/day4.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	total := 0

	for startCol := 0; startCol < len(lines[0])-2; startCol++ {
		for startRow := 0; startRow < len(lines)-2; startRow++ {
			diagonal := ""
			for i := 0; i < 3; i++ {
				diagonal += string(lines[startRow+i][startCol+i])
			}
			if diagonal == "MAS" || diagonal == "SAM" {
				diagonal = ""
				for i := 0; i < 3; i++ {
					diagonal += string(lines[startRow+i][startCol+2-i])
					if diagonal == "MAS" || diagonal == "SAM" {
						total++
					}
				}
			}
		}
	}

	return total
}
