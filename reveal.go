package main

import (
	"image"
	"image/color"
)

func Reveal(img image.Image) string {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Iterate over each pixel
	result := make([]byte, 0, width*height/8)
	curr := byte(0)
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			im := img.At(x, y).(color.NRGBA)

			curr = decode(curr, im.R, count)
			curr = decode(curr, im.G, count+1)
			curr = decode(curr, im.B, count+2)
			curr = decode(curr, im.A, count+3)

			count += 4

			if count&7 == 0 {
				result = append(result, curr)
				curr = byte(0)
			}
		}
	}
	return string(result)
}

func decode(char byte, c uint8, i int) uint8 {
	return char | ((c & 1) << (i & 7))
}
