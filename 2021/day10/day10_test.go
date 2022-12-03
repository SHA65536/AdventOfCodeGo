package day10

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, ok := CorruptSyntax(input)
	if ok != nil || res != "413733" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CorruptSyntax(input)
	}
}

func TestStar2(t *testing.T) {
	res, ok := IncompleteSyntax(input)
	if ok != nil || res != "3354640192" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IncompleteSyntax(input)
	}
}
