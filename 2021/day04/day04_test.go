package day04

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := Bingo(input)
	if err != nil {
		t.FailNow()
	}
	if res != "35711" {
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bingo(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := BingoLast(input)
	if err != nil {
		t.FailNow()
	}
	if res != "5586" {
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BingoLast(input)
	}
}
