package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var numWords = []string{
	"1",
	"one",
	"2",
	"two",
	"3",
	"three",
	"4",
	"four",
	"5",
	"five",
	"6",
	"six",
	"7",
	"seven",
	"8",
	"eight",
	"9",
	"nine",
}

var wordToNum = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type kv struct {
	Key   string
	Value int
}

func main() {
	file, err := os.Open("/home/nf/projects/go/advent-of-code-23/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		res += lineRes(line)
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
	}

	fmt.Printf("Result: %d\n", res)
}

func partOne(scanner *bufio.Scanner) {
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		first, last := 0, 0

		for _, ch := range line {
			if unicode.IsDigit(ch) {
				int, err := strconv.Atoi(string(ch))
				if err != nil {
					log.Println(err)
				}

				if first == 0 {
					first = int
				}

				last = int
			}
		}

		lineCalibrationValue := first*10 + last

		sum += lineCalibrationValue

		fmt.Println(line)
		fmt.Printf("first: %d, last: %d\n", first, last)
		fmt.Printf("line calibration value %d\n", lineCalibrationValue)
		fmt.Printf("sum till now %d\n", sum)
	}

	fmt.Printf("Sum of calibrationvalues: %d\n", sum)
}

func lineRes(line string) int {
	firstAppearence := getFirstPos(line)
	lastAppearence := getLastPos(line)

	sortKVInc(firstAppearence)
	sortKVDec(lastAppearence)

	firstNum := firstAppearence[0].Key
	lastNum := lastAppearence[0].Key

	firstNum = getNumLiteral(firstNum)
	lastNum = getNumLiteral(lastNum)

	first, err := strconv.Atoi(firstNum)
	if err != nil {
		fmt.Println(err)
	}

	last, err := strconv.Atoi(lastNum)
	if err != nil {
		fmt.Println(err)
	}

	res := getCalibrationValue(first, last)

	return res
}

func getFirstPos(line string) []kv {
	var posInLine []kv

	for _, word := range numWords {
		pos := strings.Index(line, word)
		if pos >= 0 {
			posInLine = append(posInLine, kv{word, pos})
		}
	}

	return posInLine
}

func getLastPos(line string) []kv {
	var posInLine []kv

	for _, word := range numWords {
		pos := strings.LastIndex(line, word)
		if pos >= 0 {
			posInLine = append(posInLine, kv{word, pos})
		}
	}

	return posInLine
}

func getNumLiteral(word string) string {
	if len(word) > 1 {
		word = wordToNum[word]
	}
	return word
}

func getCalibrationValue(first, last int) int {
	return first*10 + last
}

func sortKVInc(input []kv) {
	sort.Slice(input, func(i, j int) bool {
		return input[i].Value < input[j].Value
	})
}

func sortKVDec(input []kv) {
	sort.Slice(input, func(i, j int) bool {
		return input[i].Value > input[j].Value
	})
}
