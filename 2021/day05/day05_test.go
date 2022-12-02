package day05

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := VentOverlap(input)
	if err != nil {
		t.FailNow()
	}
	if res != "6225" {
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VentOverlap(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := VentOverlapDiag(input)
	if err != nil {
		t.FailNow()
	}
	if res != "22116" {
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VentOverlapDiag(input)
	}
}
