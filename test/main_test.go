package test

import (
	"image/png"
	"os"
	"testing"
)

func BenchmarkRevealOld(b *testing.B) {
	file, _ := os.Open("output.png")
	defer file.Close()
	img, _ := png.Decode(file)
	for i := 0; i < b.N; i++ {
		Reveal1(img)
	}
}

func BenchmarkRevealNew(b *testing.B) {
	file, _ := os.Open("output.png")
	defer file.Close()
	img, _ := png.Decode(file)
	for i := 0; i < b.N; i++ {
		Reveal2(img)
	}
}
