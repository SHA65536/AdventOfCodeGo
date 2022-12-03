package day07

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, ok := HorizontalAlign(input)
	if ok != nil || res != "342641" {
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HorizontalAlign(input)
	}
}

func TestStar2(t *testing.T) {
	res, ok := HorizontalAlignComplex(input)
	if ok != nil || res != "93006301" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HorizontalAlignComplex(input)
	}
}
