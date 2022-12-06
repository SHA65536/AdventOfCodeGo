package day15

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := LowestPath(input, 1)
	if err != nil || res != "" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LowestPath(input, 1)
	}
}

func TestStar2(t *testing.T) {
	res, err := LowestPath(input, 5)
	if err != nil || res != "" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LowestPath(input, 5)
	}
}
