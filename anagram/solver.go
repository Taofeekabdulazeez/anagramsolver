package anagram

import (
	"strings"

	"github.com/Taofeekabdulazeez/anagramsolver/internal/util"
)

func FindAnagrams(word string) (anagrams []string) {
	chars := getUniqueCharacters(word)

	for _, ch := range chars {
		words, _ := util.ReadAngramFile(ch)
		for _, str := range words {
			if IsAnagram(word, str) {
				anagrams = append(anagrams, str)
			}
		}
	}

	return anagrams
}

func IsAnagram(str, substr string) bool {
	str = strings.ToLower(str)
	substr = strings.ToLower(substr)

	if len(substr) > len(str) {
		return false
	}

	strFreq := getCharacterFrequency(str)
	substrFreq := getCharacterFrequency(substr)

	for key, value := range substrFreq {
		if count, exist := strFreq[key]; !exist || value > count {
			return false
		}
	}

	return true
}

func getCharacterFrequency(str string) map[rune]int {
	freq := make(map[rune]int)
	for _, ch := range str {
		if value, exist := freq[ch]; !exist {
			freq[ch] = 1
		} else {
			freq[ch] = value + 1
		}
	}
	return freq
}

func getUniqueCharacters(str string) (unique_chars []rune) {
	found := make(map[rune]bool)

	for _, ch := range str {
		if !found[ch] {
			unique_chars = append(unique_chars, ch)
			found[ch] = true
		}
	}

	return unique_chars
}
