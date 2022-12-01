package day01

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := HeightIncrease(input)
	if err != nil {
		t.FailNow()
	}
	if res != "1233" {
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeightIncrease(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := HeightIncreaseThree(input)
	fmt.Println(res)
	if err != nil {
		t.FailNow()
	}
	if res != "1275" {
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeightIncreaseThree(input)
	}
}
