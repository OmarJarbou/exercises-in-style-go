package main

import "strings"

func splitWords(text string, stop_words map[string]struct{}, ch chan []string) {
	words := []string{}
	temp_word := ""
	for _, char := range text {
		if (string(char) >= "a" && string(char) <= "z") || (string(char) >= "A" && string(char) <= "Z") || (string(char) >= "0" && string(char) <= "9") {
			temp_word += string(char)
			continue
		}

		if temp_word == "" {
			continue
		}

		word_lower := strings.ToLower(temp_word)
		if _, ok := stop_words[word_lower]; !ok {
			words = append(words, temp_word)
		}
		temp_word = ""
	}

	ch <- words
}
