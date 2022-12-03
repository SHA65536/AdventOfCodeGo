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
		var h1, h2 = map[byte]bool{}, map[byte]bool{}
		// Getting unique letters from first half
		for i := 0; i < len(line)/2; i++ {
			h1[line[i]] = true
		}
		// Getting unique letters from second half
		for i := len(line) / 2; i < len(line); i++ {
			h2[line[i]] = true
		}
		// Summing the score of intersection
		for num := range Intersect(h1, h2) {
			res += LetScore(num)
		}
	}
	return strconv.Itoa(res), nil
}

func CommonThree(items string) (string, error) {
	var res int
	// Looping over input
	for i := 0; i < len(items); i++ {
		var h1, h2, h3 map[byte]bool
		// Getting unique letters of first line
		for h1 = map[byte]bool{}; items[i] != '\n'; i++ {
			h1[items[i]] = true
		}
		i++
		// Getting unique letters of second line
		for h2 = map[byte]bool{}; items[i] != '\n'; i++ {
			h2[items[i]] = true
		}
		i++
		// Getting unique letters of third line
		for h3 = map[byte]bool{}; i < len(items) && items[i] != '\n'; i++ {
			h3[items[i]] = true
		}
		// Summing score of intersection of all three
		for num := range Intersect(Intersect(h1, h2), h3) {
			res += LetScore(num)
		}
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

// Returns common items between two maps
func Intersect(m1, m2 map[byte]bool) map[byte]bool {
	var res = map[byte]bool{}
	for k := range m1 {
		if m2[k] {
			res[k] = true
		}
	}
	return res
}
