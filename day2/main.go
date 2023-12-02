package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, err := os.Open("/home/nf/projects/go/advent-of-code-23/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	res := getPowerOfCubes(scanner)

	if scanner.Err() != nil {
		log.Println(scanner.Err())
	}

	fmt.Printf("Result: %d\n", res)
}

func getPowerOfCubes(scanner *bufio.Scanner) int {
	res := 0

	for scanner.Scan() {
		game := scanner.Text()
		minCubes := getMinCubesForGame(game)

		powerOfCubes := 1
		for _, v := range minCubes {
			powerOfCubes *= v
		}

		res += powerOfCubes
	}

	return res
}

func getMinCubesForGame(game string) map[string]int {
	minCubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	rounds := strings.Split(game, ";")

	for _, round := range rounds {
		for k := range cubes {
			colorPos := strings.Index(round, k)
			if colorPos < 0 {
				continue
			}
			count := round[colorPos-3 : colorPos]
			count = strings.TrimSpace(count)
			countAsInt := toInt(count)

			if minCubes[k] < countAsInt {
				minCubes[k] = countAsInt
			}
		}
	}

	return minCubes
}

func getValidGamesIndexSum(scanner *bufio.Scanner) int {
	res := 0
	game := 0

	for scanner.Scan() {
		game++
		line := scanner.Text()
		if isGamePossible(line) {
			res += game
		}
	}

	return res
}

func isGamePossible(game string) bool {
	rounds := strings.Split(game, ";")

	for _, round := range rounds {
		if !isRoundPossible(round) {
			return false
		}
	}

	return true
}

func isRoundPossible(round string) bool {
	for k, v := range cubes {
		colorPos := strings.Index(round, k)
		if colorPos < 0 {
			continue
		}
		count := round[colorPos-3 : colorPos]
		count = strings.TrimSpace(count)
		countAsInt := toInt(count)

		if countAsInt > v {
			return false
		}
	}
	return true
}

func toInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
