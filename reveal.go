package main

import (
	"image"
	"image/color"
	"log"
)

func Reveal(img image.Image, opts *Opts) (string, int64) {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	availablePixels := make([]Pixel, 0, width*height/8)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := img.At(x, y).(color.NRGBA)
			if checkOptions(img, x, y, opts.Conditions) {
				availablePixels = append(availablePixels, Pixel{i, x, y})
			}
		}
	}

	applyTransformers(&availablePixels, opts.Transforms)

	result := make([]byte, 0, width*height/8)
	curr := byte(0)
	count := 0
	for _, pixel := range availablePixels {
		for _, c := range []byte{pixel.Color.R, pixel.Color.G, pixel.Color.B} {
			curr = decode(curr, c, count)
			count++
			if count&7 == 0 {
				result = append(result, curr)
				curr = byte(0)
			}
		}

	}

	lilInt := result[:4]
	size := bytesToInt64LE([4]byte(lilInt))

	log.Println(size, lilInt)

	return string(result[4:][:size]), size
}

func decode(char byte, c uint8, i int) uint8 {
	return char | ((c & 1) << (i & 7))
}
