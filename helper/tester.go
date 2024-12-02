package helper

import (
	"embed"
	"log"
	"testing"
)

type Solution func(*InputReader) (string, error)

func TestStar(t *testing.T, sol Solution, fs embed.FS, res_small, res_big string) {
	input, err := NewInputReaderEmbed(fs, "small_input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err := sol(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != res_small {
		log.Println(res)
		t.FailNow()
	}

	input, err = NewInputReaderEmbed(fs, "input.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	res, err = sol(input)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if res != res_big {
		log.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar(b *testing.B, sol Solution, fs embed.FS) {
	for i := 0; i < b.N; i++ {
		input, _ := NewInputReaderEmbed(fs, "input.txt")
		sol(input)
	}
}
