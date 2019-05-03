package entropy

/**************************************************************************
	This package will calculate entropy for given data string using:
	Probability model and iid(independent and identical distribution)
	assumption.
**************************************************************************/

import (
	"bufio"
	"math"
	"os"
)

func getStringData() string {
	in := bufio.NewReader(os.Stdin)

	// \n is also included in the data.
	line, _ := in.ReadString('\n')

	return line
}

// DataFrequencyMap will take data from input
// and prepare a data frequency map.
func DataFrequencyMap() map[interface{}]int {
	frequencyMap := make(map[interface{}]int)

	data := getStringData()

	for _, char := range data {
		// Initially the value is 0 for all characters and we
		// increase it by 1 whenever we encounter the character.
		frequencyMap[char] += 1
	}

	return frequencyMap
}

// DataProbabilityDistribution will return an array
// that holds the probability distribution of data
func DataProbabilityDistribution(dataFreqMap map[interface{}]int) []float32 {
	var probabilityDist []float32
	var totalChars int

	for _, freq := range dataFreqMap {
		totalChars += freq
	}

	for _, freq := range dataFreqMap {
		probabilityDist = append(probabilityDist, float32(freq)/float32(totalChars))
	}

	return probabilityDist
}

// Entropy will calculate the entropy for given
// data frequency map and return float32 value
// which represents entropy as (x)bits/sample
func Entropy(dataFreqMap map[interface{}]int) float32 {
	var entropy float32

	probDist := DataProbabilityDistribution(dataFreqMap)

	// entropy = sigma(-P(Ai)log(P(Ai)))
	/*
	 base of log  |   unit of entropy
	      2       |    bits/sample
	      e       |    nats/sample
	      10      |   hartleys/sample
	*/

	for _, probability := range probDist {
		entropy -= probability * float32(math.Log2(float64(probability)))
	}

	return entropy
}
