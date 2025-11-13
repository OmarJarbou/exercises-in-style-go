package main

import (
	"sort"
)

func Reduce(countWords func([]string, chan map[string]int), splits [][]string) []WordFreq {
	words_freqs_maps := []map[string]int{}
	channels := make([]chan map[string]int, len(splits))
	for i, split := range splits {
		channels[i] = make(chan map[string]int)
		go countWords(split, channels[i])
	}
	for i, _ := range splits {
		words_freqs_map := <-channels[i]
		words_freqs_maps = append(words_freqs_maps, words_freqs_map)
	}

	result_map := map[string]int{}
	for _, words_freqs_map := range words_freqs_maps {
		for word, _ := range words_freqs_map {
			result_map[word] += words_freqs_map[word]
		}
	}

	words_freqs := []WordFreq{}
	for word, freq := range result_map {
		words_freqs = append(words_freqs, WordFreq{
			word: word,
			freq: freq,
		})
	}

	sort.Slice(words_freqs, func(i, j int) bool {
		return words_freqs[i].freq > words_freqs[j].freq
	})

	return words_freqs
}
