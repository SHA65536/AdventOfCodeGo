package day09

import "strconv"

// https://adventofcode.com/2021/day/9

func RiskLevels(heights string) (string, error) {
	var res int
	var floor [][]int
	// Loading the floor
	for i := 0; i < len(heights); i++ {
		var cur []int
		for ; i < len(heights) && heights[i] != '\n'; i++ {
			cur = append(cur, int(heights[i]-'0'))
		}
		floor = append(floor, cur)
	}
	// Finding low points
	for i := range floor {
		for j := range floor[i] {
			if isLowest(floor, i, j) {
				res += floor[i][j] + 1
			}
		}
	}
	return strconv.Itoa(res), nil
}

func BasinMap(heights string) (string, error) {
	var floor [][]int
	var flows = map[[2]int][2]int{}
	var basins = map[[2]int]int{}
	var maxBasins [3]int
	// Loading the floor
	for i := 0; i < len(heights); i++ {
		var cur []int
		for ; i < len(heights) && heights[i] != '\n'; i++ {
			cur = append(cur, int(heights[i]-'0'))
		}
		floor = append(floor, cur)
	}
	// Calculating where each height is flowing to
	for i := range floor {
		for j := range floor[i] {
			if floor[i][j] != 9 {
				calcFlow(floor, i, j, flows)
			}
		}
	}
	// Calculating size of basins
	for _, v := range flows {
		basins[v]++
	}
	// Finding max 3 basins
	for _, v := range basins {
		UpdateMaxThree(&maxBasins, v)
	}
	return strconv.Itoa(maxBasins[0] * maxBasins[1] * maxBasins[2]), nil
}

// isLowest checks if coordinates have any smaller neighbors
func isLowest(floor [][]int, i, j int) bool {
	if i > 0 && floor[i][j] >= floor[i-1][j] {
		return false
	}
	if j > 0 && floor[i][j] >= floor[i][j-1] {
		return false
	}
	if i < len(floor)-1 && floor[i][j] >= floor[i+1][j] {
		return false
	}
	if j < len(floor[0])-1 && floor[i][j] >= floor[i][j+1] {
		return false
	}
	return true
}

// calcFlows updates given coordinate flow endpoint recursively
func calcFlow(floor [][]int, i, j int, flow map[[2]int][2]int) {
	// Checking if outside
	if i < 0 || i > len(floor) || j < 0 || j > len(floor[0]) {
		return
	}
	// Checking if already calculated
	if _, ok := flow[[2]int{i, j}]; ok {
		return
	}
	// Checking if lowest
	if isLowest(floor, i, j) {
		// If lowest, endpoint is itself
		flow[[2]int{i, j}] = [2]int{i, j}
		return
	}
	if i > 0 && floor[i][j] > floor[i-1][j] {
		calcFlow(floor, i-1, j, flow)
		flow[[2]int{i, j}] = flow[[2]int{i - 1, j}]
	} else if j > 0 && floor[i][j] > floor[i][j-1] {
		calcFlow(floor, i, j-1, flow)
		flow[[2]int{i, j}] = flow[[2]int{i, j - 1}]
	} else if i < len(floor)-1 && floor[i][j] > floor[i+1][j] {
		calcFlow(floor, i+1, j, flow)
		flow[[2]int{i, j}] = flow[[2]int{i + 1, j}]
	} else if j < len(floor[0])-1 && floor[i][j] > floor[i][j+1] {
		calcFlow(floor, i, j+1, flow)
		flow[[2]int{i, j}] = flow[[2]int{i, j + 1}]
	}
}

func UpdateMaxThree(arr *[3]int, val int) {
	// Bubbling the value down
	for i := len(arr) - 1; i >= 0; i-- {
		if val > arr[i] {
			val, arr[i] = arr[i], val
		}
	}
}
