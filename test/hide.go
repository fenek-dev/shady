package test

import (
	"image"
	"image/color"
)

var text = "I love you!"

func Hide1(img image.Image) image.Image {

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
					R: transform1(char, i.R, count%8),
					G: transform1(char, i.G, (count+1)%8),
					B: transform1(char, i.B, (count+2)%8),
					A: transform1(char, i.A, (count+3)%8),
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

func transform1(char byte, c uint8, idx int) uint8 {
	if char&(1<<idx) != 0 {
		return c | 1
	} else {
		return c & 0
	}
}

func Hide2(img image.Image) image.Image {

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
					R: transform2(char, i.R, count),
					G: transform2(char, i.G, count+1),
					B: transform2(char, i.B, count+2),
					A: transform2(char, i.A, count+3),
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

func transform2(char byte, c uint8, idx int) uint8 {
	if char&(1<<(idx&7)) != 0 {
		return c | 1
	} else {
		return c & 0
	}
}
