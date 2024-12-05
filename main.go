package main

import (
	"AOC2024/day1"
	"AOC2024/day2"
	"AOC2024/day3"
	"AOC2024/day4"
	"AOC2024/day5"
	"AOC2024/utils"
	"fmt"
)

func main() {
	fmt.Printf("Day 1, Part 1: %d\n\n", utils.Benchmark(day1.Part1, 100))
	fmt.Printf("Day 1, Part 2: %d\n\n", utils.Benchmark(day1.Part2, 100))
	fmt.Printf("Day 2, Part 1: %d\n\n", utils.Benchmark(day2.Part1, 100))
	fmt.Printf("Day 2, Part 2: %d\n\n", utils.Benchmark(day2.Part2, 100))
	fmt.Printf("Day 3, Part 1: %d\n\n", utils.Benchmark(day3.Part1, 100))
	fmt.Printf("Day 3, Part 2: %d\n\n", utils.Benchmark(day3.Part2, 100))
	fmt.Printf("Day 4, Part 1: %d\n\n", utils.Benchmark(day4.Part1, 100))
	fmt.Printf("Day 4, Part 2: %d\n\n", utils.Benchmark(day4.Part2, 100))
	fmt.Printf("Day 5, Part 1: %d\n\n", utils.Benchmark(day5.Part1, 100))
	fmt.Printf("Day 5, Part 2: %d\n\n", utils.Benchmark(day5.Part2, 1))
}
