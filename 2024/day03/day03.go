package day03

import (
	"adventofcode/helper"
	"regexp"
	"strconv"
	"strings"
)

var pat = regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var text, _ = input.ReadAll()

	for _, match := range pat.FindAllStringSubmatch(text, -1) {
		res += helper.MustConvNum(match[1]) * helper.MustConvNum(match[2])
	}

	return strconv.Itoa(res), nil
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var text, _ = input.ReadAll()

	var parts = strings.Split(text, "do()")

	for _, part := range parts {
		if strings.Contains(part, "don't()") {
			part = part[:strings.Index(part, "don't()")]
		}
		res += CalcInstructions(part)
	}

	return strconv.Itoa(res), nil
}

func CalcInstructions(in string) int {
	var res int
	for _, match := range pat.FindAllStringSubmatch(in, -1) {
		res += helper.MustConvNum(match[1]) * helper.MustConvNum(match[2])
	}
	return res
}
