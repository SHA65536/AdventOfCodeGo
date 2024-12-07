package day07

import (
	"adventofcode/helper"
	"embed"
	"testing"
)

//go:embed *.txt
var embed_fs embed.FS

func TestStar1(t *testing.T) {
	helper.TestStar(t, Star1, embed_fs, "3749", "1620690235709")
}

func BenchmarkStar1(b *testing.B) {
	helper.BenchmarkStar(b, Star1, embed_fs)
}

func TestStar2(t *testing.T) {
	helper.TestStar(t, Star2, embed_fs, "11387", "145397611075341")
}

func BenchmarkStar2(b *testing.B) {
	helper.BenchmarkStar(b, Star2, embed_fs)
}

func TestStar1Iter(t *testing.T) {
	helper.TestStar(t, Star1Iter, embed_fs, "3749", "1620690235709")
}

func BenchmarkStar1Iter(b *testing.B) {
	helper.BenchmarkStar(b, Star1Iter, embed_fs)
}

func TestStar2Iter(t *testing.T) {
	helper.TestStar(t, Star2Iter, embed_fs, "11387", "145397611075341")
}

func BenchmarkStar2Iter(b *testing.B) {
	helper.BenchmarkStar(b, Star2Iter, embed_fs)
}
