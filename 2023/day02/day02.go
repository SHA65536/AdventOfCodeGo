package day02

import (
	"adventofcode/helper"
	"strconv"
	"strings"
)

type colors [26]int

const (
	red   = 'r' - 'a'
	blue  = 'b' - 'a'
	green = 'g' - 'a'
)

var ask colors = colors{red: 12, green: 13, blue: 14}

func CubesGame(input *helper.InputReader) (string, error) {
	var res int

eachline:
	for line := range input.IterateLines {
		words := strings.Split(line, " ")
		for i := 2; i < len(words); i += 2 {
			if ask[words[i+1][0]-'a'] < helper.MustConvNum(words[i]) {
				continue eachline
			}
		}
		res += helper.MustConvNum(words[1][:len(words[1])-1])
	}

	return strconv.Itoa(res), nil
}

func CubesGameFewest(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		words := strings.Split(line, " ")
		var maxi colors
		for i := 2; i < len(words); i += 2 {
			maxi[words[i+1][0]-'a'] = max(maxi[words[i+1][0]-'a'], helper.MustConvNum(words[i]))
		}
		res += maxi[red] * maxi[green] * maxi[blue]
	}

	return strconv.Itoa(res), nil
}
