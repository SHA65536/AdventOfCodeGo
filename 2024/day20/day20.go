package day20

import (
	"adventofcode/helper"
	"strconv"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	track, s, e := parseTrack(input)

	/*var best = findBestPath(track, s, e)

	for i := range track {
		for j := range track[i] {
			if track[i][j] == '#' {
				track[i][j] = '.'
				cheat := findBestPath(track, s, e)
				if best-cheat >= 100 {
					res++
				}
				track[i][j] = '#'
			}
		}
	}*/

	path := findPath(track, s, e)

	for s := range path {
		for e := s; e < len(path); e++ {
			mDist := abs(path[s].C-path[e].C) + abs(path[s].R-path[e].R)
			if mDist <= 2 && (e-s)-mDist >= 100 {
				res++
			}
		}
	}

	return strconv.Itoa(res), nil
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	track, s, e := parseTrack(input)
	path := findPath(track, s, e)

	for s := range path {
		for e := s; e < len(path); e++ {
			mDist := abs(path[s].C-path[e].C) + abs(path[s].R-path[e].R)
			if mDist <= 20 && (e-s)-mDist >= 100 {
				res++
			}
		}
	}

	return strconv.Itoa(res), nil
}

func abs(a int) int {
	return max(a, -a)
}

type Coord struct {
	R, C int
}

/*func findBestPath(track [][]byte, s, e Coord) int {
	var cur, next []Coord
	cur = []Coord{s}

	var visited = map[Coord]bool{s: true}

	var res int

	for len(cur) > 0 {
		for _, pos := range cur {
			if pos == e {
				return res
			}
			for r, c := range helper.IterateAdjacentC(pos.R, pos.C) {
				if !helper.InsideC(track, r, c) {
					continue
				}
				if !visited[Coord{r, c}] && track[r][c] != '#' {
					visited[Coord{r, c}] = true
					next = append(next, Coord{r, c})
				}
			}
		}
		res++
		cur, next = next, cur[:0]
	}

	return res
}*/

func parseTrack(input *helper.InputReader) (track [][]byte, s, e Coord) {
	var i int
	for line := range input.IterateLines {
		track = append(track, []byte(line))
		for j := range line {
			if line[j] == 'S' {
				s.R = i
				s.C = j
			} else if line[j] == 'E' {
				e.R = i
				e.C = j
			}
		}
		i++
	}

	return track, s, e
}

func findPath(track [][]byte, s, e Coord) []Coord {
	var path = []Coord{}
	var pos int

	for s != e {
		path = append(path, s)
		track[s.R][s.C] = '#'

		for r, c := range helper.IterateAdjacentC(s.R, s.C) {
			if !helper.InsideC(track, r, c) {
				continue
			}

			if track[r][c] != '#' {
				s = Coord{r, c}
				break
			}
		}
		pos++
	}

	path = append(path, e)

	return path
}
