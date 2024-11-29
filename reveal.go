package main

import (
	"image"
	"image/color"
)

func Reveal(img image.Image) string {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	pixelData := make([]byte, 0, width*height*4)

	// Iterate over each pixel
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := img.At(x, y).(color.NRGBA)
			pixelData = append(pixelData, i.R, i.G, i.B, i.A)
		}
	}
	curr := byte(0)
	result := make([]byte, 0, len(pixelData)/8)
	for i, c := range pixelData {
		curr = curr | ((c & 1) << (i % 8))
		if (i+1)%8 == 0 {
			result = append(result, curr)
			curr = byte(0)
			continue
		}
	}
	return string(result)
}
