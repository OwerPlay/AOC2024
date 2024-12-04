package utils

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// ReadLines reads all lines from a file and returns them as a slice of strings.
func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Benchmark(f func() int, iterations int) int {
	result := f()

	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = f()
	}
	elapsed := time.Since(start)

	fmt.Printf("Execution time (%d iterations): %s\n", iterations, elapsed)
	fmt.Printf("Avg. execution time: %s\n", elapsed/time.Duration(iterations))

	return result
}
