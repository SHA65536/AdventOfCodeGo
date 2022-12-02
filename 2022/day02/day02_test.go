package day02

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := RockPaperScissors(input)
	if err != nil {
		t.FailNow()
	}
	if res != "13221" {
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RockPaperScissors(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := RockPaperScissorsModified(input)
	if err != nil {
		t.FailNow()
	}
	if res != "13131" {
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RockPaperScissorsModified(input)
	}
}
