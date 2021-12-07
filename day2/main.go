package main

import (
	"fmt"
	"github.com/burybind/adventofcode2021"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("/Users/brendan.ashton/go/src/github.com/burybind/adventofcode2021/day2/.input.txt")
	adventofcode2021.Check(err)

	lines := strings.Split(string(data), "\n")

	mostLines := lines
	leastLines := lines

	index := 0
	for len(mostLines) > 1 {
		most, _ := mostAndLeastCommonBitAtIndex(mostLines, index)
		mostLines = filterOutNonMatches(mostLines, strconv.Itoa(most), index)
		index++
	}

	index = 0
	for len(leastLines) > 1 {
		_, least := mostAndLeastCommonBitAtIndex(leastLines, index)
		leastLines = filterOutNonMatches(leastLines, strconv.Itoa(least), index)
		index++
	}

	fmt.Printf("most lines: %v\n", mostLines)
	fmt.Printf("least lines: %v\n", leastLines)

	oxygenRating, scrubberRating := binaryToInt(mostLines[0]), binaryToInt(leastLines[0])
	fmt.Printf("oxygen: %v\n", oxygenRating)
	fmt.Printf("scrubber: %v\n", scrubberRating)
}

func filterOutNonMatches(lines []string, filter string, index int) []string {
	filteredList := []string{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		if string(line[index]) == filter {
			filteredList = append(filteredList, line)
		}
	}
	return filteredList
}

func mostAndLeastCommonBitAtIndex(lines []string, index int) (int, int) {
	gamma := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	epsilon := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	linecount := float64(0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		linecount++
		for i := 0; i < len(line); i++ {
			val, _ := strconv.Atoi(string(line[i]))
			gamma[i] += val
		}
	}

	for i, f := range gamma {
		avg := float64(f) / (linecount)
		if avg >= float64(0.5) {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}

	return gamma[index], epsilon[index]
}

func binaryToInt(val string) int64 {
	gammaInt, _ := strconv.ParseInt(val, 2, 64)
	return gammaInt
}
