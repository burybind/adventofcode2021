package main

import (
	"fmt"
	"github.com/burybind/adventofcode2021"
	"io/ioutil"
	"strconv"
	"strings"
)

func main(){
	data, err := ioutil.ReadFile("/Users/brendan.ashton/go/src/github.com/burybind/adventofcode2021/day2/input.txt")
	adventofcode2021.Check(err)

	lines := strings.Split(string(data), "\n")

	gamma := [12]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	epsilon := [12]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	linecount := float64(0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		linecount++
		for i := 0; i < len(line); i++ {
			val, _ := strconv.Atoi(string(line[i]))
			gamma[i] += float64(val)
		}
	}

	for i, f := range gamma {
		avg := f/(linecount)
		if avg >= float64(0.5) {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}


	fmt.Printf("gamma:   %v\n", gamma)
	fmt.Printf("epsilon: %v\n", epsilon)

	gammaStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(gamma)), ""), "[]")
	epsilonStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(epsilon)), ""), "[]")
	fmt.Printf("gammaStr:   %v\n", gammaStr)
	fmt.Printf("epsilonStr: %v\n", epsilonStr)

	gammaInt, _ := strconv.ParseInt(gammaStr, 2, 64)
	fmt.Printf("gammaInt: %v\n", gammaInt)
	epsilonInt, _ := strconv.ParseInt(epsilonStr, 2, 64)
	fmt.Printf("gammaInt: %v\n", epsilonInt)
}
