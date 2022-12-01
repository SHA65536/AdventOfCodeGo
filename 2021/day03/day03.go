package day03

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/3

func GammaEpsilon(report string) (string, error) {
	var length, countLines int
	var gamma, epsilon int
	var bits []int
	// Finding the length of each line
	for i := range report {
		if report[i] == '\n' {
			length = i
			break
		}
	}
	bits = make([]int, length)
	// Looping over lines
	for i := 0; i < len(report); i++ {
		// Counting the number of ones
		for j := 0; i < len(report) && report[i] != '\n'; j++ {
			if report[i] == '1' {
				bits[j]++
			}
			i++
		}
		countLines++
	}
	// Setting the bits in the number
	for i := range bits {
		gamma <<= 1
		epsilon <<= 1
		if bits[i]*2 > countLines {
			gamma++
		} else {
			epsilon++
		}
	}
	// Epsilon is always not-Gamma
	return strconv.Itoa(gamma * epsilon), nil
}

func Oxygen(report string) (string, error) {
	var lines []string = strings.Split(report, "\n")
	var zeroes, ones []string
	var cur = lines
	// Running until left with one number
	for bit := 0; len(cur) > 1; bit++ {
		zeroes = make([]string, 0, len(cur))
		ones = make([]string, 0, len(cur))
		for i := range cur {
			// Dividing the numbers into numbers with 0 in the current bit
			// and numbers with 1 in the current bit
			if cur[i][bit] == '0' {
				zeroes = append(zeroes, cur[i])
			} else {
				ones = append(ones, cur[i])
			}
		}
		// Picking the most common one
		if len(zeroes) > len(ones) {
			cur = zeroes
		} else {
			cur = ones
		}
		bit++
	}
	// Converting to decimal
	oxygen, err := strconv.ParseInt(cur[0], 2, 64)
	if err != nil {
		return "", err
	}

	cur = lines
	// Running until left with one number
	for bit := 0; len(cur) > 1; bit++ {
		zeroes = make([]string, 0, len(cur))
		ones = make([]string, 0, len(cur))
		for i := range cur {
			// Dividing the numbers into numbers with 0 in the current bit
			// and numbers with 1 in the current bit
			if cur[i][bit] == '0' {
				zeroes = append(zeroes, cur[i])
			} else {
				ones = append(ones, cur[i])
			}
		}
		// Picking the least common one
		if len(zeroes) <= len(ones) {
			cur = zeroes
		} else {
			cur = ones
		}
		bit++
	}
	// Converting to decimal
	co2, err := strconv.ParseInt(cur[0], 2, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(co2*oxygen, 10), nil
}
