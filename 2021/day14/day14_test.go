package day14

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := PolymerGenerate(input, 10)
	if err != nil || res != "2112" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PolymerGenerate(input, 10)
	}
}

func TestStar2(t *testing.T) {
	res, err := PolymerGenerate(input, 40)
	if err != nil || res != "3243771149914" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PolymerGenerate(input, 40)
	}
}
