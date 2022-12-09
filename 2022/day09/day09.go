package day09

import (
	"math"
	"strconv"
)

// https://adventofcode.com/2022/day/9

var dirs = map[byte][2]int{
	'R': {0, 1}, 'L': {0, -1}, // Directions
	'U': {-1, 0}, 'D': {1, 0},
	's': {0, 0}, 'z': {1, 1}, // Diagonals  and center for around check
	'c': {1, -1}, 'q': {-1, 1},
	'e': {-1, -1},
}

func TailPos(input string) (string, error) {
	var val int
	var head, tail, dir [2]int
	var seen = map[[2]int]bool{tail: true}

	for i := 0; i < len(input); i++ {
		// Loading direction
		dir = dirs[input[i]]
		i += 2
		// Loading value
		for val = 0; i < len(input) && input[i] != '\n'; i++ {
			val = val*10 + int(input[i]-'0')
		}
		// Moving value
		for j := 0; j < val; j++ {
			head[0], head[1] = head[0]+dir[0], head[1]+dir[1]
			if !isAround(tail, head) {
				tail[0] = head[0] - dir[0]
				tail[1] = head[1] - dir[1]
			}
			seen[tail] = true
		}
	}

	return strconv.Itoa(len(seen)), nil
}

func NineTailPos(input string) (string, error) {
	var val int
	var head, dir [2]int
	var tails [][2]int = make([][2]int, 9)
	var seen = map[[2]int]bool{tails[8]: true}

	for i := 0; i < len(input); i++ {
		// Loading direction
		dir = dirs[input[i]]
		i += 2
		// Loading value
		for val = 0; i < len(input) && input[i] != '\n'; i++ {
			val = val*10 + int(input[i]-'0')
		}
		// Moving value
		for j := 0; j < val; j++ {
			// Moving head according to dir
			head[0], head[1] = head[0]+dir[0], head[1]+dir[1]
			var prev = head
			for i := range tails {
				// If knot is not around the previous knot
				if !isAround(tails[i], prev) {
					// Check which move will make it closest
					tails[i] = closestSide(tails[i], prev)
				}
				prev = tails[i]
			}
			// Recording the last tail move
			seen[tails[len(tails)-1]] = true
		}
	}

	return strconv.Itoa(len(seen)), nil
}

func isAround(a, b [2]int) bool {
	for _, dir := range dirs {
		n := [2]int{a[0] + dir[0], a[1] + dir[1]}
		if n == b {
			return true
		}
	}
	return false
}

func closestSide(a, b [2]int) [2]int {
	var mDist float64 = 1000
	var curTypes []byte
	var mPos [2]int
	if a[0] == b[0] || a[1] == b[1] {
		// If in the same column or row, move straight
		curTypes = []byte("RDUL") // Straight moves
	} else {
		// If not in the same column or row, move diagonal
		curTypes = []byte("zcqe") // Diagonal moves
	}
	for _, d := range curTypes {
		// New position
		n := [2]int{a[0] + dirs[d][0], a[1] + dirs[d][1]}
		// Distance of new position to previous
		dist := math.Sqrt(abs(n[0]-b[0])*abs(n[0]-b[0]) + abs(n[1]-b[1])*abs(n[1]-b[1]))
		// Picking minimum distance
		if dist < mDist {
			mDist = dist
			mPos = n
		}
	}
	return mPos
}

func abs(a int) float64 {
	if a < 0 {
		return float64(-a)
	}
	return float64(a)
}
