package main

import (
	"fmt"
	"github.com/burybind/adventofcode2021"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}
type line struct {
	start, finish point
}

func main() {
	data, err := ioutil.ReadFile("/Users/brendan.ashton/go/src/github.com/burybind/adventofcode2021/day4/input.txt")
	adventofcode2021.Check(err)

	txtlines := strings.Split(strings.TrimSpace(string(data)), "\n")
	allLines := []line{}
	for _, line := range txtlines {
		allLines = append(allLines, convertTxtToLine(line))
	}

	fmt.Printf("lines - count: %d start/finish: %v", len(allLines), allLines)
}

func convertTxtToLine(txtline string) line {
	pointsStr := strings.Split(txtline, "->")
	return line{
		start:  coordinateStrToPoint(pointsStr[0]),
		finish: coordinateStrToPoint(pointsStr[1]),
	}
}

func coordinateStrToPoint(coordinate string) point {
	xyStr := strings.Split(coordinate, ",")
	x, _ := strconv.Atoi(xyStr[0])
	y, _ := strconv.Atoi(xyStr[1])

	return point{
		x: x,
		y: y,
	}
}
