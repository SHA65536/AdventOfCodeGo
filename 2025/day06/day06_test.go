package day06

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
	if res != "4277556" {
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
	if res != "5977759036837" {
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
	if res != "3263827" {
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
	if res != "9630000828442" {
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
