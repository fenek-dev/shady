package main

import (
	"image/png"
	"log"
	"os"
)

func main() {
	var text = "Lorem Ipsum is simply dummy text of the printing and typesetting industry."

	fd, err := os.Open("peppo.png")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	img, _ := png.Decode(fd)

	opts := []Option{
		WithoutEmpty(),
		WithoutBreakingConsistency(),
	}

	newImg := Hide(img, text, opts...)

	CreateImage(newImg)

	result := Reveal(newImg, opts...)

	r := result[:len(text)]
	log.Println(r == text)
	log.Println(r)
}
