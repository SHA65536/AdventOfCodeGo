package day03

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

	res, err := Star1(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "4361" {
		log.Println(res)
		t.FailNow()
	}

	input, err = helper.NewInputReader("input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err = Star1(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "557705" {
		log.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input, _ := helper.NewInputReader("small_input.txt")
		Star1(input)
	}
}

func TestStar2(t *testing.T) {
	input, err := helper.NewInputReader("small_input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err := Star2(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "467835" {
		log.Println(res)
		t.FailNow()
	}

	input, err = helper.NewInputReader("input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err = Star2(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "84266818" {
		log.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input, _ := helper.NewInputReader("small_input.txt")
		Star2(input)
	}
}
