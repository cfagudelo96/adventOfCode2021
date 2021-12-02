package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	argsLength = 2
	windowSize = 3
)

func main() {
	if len(os.Args) != argsLength {
		log.Fatal("Should provide just the input path as an argument")
	}

	inputPath := os.Args[1]

	mi, err := countMeasurementIncreases(inputPath)
	if err != nil {
		log.Fatalf("Couldn't count measurement increases: %v", err)
	}

	log.Printf("Measurement increases: %v", mi)

	miw, err := countMeasurementIncreasesSlidingWindow(inputPath, windowSize)
	if err != nil {
		log.Fatalf("Couldn't count measurement increases in windows: %v", err)
	}

	log.Printf("Measurement increases with sliding window (size %d): %v", windowSize, miw)
}

func countMeasurementIncreases(inputPath string) (int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, fmt.Errorf("opening file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	var c, m int

	for scanner.Scan() {
		measurementI, err := parseInputLine(scanner.Text())
		if err != nil {
			return 0, err
		}

		if m != 0 && measurementI > m {
			c++
		}

		m = measurementI
	}

	return c, nil
}

func countMeasurementIncreasesSlidingWindow(inputPath string, windowSize int) (int, error) {
	var sums []int

	file, err := os.Open(inputPath)
	if err != nil {
		return 0, fmt.Errorf("opening file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	var c, i int

	for scanner.Scan() {
		sums = append(sums, 0)

		measurementI, err := parseInputLine(scanner.Text())
		if err != nil {
			return 0, err
		}

		for j := 0; j < windowSize; j++ {
			if i-j < 0 {
				break
			}

			sums[i-j] += measurementI
		}

		if i >= windowSize && sums[i-windowSize+1] > sums[i-windowSize] {
			c++
		}
		i++
	}

	return c, nil
}

func parseInputLine(inputLine string) (int, error) {
	input, err := strconv.Atoi(inputLine)
	if err != nil {
		return 0, fmt.Errorf("parsing to int: %w", err)
	}

	return input, nil
}
