package main

import (
	"strings"
	"unicode"
)

func RemoveEven(slice []int) []int {
	var ans []int
	for i := range slice {
		if slice[i] % 2 != 0 {
			ans = append(ans, slice[i])
		}
	}
	return ans
}

func PowerGenerator(base int) (func() int) {
	val := 1
	return func() (ans int) {
		val *= base
		ans = val
		return
	}
}

func DifferentWordsCount(text string) int {
	runes := append([]rune(strings.ToLower(text)), rune(' '))
	words := make(map[string]int)
	lastIndex := -1
	for idx, r := range runes {
		if !unicode.IsLetter(r) {
			if idx-lastIndex > 1 {
				words[string(runes[lastIndex+1:idx])] = idx
			}
			lastIndex = idx
		}
	}
	return len(words)
}
