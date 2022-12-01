package day01

import "strconv"

// https://adventofcode.com/2021/day/1

func HeightIncrease(heights string) (string, error) {
	var res int
	var cur, prev int
	// Looping over characters
	for i := range heights {
		if heights[i] >= '0' && heights[i] <= '9' {
			// If it's a digit, add to cur
			cur = cur*10 + int(heights[i]-'0')
		} else {
			// If it's not a digit our number ends
			if cur > prev {
				res++
			}
			prev = cur
			cur = 0
		}
	}
	// Last height is outside the loop
	if cur > prev {
		res++
	}
	// Minus 1 to compensate for the first height
	return strconv.Itoa(res - 1), nil
}

func HeightIncreaseThree(heights string) (string, error) {
	var res, cur int
	var prev [3]int
	// Looping over characters
	for i := range heights {
		if heights[i] >= '0' && heights[i] <= '9' {
			// If it's a digit, add to cur
			cur = cur*10 + int(heights[i]-'0')
		} else {
			// If it's not a digit our number ends
			// Comparing last 3 heights to last 2 heights + current height
			if prev[0]+prev[1]+prev[2] < prev[1]+prev[2]+cur {
				res++
			}
			// Moving the heights back
			prev[0], prev[1], prev[2] = prev[1], prev[2], cur
			cur = 0
		}
	}
	// Last height is outside the loop
	if prev[0]+prev[1]+prev[2] < prev[1]+prev[2]+cur {
		res++
	}
	// Minus 1 to compensate for the first three heights
	return strconv.Itoa(res - 3), nil
}
