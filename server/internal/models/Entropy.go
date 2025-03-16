package main

import (
	"fmt"
	"math"
)

func calculateDirectionEntropy(userVotes map[int][4]int) float64 {
	totalVotes := [4]float64{0, 0, 0, 0}
	for _, votes := range userVotes {
		for dir := 0; dir < 4; dir++ {
			totalVotes[dir] += float64(votes[dir])
		}
	}

	total := 0.0
	for _, count := range totalVotes {
		total += count
	}

	if total == 0 {
		return 0
	}

	entropy := 0.0
	for _, count := range totalVotes {
		if count > 0 {
			p := count / total
			entropy -= p * math.Log2(p)
		}
	}

	maxEntropy := math.Log2(4)
	normalizedEntropy := entropy / maxEntropy

	return normalizedEntropy
}