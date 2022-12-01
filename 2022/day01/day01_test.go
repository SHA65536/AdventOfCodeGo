package day01

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var calories string

func TestStar1(t *testing.T) {
	res, err := MaxCalories(calories)
	if err != nil {
		t.FailNow()
	}
	if res != "71924" {
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxCalories(calories)
	}
}

func TestStar2(t *testing.T) {
	res, err := MaxCaloriesThree(calories)
	if err != nil {
		t.FailNow()
	}
	if res != "210406" {
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxCaloriesThree(calories)
	}
}
