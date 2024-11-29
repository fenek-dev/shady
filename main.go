package main

import (
	"image/png"
	"log"
	"os"
)

var text = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."

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
	log.Println(result[:len(text)] == text)
}
