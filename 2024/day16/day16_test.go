package day16

import (
	"adventofcode/helper"
	"embed"
	"testing"
)

//go:embed *.txt
var embed_fs embed.FS

func TestStar1(t *testing.T) {
	helper.TestStar(t, Star1, embed_fs, "7036", "95476")
}

func BenchmarkStar1(b *testing.B) {
	helper.BenchmarkStar(b, Star1, embed_fs)
}

func TestStar2(t *testing.T) {
	helper.TestStar(t, Star2, embed_fs, "45", "511")
}

func BenchmarkStar2(b *testing.B) {
	helper.BenchmarkStar(b, Star2, embed_fs)
}
