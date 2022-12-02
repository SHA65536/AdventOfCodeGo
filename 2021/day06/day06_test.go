package day06

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := FishPop(input)
	if err != nil {
		t.FailNow()
	}
	if res != "374994" {
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FishPop(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := FishPopMore(input)
	if err != nil {
		t.FailNow()
	}
	if res != "1686252324092" {
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FishPopMore(input)
	}
}
