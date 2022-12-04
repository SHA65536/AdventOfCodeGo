package day06

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := CommonItem(input)
	if err != nil || res != "7763" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CommonItem(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := CommonThree(input)
	if err != nil || res != "2569" {
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CommonThree(input)
	}
}
