package day02

import (
	"adventofcode/helper"
	"log"
	"testing"
)

func TestStar1(t *testing.T) {
	input, err := helper.NewInputReader("small_input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err := CubesGame(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "8" {
		log.Println(res)
		t.FailNow()
	}

	input, err = helper.NewInputReader("input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err = CubesGame(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "2265" {
		log.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input, _ := helper.NewInputReader("small_input.txt")
		CubesGame(input)
	}
}

func TestStar2(t *testing.T) {
	input, err := helper.NewInputReader("small_input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err := CubesGameFewest(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "2286" {
		log.Println(res)
		t.FailNow()
	}

	input, err = helper.NewInputReader("input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err = CubesGameFewest(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "64097" {
		log.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input, _ := helper.NewInputReader("small_input.txt")
		CubesGameFewest(input)
	}
}
