package day04

import (
	"adventofcode/helper"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		var cur int
		var win, got = parseLine(line)
		for _, num := range got {
			if win[num] {
				if cur == 0 {
					cur = 1
				} else {
					cur *= 2
				}
			}
		}
		res += cur
	}

	return strconv.Itoa(res), nil
}

func parseLine(line string) (win [100]bool, got []int) {
	var past bool
	for _, word := range strings.Fields(line)[1:] {
		if word == "|" {
			past = true
			continue
		}
		if past {
			got = append(got, helper.MustConvNum(word))
		} else {
			win[helper.MustConvNum(word)] = true
		}
	}

	return win, got
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var cards = make([][2]int, 0, 300)

	for line := range input.IterateLines {
		var win, got = parseLine(line)
		var cur int
		for _, num := range got {
			if win[num] {
				cur++
			}
		}
		cards = append(cards, [2]int{cur, 1})
	}

	for i, card := range cards {
		for j := i + 1; j < len(cards) && j-i <= card[0]; j++ {
			cards[j][1] += card[1]
		}
		res += card[1]
	}

	return strconv.Itoa(res), nil
}
