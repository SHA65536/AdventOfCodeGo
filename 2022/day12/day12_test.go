package day12

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := LeastSteps(input)
	if err != nil || res != "447" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeastSteps(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := LeastStepsAny(input)
	if err != nil || res != "446" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeastStepsAny(input)
	}
}
