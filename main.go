package main

import (
	"image/png"
	"log"
	"os"
)

func main() {
	key := []byte("thisis32byteslongpassphrase12341")
	seed := []byte("thisisalongseed")
	text := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum is simply dummy text of the printing and"

	cipher, err := EncryptAES(text, key)
	if err != nil {
		panic(err)
	}

	fd, err := os.Open("peppo.png")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	img, _ := png.Decode(fd)

	opts := &Opts{
		Conditions: []Condition{
			EmptyCondition(),
		},
		Transforms: []Transformer{
			ReverseTransform(),
			SimpleEllipticTransformer(seed),
		},
	}

	newImg := Hide(img, cipher, opts)

	CreateImage(newImg)

	result, size := Reveal(newImg, opts)
	result, err = DecryptAES(result[:], key)
	if err != nil {
		panic(err)
	}

	r := result
	log.Println(r == text, size == int64(len(cipher)))
	log.Println(r)
}
