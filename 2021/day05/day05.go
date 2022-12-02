package day05

import "strconv"

// https://adventofcode.com/2021/day/5

func VentOverlap(vents string) (string, error) {
	var ventMap = map[[2]int]int{}
	var overlapping int
	// Looping over vents
	for i := 0; i < len(vents); i++ {
		var x1, y1, x2, y2 int
		// Getting x1 number
		for x1 = 0; vents[i] >= '0' && vents[i] <= '9'; i++ {
			x1 = x1*10 + int(vents[i]-'0')
		}
		i++
		// Getting y1 number
		for y1 = 0; vents[i] >= '0' && vents[i] <= '9'; i++ {
			y1 = y1*10 + int(vents[i]-'0')
		}
		i += 4 // Skipping the arrow
		// Getting x2 number
		for x2 = 0; vents[i] >= '0' && vents[i] <= '9'; i++ {
			x2 = x2*10 + int(vents[i]-'0')
		}
		i++
		// Getting y2 number
		for y2 = 0; i < len(vents) && vents[i] >= '0' && vents[i] <= '9'; i++ {
			y2 = y2*10 + int(vents[i]-'0')
		}
		// If line is not diagonal or vertical, skip
		if x1 != x2 && y1 != y2 {
			continue
		}
		// Get affected vents
		for _, p := range GetVents(x1, y1, x2, y2) {
			ventMap[p]++
		}
	}
	for _, v := range ventMap {
		// Count spots with more than 2 vents
		if v >= 2 {
			overlapping++
		}
	}
	return strconv.Itoa(overlapping), nil
}

func VentOverlapDiag(vents string) (string, error) {
	var ventMap = map[[2]int]int{}
	var overlapping int
	// Looping over vents
	for i := 0; i < len(vents); i++ {
		var x1, y1, x2, y2 int
		// Getting x1 number
		for x1 = 0; vents[i] >= '0' && vents[i] <= '9'; i++ {
			x1 = x1*10 + int(vents[i]-'0')
		}
		i++
		// Getting y1 number
		for y1 = 0; vents[i] >= '0' && vents[i] <= '9'; i++ {
			y1 = y1*10 + int(vents[i]-'0')
		}
		i += 4 // Skipping the arrow
		// Getting x2 number
		for x2 = 0; vents[i] >= '0' && vents[i] <= '9'; i++ {
			x2 = x2*10 + int(vents[i]-'0')
		}
		i++
		// Getting y2 number
		for y2 = 0; i < len(vents) && vents[i] >= '0' && vents[i] <= '9'; i++ {
			y2 = y2*10 + int(vents[i]-'0')
		}
		// Getting all vents
		for _, p := range GetVents(x1, y1, x2, y2) {
			ventMap[p]++
		}
	}
	for _, v := range ventMap {
		// Count spots with more than 2 vents
		if v >= 2 {
			overlapping++
		}
	}
	return strconv.Itoa(overlapping), nil
}

func GetVents(x1, y1, x2, y2 int) [][2]int {
	var res [][2]int
	var stepX, stepY int
	// Since lines can only be vertical, horizontal, or at 45 degrees
	// each step can be either 0, -1 or 1
	if x1 > x2 {
		// X downwards
		stepX = -1
	} else if x1 < x2 {
		// X upwards
		stepX = 1
	}
	if y1 > y2 {
		// Y to the left
		stepY = -1
	} else if y1 < y2 {
		// Y to the right
		stepY = 1
	}
	// Moving until reaching the end of the line
	for x1 != x2 || y1 != y2 {
		// Adding vent and applying the step
		res = append(res, [2]int{x1, y1})
		x1 += stepX
		y1 += stepY
	}
	// Last vent should be counted too
	res = append(res, [2]int{x1, y1})
	return res
}
