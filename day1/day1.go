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

type window struct {
	sum         int
	currentSize int
}

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
	windows := make([]window, 0)

	file, err := os.Open(inputPath)
	if err != nil {
		return 0, fmt.Errorf("opening file: %w", err)
	}

	var c, initialWI, currentWI int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		windows = append(windows, window{})

		measurementI, err := parseInputLine(scanner.Text())
		if err != nil {
			return 0, err
		}

		for i := 0; i <= currentWI-initialWI; i++ {
			windows[initialWI+i].sum += measurementI
			windows[initialWI+i].currentSize++
		}

		if windows[initialWI].currentSize == windowSize {
			if initialWI >= 1 && windows[initialWI].sum > windows[initialWI-1].sum {
				c++
			}

			initialWI++
		}
		currentWI++
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
