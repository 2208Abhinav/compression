package main

import (
	"compression/entropy"
	"fmt"
)

func main() {
	dataFreqMap := entropy.DataFrequencyMap()
	dataEntropy := entropy.Entropy(dataFreqMap)

	fmt.Printf("Entropy: %.3f bits/sample\n", dataEntropy)
}
