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
	argsLength = 2
)

type bingoBoard struct {
	won           bool
	values        [5][5]bool
	valuePosition map[int][2]int
}

func (b bingoBoard) isWinningRow(n int) bool {
	winningRow := true
	for i := 0; i < len(b.values[n]) && winningRow; i++ {
		winningRow = winningRow && b.values[n][i]
	}

	return winningRow
}

func (b bingoBoard) isWinningColumn(n int) bool {
	winningColumn := true
	for i := 0; i < len(b.values) && winningColumn; i++ {
		winningColumn = winningColumn && b.values[i][n]
	}

	return winningColumn
}

func (b bingoBoard) sumUnmarkedNumbers() int {
	var sum int

	for number, coordinates := range b.valuePosition {
		if !b.values[coordinates[0]][coordinates[1]] {
			sum += number
		}
	}

	return sum
}

func main() {
	if len(os.Args) != argsLength {
		log.Fatal("Should provide just the input path as an argument")
	}

	inputPath := os.Args[1]

	numbers, boards, err := parseInput(inputPath)
	if err != nil {
		log.Fatalf("Couldn't parse the input: %v", err)
	}

	winningBoard, winningNumber := findLastWinningBoard(numbers, boards)
	sum := winningBoard.sumUnmarkedNumbers()
	log.Printf("The sum of the unmarked numbers in the winning board is: %d", sum)
	log.Printf("The winning number was: %d", winningNumber)
	log.Printf("Sum*winning number: %d", sum*winningNumber)
}

func parseInput(inputPath string) ([]int, []*bingoBoard, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, nil, fmt.Errorf("opening file: %w", err)
	}

	var (
		numbersCalled []int
		boards        []*bingoBoard
	)

	scanner := bufio.NewScanner(file)

	var (
		numbersCalledRead bool
		i, j              int
	)

	currentBoard := &bingoBoard{
		valuePosition: make(map[int][2]int),
	}

	for scanner.Scan() {
		inputLine := scanner.Text()

		if !numbersCalledRead {
			numbersCalled, err = parseNumbersCalled(inputLine)
			if err != nil {
				return nil, nil, fmt.Errorf("reading the numbers called: %w", err)
			}

			numbersCalledRead = true

			scanner.Scan() // Read empty line after numbers called.

			continue
		}

		if inputLine == "" {
			boards = append(boards, currentBoard)
			currentBoard = &bingoBoard{
				valuePosition: make(map[int][2]int),
			}
			i = 0
			j = 0

			continue
		}

		rowSplit := strings.Split(inputLine, " ")
		for _, ns := range rowSplit {
			if ns == "" {
				continue
			}

			n, err := strconv.Atoi(ns)
			if err != nil {
				return nil, nil, fmt.Errorf("parsing board entry: %w", err)
			}

			currentBoard.valuePosition[n] = [2]int{i, j}
			j++
		}

		j = 0
		i++
	}

	boards = append(boards, currentBoard)

	return numbersCalled, boards, nil
}

func parseNumbersCalled(inputLine string) ([]int, error) {
	numbersCalledSplit := strings.Split(inputLine, ",")
	numbersCalled := make([]int, len(numbersCalledSplit))

	for i, v := range numbersCalledSplit {
		number, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("parsing number with index %d: %w", i, err)
		}

		numbersCalled[i] = number
	}

	return numbersCalled, nil
}

func findWinningBoard(numbers []int, boards []*bingoBoard) (*bingoBoard, int) {
	for _, n := range numbers {
		for _, board := range boards {
			if coordinates, ok := board.valuePosition[n]; ok {
				board.values[coordinates[0]][coordinates[1]] = true
				if board.isWinningRow(coordinates[0]) || board.isWinningColumn(coordinates[1]) {
					return board, n
				}
			}
		}
	}

	return nil, 0
}

func findLastWinningBoard(numbers []int, boards []*bingoBoard) (*bingoBoard, int) {
	var winningBoardsCount int

	for _, n := range numbers {
		for _, board := range boards {
			if coordinates, ok := board.valuePosition[n]; ok {
				board.values[coordinates[0]][coordinates[1]] = true
				if board.isWinningRow(coordinates[0]) || board.isWinningColumn(coordinates[1]) {
					if !board.won {
						board.won = true
						winningBoardsCount++

						if winningBoardsCount == len(boards) {
							return board, n
						}
					}
				}
			}
		}
	}

	return nil, 0
}
