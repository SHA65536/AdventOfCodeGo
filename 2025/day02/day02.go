package day02

import (
	"adventofcode/helper"
	"math"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int
	all, _ := input.ReadAll()
	ranges := strings.Split(all, ",")
	for _, r := range ranges {
		start, end := strings.Split(r, "-")[0], strings.Split(r, "-")[1]
		sInt, eInt := helper.MustConvNum(start), helper.MustConvNum(end)
		firstHalf := helper.MustConvNum(start[:len(start)/2])
		var dub = double(firstHalf)
		for dub < sInt {
			firstHalf++
			dub = double(firstHalf)
		}
		for dub <= eInt {
			res += dub
			firstHalf++
			dub = double(firstHalf)
		}
	}
	return strconv.Itoa(res), nil
}

func double(n int) int {
	digits := int(math.Log10(float64(n))) + 1
	mul := int(math.Pow10(digits))
	return n*mul + n
}

func Star2(input *helper.InputReader) (string, error) {
	var res int
	all, _ := input.ReadAll()
	ranges := strings.Split(all, ",")
	for _, r := range ranges {
		start, end := strings.Split(r, "-")[0], strings.Split(r, "-")[1]
		sInt, eInt := helper.MustConvNum(start), helper.MustConvNum(end)

		for i := sInt; i <= eInt; i++ {
			str := strconv.Itoa(i)
			// Try split it to different sizes and check for repitition
			for size := 1; size < len(str); size++ {
				// Only if it can be split to even sized chunks
				if len(str)%size != 0 {
					continue
				}
				var invalid = true
				for j := 0; j < len(str)/size; j++ {
					if str[j*size:(j+1)*size] != str[0:size] {
						invalid = false
						break
					}
				}
				if invalid {
					res += i
					break // Avoid double counting numbers that can be split multiple ways :)
				}
			}
		}
	}
	return strconv.Itoa(res), nil
}
