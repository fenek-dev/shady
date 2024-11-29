package main

import (
	"image"
	"image/color"
)

type Option func(img image.Image, c color.NRGBA, x, y int) bool

func checkOptions(img image.Image, x, y int, opts []Option) bool {
	for _, opt := range opts {
		if opt(img, img.At(x, y).(color.NRGBA), x, y) {
			return false
		}
	}
	return true
}

func WithoutEmpty() Option {
	return func(img image.Image, c color.NRGBA, x, y int) bool {
		return c.R <= 1 && c.G <= 1 && c.B <= 1
	}
}

func WithoutBreakingConsistency() Option {
	return func(img image.Image, c color.NRGBA, x, y int) bool {
		bounds := img.Bounds()
		width, height := bounds.Dx(), bounds.Dy()

		// get next pixel
		if x+1 < width {
			next := img.At(x+1, y).(color.NRGBA)
			return c.R == next.R && c.G == next.G && c.B == next.B
		} else if y+1 < height {
			next := img.At(0, y+1).(color.NRGBA)
			return c.R == next.R && c.G == next.G && c.B == next.B
		}

		return false
	}
}
