package day08

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := VisibleTrees(input)
	if err != nil || res != "1835" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VisibleTrees(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := ScenicScore(input)
	if err != nil || res != "263670" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ScenicScore(input)
	}
}
