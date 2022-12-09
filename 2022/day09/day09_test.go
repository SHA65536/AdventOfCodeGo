package day09

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := TailPos(input)
	if err != nil || res != "6087" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TailPos(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := NineTailPos(input)
	if err != nil || res != "2493" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NineTailPos(input)
	}
}
