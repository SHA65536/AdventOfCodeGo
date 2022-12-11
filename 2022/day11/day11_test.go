package day11

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := MonkeyBusiness(input, 20, 3)
	if err != nil || res != "57838" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MonkeyBusiness(input, 20, 3)
	}
}

func TestStar2(t *testing.T) {
	res, err := MonkeyBusiness(input, 10000, 1)
	if err != nil || res != "15050382231" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MonkeyBusiness(input, 10000, 1)
	}
}
