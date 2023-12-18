package day01

import (
	"fmt"
	"strconv"
	"strings"
)

// https://adventofcode.com/2023/day/1

func TrebuchetCalibration(input string) (string, error) {
	var res int
	var lines = strings.Split(input, "\n")
	for line := range lines {
		for i := 0; i < len(lines[line]); i++ {
			if isDigit(lines[line][i]) {
				res += int(lines[line][i]-'0') * 10
				break
			}
		}

		for i := len(lines[line]) - 1; i >= 0; i-- {
			if isDigit(lines[line][i]) {
				res += int(lines[line][i] - '0')
				break
			}
		}
	}

	return strconv.Itoa(res), nil
}

var numbers = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
	"six": 6, "seven": 7, "eight": 8, "nine": 9, "zero": 0,
}

var numbers_rev = map[string]int{
	"eno": 1, "owt": 2, "eerht": 3, "ruof": 4, "evif": 5,
	"xis": 6, "neves": 7, "thgie": 8, "enin": 9, "orez": 0,
}

func TrebuchetCalibration2(input string) (string, error) {
	var res int
	var lines = strings.Split(input, "\n")

	for line := range lines {
		res += findNumber(lines[line]) * 10
		res += findNumberRev(lines[line])
	}

	fmt.Println(res)
	return strconv.Itoa(res), nil
}

func findNumber(line string) int {
	for i := range line {
		if isDigit(line[i]) {
			return int(line[i] - '0')
		}
	eachnum:
		for num := range numbers {
			if i+len(num) > len(line) {
				continue
			}
			for j := range num {
				if line[i+j] != num[j] {
					continue eachnum
				}
			}
			return numbers[num]
		}
	}
	return 0
}

func findNumberRev(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if isDigit(line[i]) {
			return int(line[i] - '0')
		}
	eachnum:
		for num := range numbers_rev {
			if i-len(num) < 0 {
				continue
			}
			for j := range num {
				if line[i-j] != num[j] {
					continue eachnum
				}
			}
			return numbers_rev[num]
		}
	}
	return 0
}

func isDigit(b byte) bool { return b >= '0' && b <= '9' }
