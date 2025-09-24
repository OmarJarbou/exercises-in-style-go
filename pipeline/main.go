package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type WordFreq struct {
	word string
	freq int
}

func main() {
	printWordsAndFreqs(orderWordsBasedOnFreq(extractWordsAndFrequenciesWithoutStopWords(readRealFile("../sample.txt"))(moveStopWordsToMap(addAsciiCharsToStopWords(normalizeStopWords(readStopWordsFile("../stop_words.txt")))))))
}

func readStopWordsFile(path string) []string {
	stop_words_file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	stop_words_data, err := io.ReadAll(stop_words_file)
	if err != nil {
		log.Fatal(err)
	}

	stop_words := strings.Split(string(stop_words_data), ",")

	return stop_words
}

func normalizeStopWords(stop_words []string) []string {
	for i, stop_word := range stop_words {
		stop_words[i] = strings.ToLower(stop_word)
	}

	return stop_words
}

func addAsciiCharsToStopWords(stop_words []string) []string {
	ascii := make([]string, 0, 94)
	for i := 33; i <= 126; i++ {
		ascii = append(ascii, string(rune(i)))
	}
	stop_words = append(stop_words, ascii...)

	return stop_words
}

func moveStopWordsToMap(stop_words []string) map[string]struct{} {
	stop_words_map := map[string]struct{}{}
	for _, stop_word := range stop_words {
		stop_words_map[stop_word] = struct{}{}
	}

	return stop_words_map
}

func readRealFile(path string) []byte {
	real_file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	file_data := make([]byte, 10000000)
	_, err = real_file.Read(file_data)
	if err != nil {
		log.Fatal(err)
	}

	return file_data
}

// Here I've done currying to transform function with two arguments to sequence
// of higher order functions each with one single argument
func extractWordsAndFrequenciesWithoutStopWords(file_data []byte) func(stop_words_map map[string]struct{}) map[string]int {
	return func(stop_words_map map[string]struct{}) map[string]int {
		freqs := map[string]int{}

		// Loop over all chars/bytes in the file
		temp_word := ""
		for _, byt := range file_data {
			// only consider alphabet chars in words
			if (string(byt) >= "a" && string(byt) <= "z") || (string(byt) >= "A" && string(byt) <= "Z") {
				temp_word += string(byt)
				continue
			}

			// if the char read is not an alphabet then save the word (only if it's not a stop word)
			word := temp_word
			if word == "" {
				continue
			}
			word_lower := strings.ToLower(word)
			temp_word = ""

			isStopWord := false
			if _, found := stop_words_map[word_lower]; found {
				isStopWord = true
			}

			if !isStopWord {
				freqs[word]++
			}
		}

		return freqs
	}
}

func orderWordsBasedOnFreq(freqs map[string]int) []WordFreq {
	words_freqs := []WordFreq{}
	for word, freq := range freqs {
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

func printWordsAndFreqs(words_freqs []WordFreq) {
	for i := 0; i < 25; i++ {
		fmt.Printf("%s: %d\n", words_freqs[i].word, words_freqs[i].freq)
	}
}
