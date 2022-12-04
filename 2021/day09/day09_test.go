package day09

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := RiskLevels(input)
	if err != nil || res != "600" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RiskLevels(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := BasinMap(input)
	if err != nil || res != "987840" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BasinMap(input)
	}
}
