package main

import (
	"AOC2024/day1"
	"AOC2024/day2"
	"AOC2024/day3"
	"AOC2024/utils"
	"fmt"
)

func main() {
	utils.Benchmark(func() { fmt.Printf("Day 1, Part 1: %d\n", day1.Part1()) })
	utils.Benchmark(func() { fmt.Printf("Day 1, Part 2: %d\n", day1.Part2()) })
	utils.Benchmark(func() { fmt.Printf("Day 2, Part 1: %d\n", day2.Part1()) })
	utils.Benchmark(func() { fmt.Printf("Day 2, Part 2: %d\n", day2.Part2()) })
	utils.Benchmark(func() { fmt.Printf("Day 3, Part 1: %d\n", day3.Part1()) })
	utils.Benchmark(func() { fmt.Printf("Day 3, Part 2: %d\n", day3.Part2()) })
}
