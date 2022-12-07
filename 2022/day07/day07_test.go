package day07

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := ParseTerminal(input)
	if err != nil || res != "1141028" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseTerminal(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := FindDelete(input)
	if err != nil || res != "8278005" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindDelete(input)
	}
}
