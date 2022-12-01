package day03

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := GammaEpsilon(input)
	if err != nil {
		t.FailNow()
	}
	if res != "4191876" {
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GammaEpsilon(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := Oxygen(input)
	if err != nil {
		t.FailNow()
	}
	if res != "3414905" {
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Oxygen(input)
	}
}
