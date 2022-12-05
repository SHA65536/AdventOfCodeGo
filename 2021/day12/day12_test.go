package day12

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := UniquePaths(input)
	if err != nil || res != "5252" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniquePaths(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := UniquePathsTwo(input)
	if err != nil || res != "147784" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniquePathsTwo(input)
	}
}
