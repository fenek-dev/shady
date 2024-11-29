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

			if count/8 <= len(text)-1 && !isEmpty(i) {
				arr := []byte{i.R, i.G, i.B}
				for arri, c := range arr {
					if count/8 >= len(text) {
						continue
					}
					char := text[count/8]
					arr[arri] = transform(char, c, count)
					count++
				}

				c := color.NRGBA{
					R: arr[0], G: arr[1], B: arr[2], A: i.A,
				}

				newImg.Set(x, y, c)
			} else {
				newImg.Set(x, y, i)
			}

		}
	}

	return newImg
}

func transform(char byte, c uint8, idx int) uint8 {
	if char&(1<<(idx&7)) != 0 {
		return c | 1
	} else {
		return c & 254
	}
}

func isEmpty(i color.NRGBA) bool {
	return i.R <= 1 && i.G <= 1 && i.B <= 1
}
