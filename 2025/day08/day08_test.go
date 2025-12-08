package day08

import (
	"adventofcode/helper"
	"embed"
	"log"
	"testing"
)

//go:embed *.txt
var embed_fs embed.FS

func TestStar1(t *testing.T) {
	input, err := helper.NewInputReaderEmbed(embed_fs, "small_input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err := Star1(input, 10)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "40" {
		log.Println(res)
		t.FailNow()
	}

	input, err = helper.NewInputReaderEmbed(embed_fs, "input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err = Star1(input, 1000)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "47040" {
		log.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input, _ := helper.NewInputReaderEmbed(embed_fs, "input.txt")
		Star1(input, 1000)
	}
}

func TestStar2(t *testing.T) {
	input, err := helper.NewInputReaderEmbed(embed_fs, "small_input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err := Star2(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "25272" {
		log.Println(res)
		t.FailNow()
	}

	input, err = helper.NewInputReaderEmbed(embed_fs, "input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err = Star2(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "4884971896" {
		log.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input, _ := helper.NewInputReaderEmbed(embed_fs, "input.txt")
		Star2(input)
	}
}
