package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	argsLength       = 2
	initialFishState = 8
	simulationDays   = 256
)

func main() {
	if len(os.Args) != argsLength {
		log.Fatal("Should provide just the input path as an argument")
	}

	inputPath := os.Args[1]

	fishes, err := parseInput(inputPath)
	if err != nil {
		log.Fatalf("Couldn't read the initial input: %v", err)
	}

	for i := 0; i < simulationDays; i++ {
		fishes = simulateDay(fishes)
	}

	log.Printf("Number of fishes: %d", len(fishes))
}

func parseInput(inputPath string) ([]int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	var fishes []int

	for scanner.Scan() {
		inputSplit := strings.Split(scanner.Text(), ",")
		for _, nString := range inputSplit {
			n, err := strconv.Atoi(nString)
			if err != nil {
				return nil, fmt.Errorf("parsing number entry: %w", err)
			}

			fishes = append(fishes, n)
		}
	}

	return fishes, nil
}

func simulateDay(fishes []int) []int {
	newFishes := make([]int, len(fishes))

	for i, fish := range fishes {
		if fish == 0 {
			newFishes[i] = 6
			newFishes = append(newFishes, initialFishState)
		} else {
			newFishes[i] = fish - 1
		}
	}

	return newFishes
}
