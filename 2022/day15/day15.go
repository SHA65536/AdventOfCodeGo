package day15

import (
	"math/big"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/15

func AreaSearched(input string, yCheck int) (string, error) {
	var res int

	var sensors = LoadSensors(input)
	var board = map[int]bool{}

	// For each sensor
	for _, s := range sensors {
		for i := s.X - s.Dist - 2; i < s.X+s.Dist+2; i++ {
			// Check all intersection with yCheck
			if CabDist(i, yCheck, s.X, s.Y) <= s.Dist {
				board[i] = true
			}
		}
	}

	// Subtract beacons that intersect
	for _, s := range sensors {
		if s.BY == yCheck {
			board[s.BX] = false
		}
	}

	// Sum up intersection with line
	for _, v := range board {
		if v {
			res++
		}
	}

	return strconv.Itoa(res), nil
}

func FindBeacon(input string) (string, error) {
	var sensors = LoadSensors(input)

	// Magic
	for _, s1 := range sensors {
		for _, s2 := range sensors {
			for _, m1 := range lim(s1.X, s1.Y, s1.BX, s1.BY, 1) {
				for _, o2 := range off(s2.X, s2.Y, s2.BX, s2.BY, 1) {
					p1, p2 := intersect(m1, o2)
					if answer(p1, p2, sensors) {
						return getAnswer(p1, p2), nil
					}
				}
			}
			for _, o1 := range off(s1.X, s1.Y, s1.BX, s1.BY, 1) {
				for _, m2 := range lim(s2.X, s2.Y, s2.BX, s2.BY, 1) {
					p1, p2 := intersect(o1, m2)
					if answer(p1, p2, sensors) {
						return getAnswer(p1, p2), nil
					}
				}
			}
		}
	}

	return "", nil
}

// lim is magic
func lim(sx, sy, bx, by, offset int) []int {
	r := CabDist(sx, sy, bx, by) + offset
	return []int{sx + sy + r, sx + sy - r}
}

// off is magic
func off(sx, sy, bx, by, offset int) []int {
	r := CabDist(sx, sy, bx, by) + offset
	return []int{sx - sy + r, sx - sy - r}
}

// intersect is magic
func intersect(sum, dif int) (int, int) {
	x := (sum + dif) / 2
	y := sum - x
	return x, y
}

// answer checks if we are within the requirements
func answer(x, y int, sensors []*Sensor) bool {
	if x < 0 || x > 4000000 || y < 0 || y > 4000000 {
		return false
	}
	for _, s := range sensors {
		if CabDist(s.X, s.Y, x, y) <= CabDist(s.X, s.Y, s.BX, s.BY) {
			return false
		}
	}
	return true
}

// getAnswer finds the tuning frequency
func getAnswer(x, y int) string {
	a := big.NewInt(int64(x))
	a.Mul(a, big.NewInt(4000000))
	a.Add(a, big.NewInt(int64(y)))
	return a.String()
}

// Loads sensors
func LoadSensors(input string) []*Sensor {
	var res []*Sensor
	for _, line := range strings.Split(input, "\n") {
		var sX, sY, bX, bY, i int
		var mod int = 1
		for i = 12; line[i] != ','; i++ {
			if line[i] == '-' {
				mod = -1
			} else {
				sX = sX*10 + int(line[i]-'0')
			}
		}
		sX *= mod
		i += 4
		mod = 1
		for ; line[i] != ':'; i++ {
			if line[i] == '-' {
				mod = -1
			} else {
				sY = sY*10 + int(line[i]-'0')
			}
		}
		sY *= mod
		i += 25
		mod = 1
		for ; line[i] != ','; i++ {
			if line[i] == '-' {
				mod = -1
			} else {
				bX = bX*10 + int(line[i]-'0')
			}
		}
		bX *= mod
		i += 4
		mod = 1
		for ; i < len(line); i++ {
			if line[i] == '-' {
				mod = -1
			} else {
				bY = bY*10 + int(line[i]-'0')
			}
		}
		bY *= mod
		res = append(res, &Sensor{sX, sY, bX, bY, CabDist(sX, sY, bX, bY)})
	}
	return res
}

type Sensor struct {
	X, Y   int
	BX, BY int
	Dist   int
}

// Manhattan distance
func CabDist(ax, ay, bx, by int) int {
	return abs(ax-bx) + abs(ay-by)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
