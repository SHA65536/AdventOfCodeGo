package day08

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := NumDigits(input)
	if err != nil || res != "476" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumDigits(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := AllDigits(input)
	if err != nil || res != "1011823" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllDigits(input)
	}
}
