package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){
	data, err := ioutil.ReadFile("/Users/brendan.ashton/go/src/github.com/burybind/adventofcode2021/day1/input.txt")
	check(err)

	lines := strings.Split(string(data), "\n")
	increases := 0
	allNums:= []int{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		currNum, err := strconv.Atoi(line)
		check(err)

		allNums = append(allNums, currNum)
	}

	for i, _ := range allNums {
		if i > 3 {
			firstCurrentIndex := i - 3
			firstPrevIndex := i - 4
			lastPrevIndex := i-1
			prevSum := sum(allNums[firstPrevIndex:lastPrevIndex])
			currSum := sum(allNums[firstCurrentIndex:i])
			if currSum > prevSum {
				increases++
			}
		}
	}

	fmt.Printf("num of increases: %d\n", increases)
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
