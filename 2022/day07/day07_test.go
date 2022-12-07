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
	if err != nil || res != "95437" {
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
	if err != nil || res != "24933642" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindDelete(input)
	}
}
