package main

import (
	"encoding/binary"
	"image"
	"image/color"
	"log"
)

type Pixel struct {
	Color color.NRGBA
	X, Y  int
}

type ArgsOptions struct {
	IgnoreEmptyChannels bool
}

type Args struct {
	Conditions   []Condition
	Transformers []Transformer
	Noisers      []Transformer
	Options      *ArgsOptions
}

func Hide(img image.Image, text string, args *Args) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	newImg := image.NewNRGBA(bounds)

	// Get all available pixels
	// so we can transform that array after
	availablePixels := make([]Pixel, 0, width*height/8)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := img.At(x, y).(color.NRGBA)
			if checkOptions(img, x, y, args.Conditions) {
				availablePixels = append(availablePixels, Pixel{i, x, y})
			}
			newImg.Set(x, y, i)
		}
	}

	// Noisers must be applied before the transformers
	applyTransformers(&availablePixels, args.Noisers)

	applyTransformers(&availablePixels, args.Transformers)

	// Add the length of the text to the beginning of the text
	lilInt := int64ToBytesLE(int64(len(text)))
	log.Println(len(text), lilInt)
	t := string(lilInt[:]) + text

	count := 0
	for _, pixel := range availablePixels {
		if count/8 <= len(t)-1 {
			arr := []byte{pixel.Color.R, pixel.Color.G, pixel.Color.B}
			for arri, c := range arr {
				if args.Options.IgnoreEmptyChannels && c <= 1 {
					continue
				}
				if count/8 >= len(t) {
					continue
				}
				char := t[count/8]
				arr[arri] = encode(char, c, count)
				count++
			}

			c := color.NRGBA{
				R: arr[0], G: arr[1], B: arr[2], A: pixel.Color.A,
			}

			newImg.Set(pixel.X, pixel.Y, c)
		} else {
			newImg.Set(pixel.X, pixel.Y, pixel.Color)
		}

	}

	return newImg
}

func encode(char byte, c uint8, idx int) uint8 {
	if char&(1<<(idx&7)) != 0 {
		return c | 1
	} else {
		return c & 254
	}
}

func int64ToBytesLE(n int64) [4]byte {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], uint32(n))
	return b
}

func bytesToInt64LE(b [4]byte) int64 {
	return int64(binary.LittleEndian.Uint32(b[:]))
}
