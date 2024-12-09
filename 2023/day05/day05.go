package day05

import (
	"adventofcode/helper"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type ConvRange struct {
	Dst, Src, Size int
}

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var seeds = parseSeeds(input)
	var chain = parseConvChain(input)

	for _, cRange := range chain {
		sort.Slice(cRange, func(i, j int) bool { return cRange[i].Src < cRange[j].Src })
	}

	res = getLocation(seeds[0], chain)
	for _, seed := range seeds[1:] {
		res = min(res, getLocation(seed, chain))
	}

	return strconv.Itoa(res), nil
}

func translate(seed int, cRange []ConvRange) int {
	for _, conv := range cRange {
		if seed >= conv.Src && seed-conv.Src <= conv.Size {
			return (seed - conv.Src) + conv.Dst
		}
	}
	return seed
}

func getLocation(seed int, chain [][]ConvRange) int {
	for _, cRange := range chain {
		seed = translate(seed, cRange)
	}
	return seed
}

func parseSeeds(input *helper.InputReader) []int {
	var _seeds, _ = input.ReadLine()
	var seeds = make([]int, 0)
	for _, word := range strings.Fields(_seeds)[1:] {
		seeds = append(seeds, helper.MustConvNum(word))
	}
	input.ReadLine()
	input.ReadLine()
	return seeds
}

func parseConvChain(input *helper.InputReader) [][]ConvRange {
	var ConvChain = make([][]ConvRange, 0)

	var curChain []ConvRange
	for line := range input.IterateLines {
		if len(line) < 2 {
			continue
		}

		if strings.Contains(line, ":") {
			ConvChain = append(ConvChain, curChain)
			curChain = make([]ConvRange, 0)
			continue
		}

		var conv ConvRange
		fmt.Sscanf(line, "%d %d %d", &conv.Dst, &conv.Src, &conv.Size)
		curChain = append(curChain, conv)
	}

	return append(ConvChain, curChain)
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var seeds = parseSeedsRange(input)
	var chain = parseConvChain(input)

	for _, cRange := range chain {
		sort.Slice(cRange, func(i, j int) bool { return cRange[i].Src < cRange[j].Src })
	}

	var locations [][2]int
	for _, seed := range seeds {
		locations = append(locations, getLocationRange(seed, chain)...)
	}

	res = locations[0][0]
	for _, location := range locations[1:] {
		res = min(res, location[0])
	}

	return strconv.Itoa(res), nil
}

func getLocationRange(seed [2]int, chain [][]ConvRange) [][2]int {
	var seeds = [][2]int{seed}
	var next = [][2]int{}
	for _, cRange := range chain {
		for _, seed := range seeds {
			next = append(next, translateRange(seed, cRange)...)
		}
		seeds, next = next, seeds
		next = next[:0]

	}
	return seeds
}

func translateRange(seed [2]int, cRange []ConvRange) [][2]int {
	var res [][2]int
	for i := 0; i < len(cRange); i++ {
		var diff = cRange[i].Dst - cRange[i].Src
		if seed[0] < cRange[i].Src { // non-translate too small
			if seed[1] < cRange[i].Src { // completely too small
				res = append(res, seed)
				return res // return early since everything is done
			}
			// add the too small part
			res = append(res, [2]int{seed[0], cRange[i].Src - 1})
			// next iteration do the part that is not too small
			seed = [2]int{cRange[i].Src, seed[1]}
			i--
			continue
		}

		// some fit the range
		if seed[0] >= cRange[i].Src && seed[0] <= cRange[i].Src+cRange[i].Size-1 {
			if seed[1] <= cRange[i].Src+cRange[i].Size { // if all in the range
				res = append(res, [2]int{seed[0] + diff, seed[1] + diff})
				return res // return early since everything is done
			}
			// add the part that fits
			res = append(res, [2]int{seed[0] + diff, cRange[i].Dst + cRange[i].Size - 1})
			// next iteration do the part that doesn't fit
			seed = [2]int{cRange[i].Src + cRange[i].Size, seed[1]}
			continue
		}
	}

	// add the rest that didn't fit
	res = append(res, seed)

	return res
}

func parseSeedsRange(input *helper.InputReader) [][2]int {
	var _seedsline, _ = input.ReadLine()
	var _seeds = strings.Fields(_seedsline)[1:]
	var seeds = make([][2]int, 0)
	for i := 0; i < len(_seeds)-1; i += 2 {
		seeds = append(seeds, [2]int{helper.MustConvNum(_seeds[i]), helper.MustConvNum(_seeds[i]) + helper.MustConvNum(_seeds[i+1]) - 1})
	}
	input.ReadLine()
	input.ReadLine()
	return seeds
}
