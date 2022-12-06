package day06

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := StartOfPacket(input, 4)
	if err != nil || res != "1093" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StartOfPacket(input, 4)
	}
}

func TestStar2(t *testing.T) {
	res, err := StartOfPacket(input, 14)
	if err != nil || res != "1093" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StartOfPacket(input, 14)
	}
}
