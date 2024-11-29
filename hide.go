package main

import (
	"image"
	"image/color"
)

func Hide(img image.Image) image.Image {

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	newImg := image.NewNRGBA(bounds)

	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := img.At(x, y).(color.NRGBA)

			index := count / 8
			if index <= len(text)-1 {
				char := text[index]
				c := color.NRGBA{
					R: transform(char, i.R, count%8),
					G: transform(char, i.G, (count+1)%8),
					B: transform(char, i.B, (count+2)%8),
					A: transform(char, i.A, (count+3)%8),
				}
				newImg.Set(x, y, c)
			} else {
				newImg.Set(x, y, i)
			}

			count += 4

		}
	}

	return newImg
}

func transform(char byte, c uint8, idx int) uint8 {
	if char&(1<<idx) != 0 {
		return c | 1
	} else {
		return c & 0
	}
}
