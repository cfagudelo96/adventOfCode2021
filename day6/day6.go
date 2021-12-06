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
	argsLength          = 2
	initialFishState    = 8
	resetFishState      = 6
	simulationDaysPart1 = 80
	simulationDaysPart2 = 256
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

	fishesStateMap := fishListToStateMap(fishes)

	for i := 0; i < simulationDaysPart1; i++ {
		simulateDayWithStateMap(fishesStateMap)
	}

	log.Printf("Number of fishes p1: %d", countFishesFromStateMap(fishesStateMap))

	for i := simulationDaysPart1; i < simulationDaysPart2; i++ {
		simulateDayWithStateMap(fishesStateMap)
	}

	log.Printf("Number of fishes p2: %d", countFishesFromStateMap(fishesStateMap))
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

/* Previous naive implementation
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
*/

func simulateDayWithStateMap(stateMap map[int]int) {
	fishInState0 := stateMap[0]
	previousStateCount := stateMap[0]

	for i := initialFishState; i >= 0; i-- {
		stateMap[i], previousStateCount = previousStateCount, stateMap[i]
		if i == resetFishState {
			stateMap[i] += fishInState0
		}
	}
}

func fishListToStateMap(fishes []int) map[int]int {
	stateMap := make(map[int]int)

	for _, fish := range fishes {
		stateMap[fish]++
	}

	return stateMap
}

func countFishesFromStateMap(stateMap map[int]int) int {
	var c int

	for _, v := range stateMap {
		c += v
	}

	return c
}
