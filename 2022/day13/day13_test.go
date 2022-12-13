package day13

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := InRightOrder(input)
	if err != nil || res != "5580" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InRightOrder(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := SortPackets(input)
	if err != nil || res != "26200" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortPackets(input)
	}
}
