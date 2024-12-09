package day6

import (
	"AOC2024/utils"
	"log"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

func Part2() int {
	// Read input file
	lines, err := utils.ReadLines("input/day6.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// Convert lines to a 2D rune slice
	originalGrid := make([][]rune, len(lines))
	for i, line := range lines {
		originalGrid[i] = []rune(line)
	}

	// Find guard's starting position and direction
	xStart, yStart, dirStart := findGuardPosition(originalGrid)

	count := 0
	for y := 0; y < len(originalGrid); y++ {
		for x := 0; x < len(originalGrid[y]); x++ {
			if x == xStart && y == yStart {
				// Can't place an obstacle at guard's starting position
				continue
			}
			if originalGrid[y][x] == '.' {
				if causesLoop(originalGrid, x, y, xStart, yStart, dirStart) {
					count++
				}
			}
		}
	}

	return count
}

func causesLoop(originalGrid [][]rune, blockX, blockY, xStart, yStart, dirStart int) bool {
	grid := copyGrid(originalGrid)

	// Place the new obstruction
	grid[blockY][blockX] = '#'

	return simulateWithLoopDetection(grid, xStart, yStart, dirStart)
}

func simulateWithLoopDetection(grid [][]rune, x, y, direction int) bool {
	visitedStates := make(map[[3]int]bool)

	for {
		// Check for loop
		state := [3]int{x, y, direction}
		if visitedStates[state] {
			// We've been here before with the same direction
			return true
		}
		visitedStates[state] = true

		nx, ny, ndir, ok := get_next_move(x, y, direction, grid)
		if !ok {
			// Guard leaves the grid or no next move possible
			return false
		}

		// Move guard
		x, y, direction = nx, ny, ndir
	}
}

func get_next_move(x, y, dir int, grid [][]rune) (int, int, int, bool) {
	// Determine forward cell based on current direction
	dx, dy := 0, 0
	switch dir {
	case UP:
		dy = -1
	case DOWN:
		dy = 1
	case LEFT:
		dx = -1
	case RIGHT:
		dx = 1
	}

	nx, ny := x+dx, y+dy

	// Check if forward cell is within bounds
	if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[ny]) {
		// Forward out of bounds means guard leaves map
		return 0, 0, 0, false
	}

	// If there's an obstacle ahead, turn right
	if grid[ny][nx] == '#' {
		ndir := turnRight(dir)
		// After turning right, we attempt to move forward in the new direction:
		dx2, dy2 := 0, 0
		switch ndir {
		case UP:
			dy2 = -1
		case DOWN:
			dy2 = 1
		case LEFT:
			dx2 = -1
		case RIGHT:
			dx2 = 1
		}
		nx2, ny2 := x+dx2, y+dy2
		// Check bounds for the new forward cell
		if ny2 < 0 || ny2 >= len(grid) || nx2 < 0 || nx2 >= len(grid[ny2]) {
			// Can't move out of bounds, guard leaves map
			return 0, 0, 0, false
		}
		if grid[ny2][nx2] == '#' {
			// If still blocked after turning right, guard can't move anywhere
			// According to the problem, if no moves are possible, guard leaves the area
			return 0, 0, 0, false
		}
		// Move forward in the new direction
		return nx2, ny2, ndir, true
	}

	// Otherwise move forward in the current direction
	return nx, ny, dir, true
}

func turnRight(dir int) int {
	// directions: UP=0, DOWN=1, LEFT=2, RIGHT=3
	// but we must keep consistent. Let's define a cycle:
	// turnRight: UP -> RIGHT, RIGHT -> DOWN, DOWN -> LEFT, LEFT -> UP
	switch dir {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}
	return UP
}

func findGuardPosition(grid [][]rune) (int, int, int) {
	for y, row := range grid {
		for x, cell := range row {
			switch cell {
			case '^':
				return x, y, UP
			case 'v':
				return x, y, DOWN
			case '<':
				return x, y, LEFT
			case '>':
				return x, y, RIGHT
			}
		}
	}
	log.Fatalf("Guard starting position not found.")
	return 0, 0, UP
}

func copyGrid(original [][]rune) [][]rune {
	newGrid := make([][]rune, len(original))
	for i, row := range original {
		nrow := make([]rune, len(row))
		copy(nrow, row)
		newGrid[i] = nrow
	}
	return newGrid
}
