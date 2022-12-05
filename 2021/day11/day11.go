package day11

import "strconv"

// https://adventofcode.com/2021/day/11

var directions = [8][2]int{
	{1, 1}, {1, 0}, {1, -1},
	{0, 1}, {0, -1},
	{-1, 1}, {-1, 0}, {-1, -1},
}

func NumberOfFlashes(octo string, steps int) (string, error) {
	var res int
	var energy [10][10]byte
	var cur = 0
	var cont bool
	// Loading energy map
	for i := range energy {
		for j := range energy[i] {
			energy[i][j] = octo[cur] - '0'
			cur++
		}
		cur++
	}
	// Calculating flashes
	for s := 0; s < steps; s++ {
		// Adding 1 to all of them
		for i := range energy {
			for j := range energy[i] {
				energy[i][j]++
			}
		}
		cont = true
		// Loop over array until no one needs to flash
		for cont {
			cont = false
			for i := range energy {
				for j := range energy[i] {
					// If needs to flash
					if energy[i][j] > 9 {
						cont = true
						res++
						energy[i][j] = 0
						// Loop over all directions and add 1
						for _, dir := range directions {
							ni, nj := i+dir[0], j+dir[1]
							if !isOutside(ni, nj) && energy[ni][nj] != 0 {
								energy[ni][nj]++
							}
						}
					}
				}
			}
		}
	}
	return strconv.Itoa(res), nil
}

func AllFlash(octo string) (string, error) {
	var res, flashed, cur int
	var energy [10][10]byte
	var cont bool
	// Loading energy map
	for i := range energy {
		for j := range energy[i] {
			energy[i][j] = octo[cur] - '0'
			cur++
		}
		cur++
	}
	// Calculating flashes until exactly 100 flashed in a single turn
	for res = 0; flashed != 100; res++ {
		// Adding 1 to all of them
		flashed = 0
		for i := range energy {
			for j := range energy[i] {
				energy[i][j]++
			}
		}
		cont = true
		// Loop over array until no one needs to flash
		for cont {
			cont = false
			for i := range energy {
				for j := range energy[i] {
					// If needs to flash
					if energy[i][j] > 9 {
						cont = true
						flashed++
						energy[i][j] = 0
						// Loop over all directions and add 1
						for _, dir := range directions {
							ni, nj := i+dir[0], j+dir[1]
							if !isOutside(ni, nj) && energy[ni][nj] != 0 {
								energy[ni][nj]++
							}
						}
					}
				}
			}
		}
	}
	return strconv.Itoa(res), nil
}

func isOutside(i, j int) bool {
	return i < 0 || i > 9 || j < 0 || j > 9
}
