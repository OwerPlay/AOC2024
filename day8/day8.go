package day8

import (
	"AOC2024/utils"
	"unicode"
)

type coord struct{ y, x int }

func gcd(a, b int) int {
	if b == 0 {
		if a < 0 {
			return -a
		}
		return a
	}
	return gcd(b, a%b)
}

func Part1() int {
	lines, err := utils.ReadLines("input/day8.txt")
	if err != nil {
		panic(err)
	}

	height := len(lines)
	if height == 0 {
		return 0
	}
	width := len(lines[0])

	antennasByFreq := make(map[rune][]coord)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ch := rune(lines[y][x])
			if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
				antennasByFreq[ch] = append(antennasByFreq[ch], coord{y, x})
			}
		}
	}

	antinodeSet := make(map[coord]bool)

	for _, coords := range antennasByFreq {
		n := len(coords)
		if n < 2 {
			continue
		}

		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				A := coords[i]
				B := coords[j]

				Ay := 2*A.y - B.y
				Ax := 2*A.x - B.x

				By := 2*B.y - A.y
				Bx := 2*B.x - A.x

				if Ay >= 0 && Ay < height && Ax >= 0 && Ax < width {
					antinodeSet[coord{Ay, Ax}] = true
				}
				if By >= 0 && By < height && Bx >= 0 && Bx < width {
					antinodeSet[coord{By, Bx}] = true
				}
			}
		}
	}

	return len(antinodeSet)
}

func Part2() int {
	lines, err := utils.ReadLines("input/day8.txt")
	if err != nil {
		panic(err)
	}

	height := len(lines)
	if height == 0 {
		return 0
	}
	width := len(lines[0])

	antennasByFreq := make(map[rune][]coord)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ch := rune(lines[y][x])
			if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
				antennasByFreq[ch] = append(antennasByFreq[ch], coord{y, x})
			}
		}
	}

	antinodeSet := make(map[coord]bool)
	lineSet := make(map[[3]int]bool)

	for _, coords := range antennasByFreq {
		n := len(coords)
		if n < 2 {
			continue
		}

		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				A := coords[i]
				B := coords[j]

				dy := B.y - A.y
				dx := B.x - A.x

				g := gcd(dy, dx)
				sy := dy / g
				sx := dx / g

				if sx < 0 {
					sx = -sx
					sy = -sy
				} else if sx == 0 && sy < 0 {
					sy = -sy
				}

				C := -sx*A.y + sy*A.x

				key := [3]int{sy, sx, C}
				if lineSet[key] {
					continue
				}
				lineSet[key] = true

				{
					Y, X := A.y, A.x
					for {
						antinodeSet[coord{Y, X}] = true
						Y += sy
						X += sx
						if Y < 0 || Y >= height || X < 0 || X >= width {
							break
						}
					}
				}

				{
					Y, X := A.y, A.x
					for {
						antinodeSet[coord{Y, X}] = true
						Y -= sy
						X -= sx
						if Y < 0 || Y >= height || X < 0 || X >= width {
							break
						}
					}
				}
			}
		}
	}

	return len(antinodeSet)
}
