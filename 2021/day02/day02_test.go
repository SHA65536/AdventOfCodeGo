package day02

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := SubmarinePos(input)
	if err != nil {
		t.FailNow()
	}
	if res != "2120749" {
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubmarinePos(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := SubmarinePosAim(input)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(res)
	if res != "2120749" {
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubmarinePosAim(input)
	}
}
