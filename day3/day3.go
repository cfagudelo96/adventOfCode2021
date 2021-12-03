package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	argsLength = 2
	binaryBase = 2
)

type submarineDiagnostic struct {
	gammaRate             int
	epsilonRate           int
	oxygenGeneratorRating int
	co2ScrubberRating     int
}

func main() {
	if len(os.Args) != argsLength {
		log.Fatal("Should provide just the input path as an argument")
	}

	inputPath := os.Args[1]

	subD, err := calculateSubmarineDiagnostic(inputPath)
	if err != nil {
		log.Fatalf("Couln't calculate the value of gamma and epsilon: %v", err)
	}

	log.Printf("Gamma: %d, Epsilon: %d", subD.gammaRate, subD.epsilonRate)
	log.Printf("The power consumption is: %d", subD.gammaRate*subD.epsilonRate)
	log.Printf("Oxygen generator rating: %d, CO2 scrubber rating: %d", subD.oxygenGeneratorRating, subD.co2ScrubberRating)
	log.Printf("Life support rating: %d", subD.oxygenGeneratorRating*subD.co2ScrubberRating)
}

func calculateSubmarineDiagnostic(inputPath string) (submarineDiagnostic, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return submarineDiagnostic{}, fmt.Errorf("opening file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	var (
		gamma, epsilon int
		bitsCount      [][2]int
		numbers        [][]rune
	)

	for scanner.Scan() {
		inputLine := scanner.Text()
		numbers = append(numbers, []rune(inputLine))

		if bitsCount == nil {
			bitsCount = make([][2]int, len([]rune(inputLine)))
		}

		for i, r := range inputLine {
			switch r {
			case '0':
				bitsCount[i][0]++
			default:
				bitsCount[i][1]++
			}
		}
	}

	for i, bitCount := range bitsCount {
		bitSum := int(math.Pow(binaryBase, float64(len(bitsCount)-i-1)))

		if bitCount[1] > bitCount[0] {
			gamma += bitSum
		} else {
			epsilon += bitSum
		}
	}

	return submarineDiagnostic{
		gammaRate:             gamma,
		epsilonRate:           epsilon,
		oxygenGeneratorRating: calculateOxygenGeneratorRating(numbers),
		co2ScrubberRating:     calculateCO2ScrubberRating(numbers),
	}, nil
}

func calculateOxygenGeneratorRating(numbers [][]rune) int {
	for i := 0; i < len(numbers[0]) && len(numbers) > 1; i++ {
		count := getBitsCountAtIndex(numbers, i)
		if count[1] >= count[0] {
			numbers = findMatchingNumbers(numbers, i, '1')
		} else {
			numbers = findMatchingNumbers(numbers, i, '0')
		}
	}

	return binaryRunesToInt(numbers[0])
}

func calculateCO2ScrubberRating(numbers [][]rune) int {
	for i := 0; i < len(numbers[0]) && len(numbers) > 1; i++ {
		count := getBitsCountAtIndex(numbers, i)
		if count[1] >= count[0] {
			numbers = findMatchingNumbers(numbers, i, '0')
		} else {
			numbers = findMatchingNumbers(numbers, i, '1')
		}
	}

	return binaryRunesToInt(numbers[0])
}

func getBitsCountAtIndex(numbers [][]rune, index int) [2]int {
	var bitsCount [2]int

	for _, number := range numbers {
		if number[index] == '0' {
			bitsCount[0]++
		} else {
			bitsCount[1]++
		}
	}

	return bitsCount
}

func binaryRunesToInt(runes []rune) int {
	var n int

	for i, r := range runes {
		if r == '1' {
			n += int(math.Pow(binaryBase, float64(len(runes)-i-1)))
		}
	}

	return n
}

func findMatchingNumbers(numbers [][]rune, position int, match rune) [][]rune {
	var matchingNumbers [][]rune

	for i := 0; i < len(numbers); i++ {
		if numbers[i][position] == match {
			matchingNumbers = append(matchingNumbers, numbers[i])
		}
	}

	return matchingNumbers
}
