package day04

import "strconv"

// https://adventofcode.com/2022/day/4

func CleaningContained(regions string) (string, error) {
	var res int
	for i := 0; i < len(regions); i++ {
		var s1, e1, s2, e2 int
		// Loading the numbers
		for s1 = 0; regions[i] != '-'; i++ {
			s1 = s1*10 + int(regions[i]-'0')
		}
		i++
		for e1 = 0; regions[i] != ','; i++ {
			e1 = e1*10 + int(regions[i]-'0')
		}
		i++
		for s2 = 0; regions[i] != '-'; i++ {
			s2 = s2*10 + int(regions[i]-'0')
		}
		i++
		for e2 = 0; i < len(regions) && regions[i] != '\n'; i++ {
			e2 = e2*10 + int(regions[i]-'0')
		}
		// Checking if all is contained
		if (s1 <= s2) == (e1 >= e2) {
			res++
		}
	}
	return strconv.Itoa(res), nil
}

func CleaningIntersect(regions string) (string, error) {
	var res int
	for i := 0; i < len(regions); i++ {
		var s1, e1, s2, e2 int
		// Loading the numbers
		for s1 = 0; regions[i] != '-'; i++ {
			s1 = s1*10 + int(regions[i]-'0')
		}
		i++
		for e1 = 0; regions[i] != ','; i++ {
			e1 = e1*10 + int(regions[i]-'0')
		}
		i++
		for s2 = 0; regions[i] != '-'; i++ {
			s2 = s2*10 + int(regions[i]-'0')
		}
		i++
		for e2 = 0; i < len(regions) && regions[i] != '\n'; i++ {
			e2 = e2*10 + int(regions[i]-'0')
		}
		// Checking intersection
		if e2 >= s1 && s2 <= e1 {
			res++
		}
	}
	return strconv.Itoa(res), nil
}
