package day04

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := CleaningContained(input)
	if err != nil || res != "487" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CleaningContained(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := CleaningIntersect(input)
	if err != nil || res != "849" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CleaningIntersect(input)
	}
}
