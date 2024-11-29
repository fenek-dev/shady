package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func CreateImage(img image.Image) {
	outFile, err := os.Create("output.png")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer func() {
		err = outFile.Close()
		if err != nil {
			log.Fatalf("Failed to close output file: %v", err)
		}
	}()

	// Encode the modified image into the new file
	err = png.Encode(outFile, img)
	if err != nil {
		log.Fatalf("Failed to encode PNG: %v", err)
	}
}
