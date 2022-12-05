package day11

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := NumberOfFlashes(input, 100)
	if err != nil || res != "1601" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberOfFlashes(input, 100)
	}
}

func TestStar2(t *testing.T) {
	res, err := AllFlash(input)
	if err != nil || res != "368" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllFlash(input)
	}
}
