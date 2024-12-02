package day24

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
)

const tMin, tMax = 200000000000000, 400000000000000

type Hail struct {
	PX, PY, PZ, VX, VY, VZ float64
}

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var stones = make([]Hail, 0, 300)

	for line := range input.IterateLines {
		var h Hail
		fmt.Sscanf(line, "%f, %f, %f @ %f, %f, %f", &h.PX, &h.PY, &h.PZ, &h.VX, &h.VY, &h.VZ)
		stones = append(stones, h)
	}

	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if WillIntersect(stones[i], stones[j]) {
				res++
			}
		}
	}

	return strconv.Itoa(res), nil
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	return strconv.Itoa(res), nil
}

func WillIntersect(h1, h2 Hail) bool {
	if h1.VX*h2.VY == h1.VY*h2.VX {
		return false
	}
	var t1, t2 float64

	t2 = ((h1.PX-h2.PX)*h1.VY - (h1.PY-h2.PY)*h1.VX) / (h2.VX*h1.VY - h2.VY*h1.VX)
	if h1.VX != 0 {
		t1 = ((h2.PX - h1.PX) + t2*h2.VX) / h1.VX
	} else {
		t1 = ((h2.PY - h1.PY) + t2*h2.VY) / h1.VY
	}
	var ix = h1.PX + t1*h1.VX
	var iy = h1.PY + t1*h1.VY

	if (h1.VX > 0 && ix <= h1.PX) || (h2.VX > 0 && ix <= h2.PX) {
		return false
	}

	if (h1.VX < 0 && ix >= h1.PX) || (h2.VX < 0 && ix >= h2.PX) {
		return false
	}

	return ix >= tMin && ix <= tMax && iy >= tMin && iy <= tMax
}
