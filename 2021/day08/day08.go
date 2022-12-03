package day08

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/8

func NumDigits(signals string) (string, error) {
	var res int
	lines := strings.Split(signals, "\n")
	for _, line := range lines {
		digits := strings.Split(strings.Split(line, " | ")[1], " ")
		for _, digit := range digits {
			// Summing the digits with known number of signals
			switch len(digit) {
			case 2, 3, 4, 7:
				res++
			}
		}
	}
	return strconv.Itoa(res), nil
}

func AllDigits(signals string) (string, error) {
	var res uint64
	lines := strings.Split(signals, "\n")
	for _, line := range lines {
		var curRes uint64
		// Finding 1, 4, 7 and 8 accoring to the number of segments
		var decoding, encoding, digits = FindWithLen(line[:58])
		for len(encoding) != 10 {
			for i := range digits {
				// Finding the rest by transitivity
				if FindWithMaps(digits[i], decoding, encoding) {
					digits = Remove(digits, i)
					break
				}
			}
		}
		// Summing up the decoded digits
		for _, digit := range strings.Split(line[61:], " ") {
			var cur Digit
			for i := range digit {
				cur[digit[i]-'a'] = true
			}
			curRes = curRes*10 + uint64(decoding[cur])
		}
		res += curRes
	}
	return strconv.FormatUint(res, 10), nil
}

func FindWithLen(input string) (map[Digit]int, map[int]Digit, []Digit) {
	var decoding = map[Digit]int{}
	var encoding = map[int]Digit{}
	var digits []Digit
	for _, digit := range strings.Split(input, " ") {
		var cur Digit
		for i := range digit {
			cur[digit[i]-'a'] = true
		}
		switch cur.Sum() {
		case 2: // Only 1 has 2 segments
			decoding[cur] = 1
			encoding[1] = cur
		case 3: // Only 7 has 3 segments
			decoding[cur] = 7
			encoding[7] = cur
		case 4: // Only 4 has 4 segments
			decoding[cur] = 4
			encoding[4] = cur
		case 7: // Only 8 has 7 segments
			decoding[cur] = 8
			encoding[8] = cur
		default: // Save the rest for further proccessing
			digits = append(digits, cur)
		}
	}
	return decoding, encoding, digits
}

func FindWithMaps(d Digit, dec map[Digit]int, enc map[int]Digit) bool {
	if d.Sum() == 6 { // Only 9, 0 and 6 have 6 segments
		if d.InCommon(enc[4]) == 4 { // Only 9 has 4 in common with 4
			dec[d] = 9
			enc[9] = d
			return true
		} else if d.InCommon(enc[1]) == 2 { // Only 0 has 2 in common with 1
			dec[d] = 0
			enc[0] = d
			return true
		} else { // Only 6 is remaining with 6 segments
			dec[d] = 6
			enc[6] = d
			return true
		}
	}
	if d.Sum() == 5 { // Only 3, 5 and 2 have 5 segments
		if d.InCommon(enc[1]) == 2 { // Only 3 has 2 in common with 1
			dec[d] = 3
			enc[3] = d
			return true
		} else if d.InCommon(enc[4]) == 3 { // Only 5 has 3 in common with 4
			dec[d] = 5
			enc[5] = d
			return true
		} else { // Only 2 is remaining with 5 segments
			dec[d] = 2
			enc[2] = d
			return true
		}
	}
	return false
}

type Digit [7]bool

func (d *Digit) Sum() int {
	var sum int
	for i := range d {
		if d[i] {
			sum++
		}
	}
	return sum
}

func (d *Digit) InCommon(in Digit) int {
	var sum int
	for i := range d {
		if d[i] && in[i] {
			sum++
		}
	}
	return sum
}

func Remove(in []Digit, idx int) []Digit {
	in[idx] = in[len(in)-1]
	in = in[:len(in)-1]
	return in
}
