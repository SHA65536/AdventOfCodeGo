package day01

import (
	_ "embed"
	"testing"
)

//go:embed small_input.txt
var smallInput string

//go:embed small_input_2.txt
var smallInput_2 string

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := TrebuchetCalibration(smallInput)
	if err != nil {
		t.FailNow()
	}
	if res != "142" {
		t.FailNow()
	}

	res, err = TrebuchetCalibration(input)
	if err != nil {
		t.FailNow()
	}
	if res != "55538" {
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrebuchetCalibration(input)
	}
}

func TestStar2(t *testing.T) {
	res, err := TrebuchetCalibration2(smallInput_2)
	if err != nil {
		t.FailNow()
	}
	if res != "281" {
		t.FailNow()
	}

	res, err = TrebuchetCalibration2(input)
	if err != nil {
		t.FailNow()
	}
	if res != "54875" {
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrebuchetCalibration2(smallInput_2)
	}
}
