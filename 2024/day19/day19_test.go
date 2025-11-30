package day19

import (
	"adventofcode/helper"
	"embed"
	"testing"
)

//go:embed *.txt
var embed_fs embed.FS

func TestStar1(t *testing.T) {
	helper.TestStar(t, Star1, embed_fs, "6", "269")
}

func BenchmarkStar1(b *testing.B) {
	helper.BenchmarkStar(b, Star1, embed_fs)
}

func TestStar2(t *testing.T) {
	helper.TestStar(t, Star2, embed_fs, "16", "758839075658876")
}

func BenchmarkStar2(b *testing.B) {
	helper.BenchmarkStar(b, Star2, embed_fs)
}

func TestStar1Map(t *testing.T) {
	helper.TestStar(t, Star1Map, embed_fs, "6", "269")
}

func BenchmarkStar1Map(b *testing.B) {
	helper.BenchmarkStar(b, Star1Map, embed_fs)
}

func TestStar2Map(t *testing.T) {
	helper.TestStar(t, Star2Map, embed_fs, "16", "758839075658876")
}

func BenchmarkStar2Map(b *testing.B) {
	helper.BenchmarkStar(b, Star2Map, embed_fs)
}
