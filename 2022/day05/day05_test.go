package day05

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := TopCrates(input)
	if err != nil || res != "BWNCQRMDB" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TopCrates(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := TopCratesMulti(input)
	if err != nil || res != "NHWZCBNBF" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TopCratesMulti(input)
	}
}
