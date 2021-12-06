package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	argsLength   = 2
	safeOverlaps = 2
)

type lineDefinition struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (l lineDefinition) minMaxX() (int, int) {
	if l.x1 < l.x2 {
		return l.x1, l.x2
	}

	return l.x2, l.x1
}

func (l lineDefinition) minMaxY() (int, int) {
	if l.y1 < l.y2 {
		return l.y1, l.y2
	}

	return l.y2, l.y1
}

func (l lineDefinition) addOverlaps(overlapsMap [][]int) {
	minX, maxX := l.minMaxX()
	minY, maxY := l.minMaxY()

	switch {
	case l.isHorizontalLine():
		for i := minX; i <= maxX; i++ {
			overlapsMap[minY][i]++
		}
	case l.isVerticalLine():
		for i := minY; i <= maxY; i++ {
			overlapsMap[i][minX]++
		}
	case l.isDiagonalLine():
		xMultiplier := 1
		yMultiplier := 1

		if l.x2 < l.x1 {
			xMultiplier = -1
		}

		if l.y2 < l.y1 {
			yMultiplier = -1
		}

		for i := 0; i <= int(math.Abs(float64(l.x1-l.x2))); i++ {
			overlapsMap[l.y1+i*yMultiplier][l.x1+i*xMultiplier]++
		}
	}
}

func (l lineDefinition) isHorizontalLine() bool {
	return l.y1 == l.y2
}

func (l lineDefinition) isVerticalLine() bool {
	return l.x1 == l.x2
}

func (l lineDefinition) isDiagonalLine() bool {
	return math.Abs(float64(l.x1-l.x2)) == math.Abs(float64(l.y1-l.y2))
}

func main() {
	if len(os.Args) != argsLength {
		log.Fatal("Should provide just the input path as an argument")
	}

	inputPath := os.Args[1]

	linesDefinitions, maxCoordinates, err := parseInput(inputPath)
	if err != nil {
		log.Fatalf("Couldn't parseaddOverlaps the input: %v", err)
	}

	overlapsMap := calculateOverlapsMap(linesDefinitions, maxCoordinates)
	r := countOverlapsEqGreaterThan(safeOverlaps, overlapsMap)
	log.Printf("Safe points: %d", r)
}

func parseInput(inputPath string) ([]lineDefinition, [2]int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, [2]int{}, fmt.Errorf("opening file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	var (
		linesDefinitions []lineDefinition
		maxX, maxY       int
	)

	for scanner.Scan() {
		inputSplit := strings.Split(scanner.Text(), " -> ")

		x1, y1, err := parseCoordinates(inputSplit[0])
		if err != nil {
			return nil, [2]int{}, fmt.Errorf("parsing 1st coordinates: %w", err)
		}

		x2, y2, err := parseCoordinates(inputSplit[1])
		if err != nil {
			return nil, [2]int{}, fmt.Errorf("parsing 2nd coordinates: %w", err)
		}

		maxX = max(maxX, x1, x2)
		maxY = max(maxY, y1, y2)

		linesDefinitions = append(linesDefinitions, lineDefinition{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		})
	}

	return linesDefinitions, [2]int{maxX, maxY}, nil
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}

	if b >= a && b >= c {
		return b
	}

	return c
}

func parseCoordinates(coordinates string) (int, int, error) {
	cSplit := strings.Split(coordinates, ",")

	x, err := strconv.Atoi(cSplit[0])
	if err != nil {
		return 0, 0, fmt.Errorf("parsing coordinate X: %w", err)
	}

	y, err := strconv.Atoi(cSplit[1])
	if err != nil {
		return 0, 0, fmt.Errorf("parsing coordinate Y: %w", err)
	}

	return x, y, nil
}

func calculateOverlapsMap(linesDefinitions []lineDefinition, maxCoordinates [2]int) [][]int {
	overlapsMap := initializeOverlapsMap(maxCoordinates)

	for _, lineDef := range linesDefinitions {
		lineDef.addOverlaps(overlapsMap)
	}

	return overlapsMap
}

func countOverlapsEqGreaterThan(n int, overlapsMap [][]int) int {
	var c int

	for i := 0; i < len(overlapsMap); i++ {
		for j := 0; j < len(overlapsMap[i]); j++ {
			if overlapsMap[i][j] >= n {
				c++
			}
		}
	}

	return c
}

func initializeOverlapsMap(maxCoordinates [2]int) [][]int {
	overlapsMap := make([][]int, maxCoordinates[1]+1)
	for i := 0; i < len(overlapsMap); i++ {
		overlapsMap[i] = make([]int, maxCoordinates[0]+1)
	}

	return overlapsMap
}
