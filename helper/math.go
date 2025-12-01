package helper

import "golang.org/x/exp/constraints"

func AbsoluteDistance[T constraints.Integer](a, b T) T {
	if a > b {
		return a - b
	}
	return b - a
}

func Mod[T constraints.Integer](dividend, divisor T) T {
	if divisor < 0 {
		divisor = -divisor
	}
	res := dividend % divisor
	if res < 0 {
		res += divisor
	}
	return res
}
