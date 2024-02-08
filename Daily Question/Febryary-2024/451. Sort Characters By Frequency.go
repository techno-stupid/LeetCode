package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "tReeerdfsf"
	fmt.Println(frequencySort(str))
}

type CharFrequency struct {
	Character byte
	Frequency int
}

func frequencySort(s string) string {
	charFrequencyMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		charFrequencyMap[s[i]]++
	}

	charFrequencies := make([]CharFrequency, 0)
	for char, freq := range charFrequencyMap {
		charFrequencies = append(charFrequencies, CharFrequency{
			Character: char,
			Frequency: freq,
		})
	}
	sort.Slice(charFrequencies, func(i, j int) bool {
		return charFrequencies[i].Frequency > charFrequencies[j].Frequency
	})

	result := ""
	for _, cf := range charFrequencies {
		for i := 0; i < cf.Frequency; i++ {
			result += string(cf.Character)
		}
	}
	return result
}
