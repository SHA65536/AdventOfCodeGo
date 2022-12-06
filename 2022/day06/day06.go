package day06

import "strconv"

// https://adventofcode.com/2022/day/6

func StartOfPacket(signal string, cons int) (string, error) {
	var i int
	var past [26]int
	// Loading the first n signals
	for i = 0; i < cons; i++ {
		past[signal[i]-'a']++
	}
	// Loading the rest
	for i = cons; i < len(signal); i++ {
		// Checking if they're all unique
		if AllOnes(past) {
			break
		}
		// Removing the first singal
		past[signal[i-cons]-'a']--
		// Adding the next signal
		past[signal[i]-'a']++
	}
	return strconv.Itoa(i), nil
}

func AllOnes(in [26]int) bool {
	for _, v := range in {
		if v > 1 {
			return false
		}
	}
	return true
}
