package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var validSymbols = "#$%&*+-/@="

var input []string

func main() {
	fmt.Println("Solution Day 3")
	scanner := getFileScanner()
	solvePartOne(scanner)
}

func solvePartOne(scanner *bufio.Scanner) {
	initInputAsArray(scanner)

	for _, line := range input {
		fmt.Println(line)
	}

	for i, line := range input {
		for j, symbol := range line {
			symbolAsNumber, err := strconv.Atoi(string(symbol))
			if err == nil {
				isPartNumber := hasAdjacentSymbol(i, j)

				fmt.Printf("Number %d is part number: %t\n", symbolAsNumber, isPartNumber)
				// TODO: For now we checked for a single digit if it has adjacent symbols, now concat following digits to numbers and check for them
			}
		}
	}
}

func hasAdjacentSymbol(line, pos int) bool {
	lineAbove := line > 0
	lineBelow := line < len(input)-1

	if (lineAbove && checkHorizontal(line-1, pos)) || checkHorizontal(line, pos) || (lineBelow && checkHorizontal(line+1, pos)) {
		return true
	}

	return false
}

func checkHorizontal(lineNumber, pos int) bool {
	line := input[lineNumber]
	symbolBefore := pos > 0
	symbolAfter := pos < len(line)-1

	if (symbolBefore && isValidSymbol(line[pos-1])) || isValidSymbol(line[pos]) || (symbolAfter && isValidSymbol(line[pos+1])) {
		return true
	}

	return false
}

func isValidSymbol(symbol byte) bool {
	return strings.Contains(validSymbols, string(symbol))
}

func getFileScanner() *bufio.Scanner {
	file, err := os.Open("/home/nf/projects/go/advent-of-code-23/day3/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
}

func initInputAsArray(scanner *bufio.Scanner) {

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
}
