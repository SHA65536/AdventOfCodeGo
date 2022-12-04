package day06

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/3

func CommonItem(items string) (string, error) {
	var res int
	// Looping over lines
	for _, line := range strings.Split(items, "\n") {
		var h1, h2 uint64
		// Getting unique letters from first half
		for i := 0; i < len(line)/2; i++ {
			h1 |= 1 << (LetScore(line[i]) - 1)
		}
		// Getting unique letters from second half
		for i := len(line) / 2; i < len(line); i++ {
			h2 |= 1 << (LetScore(line[i]) - 1)
		}
		// Getting score of intersection
		res += ScoreLet(h1 & h2)
	}
	return strconv.Itoa(res), nil
}

func CommonThree(items string) (string, error) {
	var res int
	// Looping over input
	for i := 0; i < len(items); i++ {
		var h1, h2, h3 uint64
		// Getting unique letters of first line
		for ; items[i] != '\n'; i++ {
			h1 |= 1 << (LetScore(items[i]) - 1)
		}
		i++
		// Getting unique letters of second line
		for ; items[i] != '\n'; i++ {
			h2 |= 1 << (LetScore(items[i]) - 1)
		}
		i++
		// Getting unique letters of third line
		for ; i < len(items) && items[i] != '\n'; i++ {
			h3 |= 1 << (LetScore(items[i]) - 1)
		}
		// Summing score of intersection of all three
		res += ScoreLet(h1 & h2 & h3)
	}
	return strconv.Itoa(res), nil
}

// Score letter according to guide , a-1 z-26 A-27 Z-52
func LetScore(b byte) int {
	if b >= 'a' && b <= 'z' {
		return int(b-'a') + 1
	} else {
		return int(b-'A') + 27
	}
}

// Get letter score from uint64
func ScoreLet(num uint64) int {
	var score int = 1
	for num&1 != 1 {
		num >>= 1
		score++
	}
	return score
}
