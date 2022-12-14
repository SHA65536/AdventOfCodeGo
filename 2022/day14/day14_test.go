package day14

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := HowMuchSand(input)
	if err != nil || res != "1298" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HowMuchSand(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := SandWithFloor(input)
	if err != nil || res != "25585" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SandWithFloor(input)
	}
}
