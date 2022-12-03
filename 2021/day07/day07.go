package day07

import (
	"sort"
	"strconv"
)

// https://adventofcode.com/2021/day/7

func HorizontalAlign(crabs string) (string, error) {
	var positions []int
	var median, res int
	// Loading crab positions
	for i := 0; i < len(crabs); i++ {
		var cur int
		for cur = 0; i < len(crabs) && crabs[i] != ','; i++ {
			cur = cur*10 + int(crabs[i]-'0')
		}
		positions = append(positions, cur)
	}
	// Sorting to find median
	sort.Ints(positions)
	// Finding median
	if len(positions)%2 == 0 {
		median = (positions[len(positions)/2] + positions[len(positions)/2-1]) / 2
	} else {
		median = positions[len(positions)/2]
	}
	// Finding distance from median for each crab
	for i := range positions {
		res += absDist(median, positions[i])
	}
	return strconv.Itoa(res), nil
}

func HorizontalAlignComplex(crabs string) (string, error) {
	var res1, res2, avg1, avg2, sum uint64
	var positions []uint64
	// Loading crab positions
	for i := 0; i < len(crabs); i++ {
		var cur uint64
		for cur = 0; i < len(crabs) && crabs[i] != ','; i++ {
			cur = cur*10 + uint64(crabs[i]-'0')
		}
		positions = append(positions, cur)
		// Summing positions to later find average
		sum += cur
	}
	// Finding average with floor and ceiling
	avg1 = sum / uint64(len(positions))
	avg2 = sum/uint64(len(positions)) + 1
	// Calculating distance fromed floored average and ceiling average
	for i := range positions {
		res1 += uint64(calcFuel(int(positions[i]), int(avg1)))
		res2 += uint64(calcFuel(int(positions[i]), int(avg2)))
	}
	// Picking the smaller one
	return strconv.FormatUint(min(res1, res2), 10), nil
}

// absDist calculates absolute distance
func absDist(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}

// calcFuel calculates how much fuel to cover a trip
func calcFuel(src, dst int) int {
	delta := absDist(src, dst)
	// sum of numbers from 1 to delta
	return (delta * (delta + 1)) / 2
}

func min(a, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}
