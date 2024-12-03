package main

import (
	"image/color"
	"math/rand/v2"
)

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Add noise to a pixel color
func addNoiseToColor(c color.NRGBA, noiseRange int) color.NRGBA {
	// Generate random noise for each channel
	noiseR := rand.IntN(noiseRange)
	noiseG := rand.IntN(noiseRange)
	noiseB := rand.IntN(noiseRange)

	// Add noise and clamp the values
	newR := clamp(int(c.R)+noiseR, 0, 255)
	newG := clamp(int(c.G)+noiseG, 0, 255)
	newB := clamp(int(c.B)+noiseB, 0, 255)

	return color.NRGBA{R: uint8(newR), G: uint8(newG), B: uint8(newB), A: c.A}
}

func addNoiseToColorWithoutEmptyChannels(c color.NRGBA, noiseRange int) color.NRGBA {
	newColor := addNoiseToColor(c, noiseRange)
	if c.R == 0 {
		newColor.R = 0
	}
	if c.G == 0 {
		newColor.G = 0
	}
	if c.B == 0 {
		newColor.B = 0
	}

	return newColor
}
