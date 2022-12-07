package day16

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := VersionSum(input)
	if err != nil || res != "927" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VersionSum(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := EvalPackets(input)
	if err != nil || res != "1725277876501" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalPackets(input)
	}
}
