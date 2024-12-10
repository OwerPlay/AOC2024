package day10

import (
	"AOC2024/utils"
)

func Part1() int {
	lines, _ := utils.ReadLines("input/day10.txt")
	if len(lines) == 0 {
		return 0
	}
	R, C := len(lines), len(lines[0])
	grid := make([][]int, R)
	for i := range grid {
		grid[i] = make([]int, C)
		for j := range lines[i] {
			grid[i][j] = int(lines[i][j] - '0')
		}
	}
	res := 0
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == 0 {
				seen := make([][]bool, R)
				for i := range seen {
					seen[i] = make([]bool, C)
				}
				type cell struct{ r, c int }
				q := []cell{{r, c}}
				seen[r][c] = true
				reached := map[[2]int]bool{}
				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					if grid[cur.r][cur.c] == 9 {
						reached[[2]int{cur.r, cur.c}] = true
					}
					for _, d := range dirs {
						nr, nc := cur.r+d[0], cur.c+d[1]
						if nr >= 0 && nr < R && nc >= 0 && nc < C && !seen[nr][nc] {
							if grid[nr][nc] == grid[cur.r][cur.c]+1 {
								seen[nr][nc] = true
								q = append(q, cell{nr, nc})
							}
						}
					}
				}
				res += len(reached)
			}
		}
	}
	return res
}

func Part2() int {
	lines, _ := utils.ReadLines("input/day10.txt")
	if len(lines) == 0 {
		return 0
	}
	R, C := len(lines), len(lines[0])
	grid := make([][]int, R)
	for i := range grid {
		grid[i] = make([]int, C)
		for j := range lines[i] {
			grid[i][j] = int(lines[i][j] - '0')
		}
	}
	dp := make([][]int, R)
	for i := range dp {
		dp[i] = make([]int, C)
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == 9 {
				dp[r][c] = 1
			}
		}
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for h := 8; h >= 0; h-- {
		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {
				if grid[r][c] == h {
					sum := 0
					for _, d := range dirs {
						nr, nc := r+d[0], c+d[1]
						if nr >= 0 && nr < R && nc >= 0 && nc < C && grid[nr][nc] == h+1 {
							sum += dp[nr][nc]
						}
					}
					dp[r][c] = sum
				}
			}
		}
	}
	res := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == 0 {
				res += dp[r][c]
			}
		}
	}
	return res
}
