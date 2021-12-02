package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	argsLength     = 2
	inputSplitSize = 2
)

var errInvalidInput = errors.New("invalid input")

func main() {
	if len(os.Args) != argsLength {
		log.Fatal("Should provide just the input path as an argument")
	}

	inputPath := os.Args[1]

	xPosition, zPosition, err := calculateXZPosition(inputPath)
	if err != nil {
		log.Fatalf("Couldn't find the position: %v", err)
	}

	log.Printf("Current position (X, Z) is: (%d, %d)", xPosition, zPosition)
	log.Printf("X*Z is: %d", xPosition*zPosition)

	xPosition, zPosition, err = calculateXZPositionWithAim(inputPath)
	if err != nil {
		log.Fatalf("Couldn't find the position with aim: %v", err)
	}

	log.Printf("Current position with aim (X, Z) is: (%d, %d)", xPosition, zPosition)
	log.Printf("X*Z with aim is: %d", xPosition*zPosition)
}

func calculateXZPosition(inputPath string) (int, int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, 0, fmt.Errorf("opening file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	var x, z int

	for scanner.Scan() {
		direction, amount, err := parseInputLine(scanner.Text())
		if err != nil {
			return x, z, fmt.Errorf("invalid line: %w", err)
		}

		switch direction {
		case "up":
			z -= amount
		case "down":
			z += amount
		default:
			x += amount
		}
	}

	return x, z, nil
}

func calculateXZPositionWithAim(inputPath string) (int, int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, 0, fmt.Errorf("opening file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	var x, z, aim int

	for scanner.Scan() {
		direction, amount, err := parseInputLine(scanner.Text())
		if err != nil {
			return x, z, fmt.Errorf("invalid line: %w", err)
		}

		switch direction {
		case "up":
			aim -= amount
		case "down":
			aim += amount
		default:
			x += amount
			z += aim * amount
		}
	}

	return x, z, nil
}

func parseInputLine(inputLine string) (string, int, error) {
	inputSplit := strings.Split(inputLine, " ")
	if len(inputSplit) != inputSplitSize {
		return "", 0, errInvalidInput
	}

	direction := inputSplit[0]

	amount, err := strconv.Atoi(inputSplit[1])
	if err != nil {
		return "", 0, fmt.Errorf("invalid input: %w", err)
	}

	return direction, amount, nil
}
