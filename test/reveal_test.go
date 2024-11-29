package test

import (
	"image/png"
	"os"
	"testing"
)

func BenchmarkReveal1(b *testing.B) {
	file, _ := os.Open("output.png")
	img, _ := png.Decode(file)
	for i := 0; i < b.N; i++ {
		r := Reveal1(img)
		if r[:len(text)] != text {
			b.Fatal("Reveal1 failed")
		}
	}
}

func BenchmarkReveal2(b *testing.B) {
	file, _ := os.Open("output.png")
	defer file.Close()
	img, _ := png.Decode(file)
	for i := 0; i < b.N; i++ {
		r := Reveal2(img)
		if r[:len(text)] != text {
			b.Fatal("Reveal2 failed")
		}
	}
}

func BenchmarkReveal3(b *testing.B) {
	file, _ := os.Open("output.png")
	defer file.Close()
	img, _ := png.Decode(file)
	for i := 0; i < b.N; i++ {
		r := Reveal3(img)
		if r[:len(text)] != text {
			b.Fatal("Reveal3 failed")
		}
	}
}
