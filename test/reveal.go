package test

import (
	"image"
	"image/color"
)

func Reveal1(img image.Image) string {
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

func Reveal2(img image.Image) string {
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

			if count%8 == 0 {
				result = append(result, curr)
				curr = byte(0)
				count = 0
			}
		}
	}
	return string(result)
}

func decode(char byte, c uint8, i int) uint8 {
	return char | ((c & 1) << (i % 8))
}
