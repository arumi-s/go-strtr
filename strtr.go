package strtr

import (
	"fmt"
	"slices"
	"sort"
)

func Strtr(input string, from map[string]string) string {
	if len(from) == 0 {
		return input
	}

	// Sort the keys by their length in descending order for priority
	sortedKeys := make([]string, 0, len(from))
	for key := range from {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Slice(sortedKeys, func(i, j int) bool {
		return len(sortedKeys[i]) > len(sortedKeys[j])
	})

	// Create an array of characters from the input string to facilitate replacements
	characters := []rune(input)
	fmt.Println(characters)

	// Track which indices have been replaced to prevent re-evaluation
	replacedIndices := make([]bool, len(input)*2)

	// Process each key according to its sorted order
	for _, key := range sortedKeys {
		index := 0
		keyLength := len([]rune(key))
		for index <= len(characters)-keyLength {
			// Check if current segment matches key and has not been replaced
			if !slices.Contains(replacedIndices[index:index+keyLength], true) && string(characters[index:index+keyLength]) == key {
				// Replace the segment
				to := []rune(from[key])
				characters = append(characters[:index], append(to, characters[index+keyLength:]...)...)
				// Update the replaced indices
				toLength := len(to)
				newReplacedIndices := make([]bool, toLength)
				for i := 0; i < toLength; i++ {
					newReplacedIndices[i] = true
				}
				replacedIndices = append(replacedIndices[:index], append(newReplacedIndices, replacedIndices[index+keyLength:]...)...)
				// Move index forward by the length of the replacement
				index += toLength
			} else {
				index += 1 // Move to the next character
			}
		}
	}

	return string(characters)
}
