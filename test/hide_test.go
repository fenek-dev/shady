package test

import (
	"image/png"
	"os"
	"testing"
)

func BenchmarkHide1(b *testing.B) {
	file, _ := os.Open("input.png")
	img, _ := png.Decode(file)
	for i := 0; i < b.N; i++ {
		Hide1(img)
	}
}

func BenchmarkHide2(b *testing.B) {
	file, _ := os.Open("input.png")
	img, _ := png.Decode(file)
	for i := 0; i < b.N; i++ {
		Hide2(img)
	}
}
