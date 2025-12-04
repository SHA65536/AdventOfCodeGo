package day04

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
	if res != "13" {
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
	if res != "1409" {
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
	if res != "43" {
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
	if res != "8366" {
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

func TestStar2SinglePass(t *testing.T) {
	input, err := helper.NewInputReaderEmbed(embed_fs, "small_input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err := Star2SinglePass(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "43" {
		log.Println(res)
		t.FailNow()
	}

	input, err = helper.NewInputReaderEmbed(embed_fs, "input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err = Star2SinglePass(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != "8366" {
		log.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2SinglePass(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input, _ := helper.NewInputReaderEmbed(embed_fs, "input.txt")
		Star2SinglePass(input)
	}
}
