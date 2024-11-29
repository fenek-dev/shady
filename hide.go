package main

import (
	"encoding/binary"
	"image"
	"image/color"
	"log"
)

func Hide(img image.Image, text string, opts ...Option) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	newImg := image.NewNRGBA(bounds)

	lilInt := int64ToBytesLE(int64(len(text)))
	log.Println(len(text), lilInt)
	t := string(lilInt[:]) + text
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := img.At(x, y).(color.NRGBA)

			if count/8 <= len(t)-1 && checkOptions(img, x, y, opts) {
				arr := []byte{i.R, i.G, i.B}
				for arri, c := range arr {
					if count/8 >= len(t) {
						continue
					}
					char := t[count/8]
					arr[arri] = encode(char, c, count)
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
