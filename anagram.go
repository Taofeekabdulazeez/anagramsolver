package main

import "strings"

func GetCharacterFrequency(str string) map[rune]int {
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

func IsAnagram(str, substr string) bool {
	str = strings.ToLower(str)
	substr = strings.ToLower(substr)

	if len(substr) > len(str) {
		return false
	}

	strFreq := GetCharacterFrequency(str)
	substrFreq := GetCharacterFrequency(substr)

	for key, value := range substrFreq {
		if count, exist := strFreq[key]; !exist || value > count {
			return false
		}
	}

	return true
}

func GetUniqueCharacters(str string) (unique_chars []rune) {
	found := make(map[rune]bool)

	for _, ch := range str {
		if !found[ch] {
			unique_chars = append(unique_chars, ch)
			found[ch] = true
		}
	}

	return unique_chars
}

func FindAnagrams(word string) (anagrams []string) {
	chars := GetUniqueCharacters(word)

	for _, ch := range chars {
		words, _ := ReadAngramFile(ch)
		for _, str := range words {
			if IsAnagram(word, str) {
				anagrams = append(anagrams, str)
			}
		}
	}

	return anagrams
}
