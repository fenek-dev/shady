package main

import (
	"image/png"
	"log"
	"os"
)

var text = "I love you!"

func main() {
	fd, err := os.Open("peppo.png")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	img, _ := png.Decode(fd)

	newImg := Hide(img)
	CreateImage(newImg)

	result := Reveal(newImg)
	log.Println(result[:len(text)])
}
