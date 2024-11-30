package main

import (
	"crypto/elliptic"
	"crypto/sha256"
	"math/big"
)

// Deterministic random number generator using a fixed private key and counter
func deterministicRandomInt(seed []byte, counter int, max int) int {
	// Use elliptic curve P256
	curve := elliptic.P256()

	// Create a deterministic private key based on the seed and counter
	seedWithCounter := append(seed, byte(counter))
	hash := sha256.Sum256(seedWithCounter)

	// Use the hash as the private key
	d := new(big.Int).SetBytes(hash[:])
	d = d.Mod(d, curve.Params().N) // Ensure it's within the curve's order

	// Generate a deterministic random number within the range [0, max)
	randomInt := d.Mod(d, big.NewInt(int64(max))).Int64()
	return int(randomInt)
}

// Deterministically shuffle a slice using ECC
func shuffleSliceDeterministic(slice []Pixel, seed []byte) {
	for i := range slice {
		// Generate a deterministic random index
		randIndex := deterministicRandomInt(seed, i, len(slice))

		// Swap the current element with the random index
		slice[i], slice[randIndex] = slice[randIndex], slice[i]
	}
}
