package entropy

/**************************************************************************
	This package will calculate entropy for given data string using:
	Probability model and iid(independent and identical distribution)
	assumption.
**************************************************************************/

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func getStringData() []byte {
	in := bufio.NewReader(os.Stdin)

	// \n is also included in the data.
	line, _ := in.ReadString('\n')

	return []byte(line)
}

func getFileData(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Something wrong happened while opening the file.")
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Something wrong happened while reading data from file.")
		os.Exit(2)
	}

	return data
}

func getDataForUsageMode() []byte {
	var mode int

	fmt.Println("(1) Find entropy of text entered in terminal.")
	fmt.Println("(2) Find entropy of file.")
	fmt.Println("Enter the corresponding number and hit enter.")

	_, err := fmt.Scan(&mode)
	if err != nil {
		fmt.Println("Something wrong happened. Please try again")
		os.Exit(1)
	}

	if mode != 1 && mode != 2 {
		fmt.Println("There is no mode for this number.")
		os.Exit(2)
	}

	if mode == 1 {
		data := getStringData()
		return data
	}
	fmt.Println("Drag and drop the file in terminal and then hit enter.")
	in := bufio.NewReader(os.Stdin)
	filePath, _ := in.ReadString('\n')

	correctFilePath := ""
	quoteCount := 0

	// The following code will find the correct path on
	// all operating systems.
	for _, pathChar := range filePath {
		if string(pathChar) == "'" {
			quoteCount++
			continue
		}
		if quoteCount == 2 {
			break
		}
		if pathChar == '\n' {
			continue
		}
		correctFilePath += string(pathChar)
	}

	data := getFileData(correctFilePath)
	return data
}

// DataFrequencyMap will take data from input
// and prepare a data frequency map.
func DataFrequencyMap() map[byte]int {
	frequencyMap := make(map[byte]int)

	data := getDataForUsageMode()

	for _, char := range data {
		// Initially the value is 0 for all characters and we
		// increase it by 1 whenever we encounter the character.
		frequencyMap[char] += 1
	}

	return frequencyMap
}

// DataProbabilityDistribution will return an array
// that holds the probability distribution of data
func DataProbabilityDistribution(dataFreqMap map[byte]int) []float32 {
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
func Entropy(dataFreqMap map[byte]int) float32 {
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
