package main

import (
	"image"
	"image/color"
)

func Reveal(img image.Image, opts ...Option) string {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	result := make([]byte, 0, width*height/8)
	curr := byte(0)
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			im := img.At(x, y).(color.NRGBA)
			if !checkOptions(img, x, y, opts) {
				continue
			}

			for _, c := range []byte{im.R, im.G, im.B} {
				curr = decode(curr, c, count)
				count++
				if count&7 == 0 {
					result = append(result, curr)
					curr = byte(0)
				}
			}

		}
	}
	return string(result)
}

func decode(char byte, c uint8, i int) uint8 {
	return char | ((c & 1) << (i & 7))
}
