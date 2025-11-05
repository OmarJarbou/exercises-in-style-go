package main

func countWords(words []string, ch chan map[string]int) {
	words_freqs_map := map[string]int{}
	for _, word := range words {
		words_freqs_map[word]++
	}

	ch <- words_freqs_map
}
