package day02

import (
	"adventofcode/helper"
	"embed"
	"testing"
)

//go:embed *.txt
var embed_fs embed.FS

func TestStar1(t *testing.T) {
	helper.TestStar(t, Star1, embed_fs, "2", "202")
}

func BenchmarkStar1(b *testing.B) {
	helper.BenchmarkStar(b, Star1, embed_fs)
}

func TestStar2(t *testing.T) {
	helper.TestStar(t, Star2, embed_fs, "4", "271")
}

func BenchmarkStar2(b *testing.B) {
	helper.BenchmarkStar(b, Star2, embed_fs)
}
