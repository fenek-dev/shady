package main

import "slices"

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
