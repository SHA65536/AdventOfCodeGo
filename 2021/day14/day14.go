package day14

import (
	"strconv"
)

// https://adventofcode.com/2021/day/14

func PolymerGenerate(guide string, steps int) (string, error) {
	var i int
	var occurences = map[[2]byte]uint64{}
	var freq = map[byte]uint64{}
	var instructions = map[[2]byte]byte{}
	var prev byte
	// Loading starting configuration
	for ; guide[i] != '\n'; i++ {
		if prev != 0 {
			occurences[[2]byte{prev, guide[i]}]++
		}
		prev = guide[i]
	}
	i++
	// Loading instructions
	for i++; i < len(guide); i += 8 {
		instructions[[2]byte{guide[i], guide[i+1]}] = guide[i+6]
	}
	// Creating new polymer
	for s := 0; s < steps; s++ {
		var newMap = map[[2]byte]uint64{}
		for p, times := range occurences {
			if val, ok := instructions[p]; ok {
				newMap[[2]byte{p[0], val}] += times
				newMap[[2]byte{val, p[1]}] += times
			} else {
				newMap[p] += times
			}
		}
		occurences = newMap
	}
	// Calculating frequency
	for p, v := range occurences {
		freq[p[0]] += v
		freq[p[1]] += v
	}
	freq[guide[0]]++
	freq[prev]++
	// Calculating min and max
	var most, least uint64 = 0, ^uint64(0)
	for _, v := range freq {
		if v > most {
			most = v
		}
		if v < least {
			least = v
		}
	}
	return strconv.FormatUint(most/2-least/2, 10), nil
}
