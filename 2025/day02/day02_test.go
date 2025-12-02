package day02

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

	res, err := Star1(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "1227775554" {
		log.Println(res)
		t.FailNow()
	}

	input, err = helper.NewInputReaderEmbed(embed_fs, "input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err = Star1(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "12850231731" {
		log.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input, _ := helper.NewInputReaderEmbed(embed_fs, "input.txt")
		Star1(input)
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
	if res != "4174379265" {
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
	if res != "24774350322" {
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
