package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

//var m map[byte]int

// 0b01100001

// 0b00000000 0b00000000 0b00000000 0b00000000
// 0b00000000 0b00000000 0b00000000 0b00000000

var text = "Hello, World!"

func main() {
	fd, err := os.Open("peppo.png")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	//cfg, _ := png.DecodeConfig(fd)
	//_ = cfg
	//cfg.ColorModel.Convert(color.NRGBA{})
	// Decode the PNG file
	img, err := png.Decode(fd)
	if err != nil {
		log.Fatalf("Failed to decode PNG: %v", err)
	}

	// Access the raw pixel data
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	pixelData := make([]byte, 0, width*height*4)
	//pixelData := nBytes(8 * len(text))

	// Iterate over each pixel
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := img.At(x, y).(color.NRGBA)
			pixelData = append(pixelData, i.R, i.G, i.B, i.A)
		}
	}

	// 0b01100001 0b01100010
	//log.Println(pixelData)
	count := 0
	t := []byte(text)
	for _, c := range t {
		for idx := range 8 {
			mask := bitmaskForIndex(c, idx)
			if mask == 1 {
				pixelData[count] = pixelData[count] | 1
			} else {
				pixelData[count] = pixelData[count] & 0
			}
			count++
		}
	}
	//log.Println(pixelData)

	newImg := image.NewNRGBA(bounds)
	for i := 0; i < len(pixelData); i += 4 {
		c := color.NRGBA{
			R: pixelData[i],
			G: pixelData[i+1],
			B: pixelData[i+2],
			A: pixelData[i+3],
		}
		newImg.Set(i/4%width, i/4/width, c)
	}

	outFile, err := os.Create("output.png")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outFile.Close()

	// Encode the modified image into the new file
	err = png.Encode(outFile, newImg)
	if err != nil {
		log.Fatalf("Failed to encode PNG: %v", err)
	}

	//curr := byte(0)
	//result := make([]byte, 0, len(pixelData)/8)
	//for i, c := range pixelData {
	//	curr = curr | ((c & 1) << (i % 8))
	//	if (i+1)%8 == 0 {
	//		result = append(result, curr)
	//		curr = byte(0)
	//		continue
	//	}
	//}
	//
	//log.Println(string(result), len(result))

	// Process or save pixelData as needed
	log.Printf("Extracted %d pixels from the image\n", len(pixelData))
}

// 0b01100001 5 -> 0b00000000
func bitmaskForIndex(b byte, index int) byte {
	if b&(1<<index) != 0 {
		return 1
	}
	return 0
}

func nBytes(n int) []byte {
	r := make([]byte, 0, n)
	for _ = range n {
		r = append(r, 0)
	}
	return r
}
