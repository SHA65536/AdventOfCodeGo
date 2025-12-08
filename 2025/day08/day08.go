package day08

import (
	"adventofcode/helper"
	"fmt"
	"math"
	"sort"
	"strconv"
)

type Dist struct {
	D      float64
	B1, B2 Box
}

type Box struct {
	X, Y, Z int
}

func LineDistance(a, b Box) float64 {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	dz := float64(a.Z - b.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func Star1(input *helper.InputReader, conns int) (string, error) {
	var res int = 1

	var boxes []Box
	for line := range input.IterateLines {
		var box Box
		fmt.Sscanf(line, "%d,%d,%d", &(box.X), &(box.Y), &(box.Z))
		boxes = append(boxes, box)
	}

	var distances []Dist
	var circuits = make(map[Box]Box, len(boxes))
	for i := 0; i < len(boxes); i++ {
		circuits[boxes[i]] = boxes[i]
		for j := i + 1; j < len(boxes); j++ {
			distances = append(distances, Dist{B1: boxes[i], B2: boxes[j], D: LineDistance(boxes[i], boxes[j])})
		}
	}

	sort.Slice(distances, func(i, j int) bool { return distances[i].D < distances[j].D })

	for i := range conns {
		union(circuits, distances[i].B1, distances[i].B2)
	}

	var rootSizes = map[Box]int{}
	for _, b := range boxes {
		rootSizes[find(circuits, b)]++
	}

	var a, b, c int
	for _, size := range rootSizes {
		if size > a {
			a, b, c = size, a, b
		} else if size > b {
			b, c = size, b
		} else if size > c {
			c = size
		}
	}

	res = a * b * c

	return strconv.Itoa(res), nil
}

func find[T comparable](uf map[T]T, x T) T {
	if uf[x] == x {
		return uf[x]
	} else {
		uf[x] = find(uf, uf[x])
		return uf[x]
	}
}

func union[T comparable](uf map[T]T, x, y T) {
	var rootx, rooty T
	if _, found := uf[x]; !found {
		uf[x] = x
	}
	if _, found := uf[y]; !found {
		uf[y] = y
	}
	rootx = find(uf, x)
	rooty = find(uf, y)
	uf[rootx] = rooty
}

func Star2(input *helper.InputReader) (string, error) {
	var res int = 1

	var boxes []Box
	for line := range input.IterateLines {
		var box Box
		fmt.Sscanf(line, "%d,%d,%d", &(box.X), &(box.Y), &(box.Z))
		boxes = append(boxes, box)
	}

	var distances []Dist
	var circuits = make(map[Box]Box, len(boxes))
	for i := 0; i < len(boxes); i++ {
		circuits[boxes[i]] = boxes[i]
		for j := i + 1; j < len(boxes); j++ {
			distances = append(distances, Dist{B1: boxes[i], B2: boxes[j], D: LineDistance(boxes[i], boxes[j])})
		}
	}

	sort.Slice(distances, func(i, j int) bool { return distances[i].D < distances[j].D })

	for i := range distances {
		union(circuits, distances[i].B1, distances[i].B2)
		if checkAllConnected(circuits, boxes) {
			res = distances[i].B1.X * distances[i].B2.X
			break
		}
	}

	return strconv.Itoa(res), nil
}

func checkAllConnected[T comparable](uf map[T]T, items []T) bool {
	var src = find(uf, items[0])
	for i := range items {
		if find(uf, items[i]) != src {
			return false
		}
	}
	return true
}
