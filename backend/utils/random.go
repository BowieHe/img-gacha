package utils

import (
	"math/rand"
	"time"
)

// init initializes the random number generator with current time
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateRandomSeed returns a random seed for image generation
func GenerateRandomSeed() int64 {
	return rand.Int63()
}

// GenerateRandomSeeds generates multiple random seeds
func GenerateRandomSeeds(count int) []int64 {
	seeds := make([]int64, count)
	for i := 0; i < count; i++ {
		seeds[i] = GenerateRandomSeed()
	}
	return seeds
}

// DeriveSeeds generates deterministic seeds from a base seed
// This ensures reproducibility when needed
func DeriveSeeds(baseSeed int64, count int) []int64 {
	seeds := make([]int64, count)
	rng := rand.New(rand.NewSource(baseSeed))
	for i := 0; i < count; i++ {
		seeds[i] = rng.Int63()
	}
	return seeds
}
