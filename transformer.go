package main

import (
	"slices"
)

type Transformer func(pixels *[]Pixel)

func applyTransformers(pixels *[]Pixel, transformers []Transformer) {
	for _, t := range transformers {
		t(pixels)
	}
}

func ReverseTransform() Transformer {
	return func(pixels *[]Pixel) {
		slices.Reverse(*pixels)
	}
}

func SimpleEllipticTransformer(seed []byte) Transformer {
	return func(pixels *[]Pixel) {
		shuffleSliceDeterministic(*pixels, seed)
	}
}

func SimpleNoiseTransformer(value int) Transformer {
	return func(pixels *[]Pixel) {
		for i, p := range *pixels {
			c := addNoiseToColor(p.Color, value)
			(*pixels)[i].Color = c
		}
	}
}

func SimpleNoiseWithoutEmptyChannelsTransformer(value int) Transformer {
	return func(pixels *[]Pixel) {
		for i, p := range *pixels {
			c := addNoiseToColorWithoutEmptyChannels(p.Color, value)
			(*pixels)[i].Color = c
		}
	}
}
