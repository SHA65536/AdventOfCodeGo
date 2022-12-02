package day06

import (
	"strconv"
)

// https://adventofcode.com/2021/day/6

func FishPop(initial string) (string, error) {
	var popCur = new([9]uint64)
	var res uint64
	// Loading the population
	for i := 0; i < len(initial); i++ {
		if initial[i] >= '0' && initial[i] <= '9' {
			popCur[initial[i]-'0']++
		}
	}
	for i := 0; i < 80; i++ {
		// popNext is next turn population
		var popNext = new([9]uint64)
		for j := 8; j >= 1; j-- {
			// All fish decrement the timer
			popNext[j-1] = popCur[j]
		}
		// Replicating fish go back to 6
		popNext[6] += popCur[0]
		// New fish start at 8
		popNext[8] = popCur[0]
		// Switch to next turn
		popCur = popNext
	}
	// Sum all of them up
	for i := range popCur {
		res += (*popCur)[i]
	}
	return strconv.FormatUint(res, 10), nil
}

// Same deal just with 256 days
func FishPopMore(initial string) (string, error) {
	var popCur = new([9]uint64)
	var res uint64
	// Loading the population
	for i := 0; i < len(initial); i++ {
		if initial[i] >= '0' && initial[i] <= '9' {
			popCur[initial[i]-'0']++
		}
	}
	for i := 0; i < 256; i++ {
		// popNext is next turn population
		var popNext = new([9]uint64)
		for j := 8; j >= 1; j-- {
			// All fish decrement the timer
			popNext[j-1] = popCur[j]
		}
		// Replicating fish go back to 6
		popNext[6] += popCur[0]
		// New fish start at 8
		popNext[8] = popCur[0]
		// Switch to next turn
		popCur = popNext
	}
	// Sum all of them up
	for i := range popCur {
		res += (*popCur)[i]
	}
	return strconv.FormatUint(res, 10), nil
}
