package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type StopWordsManager struct {
	ActiveWFObjects
	stop_words map[string]struct{}
}

func (swm *StopWordsManager) InitializeStopWordManager(ch chan []interface{}) {
	swm.stop_words = map[string]struct{}{}
	swm.queue = ch
}

func (swm *StopWordsManager) ReadStopWordsFile(path string) {
	stop_words_file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	stop_words_data, err := io.ReadAll(stop_words_file)
	if err != nil {
		log.Fatal(err)
	}
	stop_words_list := strings.Split(string(stop_words_data), ",")

	// normalize
	for i, stop_word := range stop_words_list {
		stop_words_list[i] = strings.ToLower(stop_word)
	}

	// store them in a map for better accessability
	for _, stop_word := range stop_words_list {
		swm.stop_words[stop_word] = struct{}{}
	}

	// add ascii chars to stop words
	for i := 33; i <= 126; i++ {
		swm.stop_words[string(rune(i))] = struct{}{}
	}
}

func (swm *StopWordsManager) IsStopWord(word string) bool {
	_, found := swm.stop_words[word]
	return found
}

func (swm *StopWordsManager) RemoveStopWordsFromASlice(words []string) []string {
	filtered_list := []string{}
	for _, word := range words {
		if !swm.IsStopWord(word) {
			filtered_list = append(filtered_list, word)
		}
	}
	return filtered_list
}
