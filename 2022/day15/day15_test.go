package day15

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := AreaSearched(input, 2000000)
	if err != nil || res != "5508234" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AreaSearched(input, 2000000)
	}
}

func TestStar2(t *testing.T) {
	res, err := FindBeacon(input)
	if err != nil || res != "10457634860779" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindBeacon(input)
	}
}
