package main

import (
	"bufio"
	"fmt"
	"github.com/burybind/adventofcode2021"
	"log"
	"os"
	"strconv"
	"strings"
)

type hasBeenCalled bool
type bingoCard [5]bingoRow
type bingoRow []*slot
type slot struct {
	num int
	hasBeenCalled
}

func main() {
	f, err := os.Open("/Users/brendan.ashton/go/src/github.com/burybind/adventofcode2021/day3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	calledNums := getDrawnNumbers(scanner)
	bingoCards := getBingoCards(scanner)
	var cardWithBingo *bingoCard
	for i, num := range calledNums {
		markUpAllBingoCards(bingoCards, num)
		cardWithBingo = checkForBingos(bingoCards)
		if cardWithBingo != nil {
			sum := cardWithBingo.sumAllUnMarked()
			product := sum * num
			fmt.Printf("winning card won in %d steps. Unmarked sum: %d. Last called num: %d. Product: %d\n", i, sum, num, product)
			break
		}
	}
}

func getDrawnNumbers(scanner *bufio.Scanner) []int {
	scanner.Scan()

	nums := convertToIntArray(scanner.Text(), ",")
	return nums
}

func convertToBingoRow(str string) bingoRow {
	arr := strings.Split(str, " ")

	m := bingoRow{}

	for _, a := range arr {
		a = strings.TrimSpace(a)
		if a == "" {
			continue
		}
		num, err := strconv.Atoi(a)
		adventofcode2021.Check(err)
		m = append(m, &slot{num, false})
	}
	return m
}

func convertToIntArray(str string, delim string) []int {
	var converted []int
	arr := strings.Split(str, delim)

	for _, a := range arr {
		num, err := strconv.Atoi(a)
		adventofcode2021.Check(err)
		converted = append(converted, num)
	}
	return converted
}

func getBingoCards(scanner *bufio.Scanner) []*bingoCard {
	var bingoCards []*bingoCard
	var card *bingoCard
	var i int

	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			bingoCards = append(bingoCards, card)
			card = &bingoCard{}
			i = 0
			continue
		}
		card[i] = convertToBingoRow(txt)
		i++
	}

	bingoCards = append(bingoCards, card)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return bingoCards
}

func checkForBingos(cards []*bingoCard) *bingoCard {
	for _, card := range cards {
		if card == nil {
			continue
		}
		if card.hasRowBingo() || card.hasColumnBingo() {
			return card
		}
	}
	return nil
}

func markUpAllBingoCards(cards []*bingoCard, num int) {
	for _, card := range cards {
		if card == nil {
			continue
		}
		card.markNumber(num)
	}
}

func (c *bingoCard) markNumber(calledNum int) {
	for _, row := range c {
		for _, s := range row {
			if s.num == calledNum {
				s.hasBeenCalled = true
				break
			}
		}
	}
}

func (c *bingoCard) hasRowBingo() bool {
	for _, row := range c {
		i := 0
		for _, v := range row {
			if v.hasBeenCalled {
				i++
			}
		}
		if i == 5 {
			return true
		}
	}
	return false
}

func (c *bingoCard) hasColumnBingo() bool {
	for i := 0; i < 5; i++ {
		if c[0][i].hasBeenCalled && c[1][i].hasBeenCalled && c[2][i].hasBeenCalled && c[3][i].hasBeenCalled && c[4][i].hasBeenCalled {
			return true
		}
	}
	return false
}

func (c *bingoCard) sumAllUnMarked() int {
	var runningTotal int
	for _, row := range c {
		for _, slot := range row {
			if !slot.hasBeenCalled {
				runningTotal += slot.num
			}
		}
	}
	return runningTotal
}
