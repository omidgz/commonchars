package main

import (
	"fmt"
	"math"
)

func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	freqMap := make(map[rune]int)
	for _, c := range words[0] {
		freqMap[c]++
	}

	for _, word := range words[1:] {
		currentMap := make(map[rune]int)
		for _, c := range word {
			currentMap[c]++
		}

		for char := range freqMap {
			if currentMap[char] > 0 {
				freqMap[char] = int(math.Min(float64(freqMap[char]), float64(currentMap[char])))
			} else {
				delete(freqMap, char)
			}
		}
	}

	result := []string{}
	for char, count := range freqMap {
		for i := 0; i < count; i++ {
			result = append(result, string(char))
		}
	}

	return result
}

func main() {
	fmt.Println(commonChars([]string{"philadelphia", "constantinopol", "persepolis"}))
}
