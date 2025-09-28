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
	tfQuarantine := TFQuarantine{}
	tfQuarantine.initializeTFQuarantine()
	// tfQuarantine.bind(readRealFile)
	// this won't work because bind expects the function signature to be
	// exactly func(input interface{}) interface{}, so instead we wrap
	// the function readRealFile inside that signature and we use
	// TYPE ASSERTION for arguments since interface{}(generic type) can be adjusted to any type
	tfQuarantine.bind(func(input interface{}) interface{} {
		path := input.(string)
		return readRealFile(path)() // for "callable values"/"functions" we simply call it
	})
	tfQuarantine.bind(func(input interface{}) interface{} {
		fileData := input.([]byte)
		return extractWordsAndFrequencies(fileData) // will return raw value, no need for call
	})
	tfQuarantine.bind(func(input interface{}) interface{} {
		freqs := input.(map[string]int)
		return removeStopWords(freqs)()
	})
	tfQuarantine.bind(func(input interface{}) interface{} {
		freqs := input.(map[string]int)
		return orderWordsBasedOnFreq(freqs)
	})
	tfQuarantine.bind(func(input interface{}) interface{} {
		wordsFreqs := input.([]WordFreq)
		return top25(wordsFreqs)
	})
	words_freqs := tfQuarantine.execute(getFilePath(os.Args)())

	fmt.Println(words_freqs)
}

func getFilePath(args []string) func() string {
	return func() string {
		if len(args) < 2 {
			log.Fatal("file name is required!")
			os.Exit(1)
		}
		return args[1]
	}
}

func readRealFile(path string) func() []byte {
	return func() []byte {
		real_file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		file_data, err := io.ReadAll(real_file)
		if err != nil {
			log.Fatal(err)
		}

		return file_data
	}
}

func extractWordsAndFrequencies(file_data []byte) map[string]int {
	freqs := map[string]int{}

	// Loop over all chars/bytes in the file
	temp_word := ""
	for _, byt := range file_data {
		// only consider alphabet chars in words
		if (string(byt) >= "a" && string(byt) <= "z") || (string(byt) >= "A" && string(byt) <= "Z") || (string(byt) >= "0" && string(byt) <= "9") {
			temp_word += string(byt)
			continue
		}

		// if the char read is not an alphabet then save the word (only if it's not empty)
		word := temp_word
		if word == "" {
			continue
		}
		temp_word = ""

		freqs[word]++
	}

	return freqs
}

func removeStopWords(words map[string]int) func() map[string]int {
	return func() map[string]int {
		// extract stop words from the corresponding file
		stop_words_file, err := os.Open("../stop_words.txt")
		if err != nil {
			log.Fatal(err)
		}

		stop_words_data, err := io.ReadAll(stop_words_file)
		if err != nil {
			log.Fatal(err)
		}

		stop_words := strings.Split(string(stop_words_data), ",")

		// normalize stop words
		for i, stop_word := range stop_words {
			stop_words[i] = strings.ToLower(stop_word)
		}

		// add printable ascii chars to stop words
		ascii := make([]string, 0, 94)
		for i := 33; i <= 126; i++ {
			ascii = append(ascii, string(rune(i)))
		}
		stop_words = append(stop_words, ascii...)

		// move stop words to a map for faster access
		stop_words_map := map[string]struct{}{}
		for _, stop_word := range stop_words {
			stop_words_map[stop_word] = struct{}{}
		}

		// remove stop words from file words
		for word, _ := range words {
			word_lower := strings.ToLower(word)
			if _, found := stop_words_map[word_lower]; found {
				delete(words, word)
			}
		}

		return words
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

func top25(words_freqs []WordFreq) string {
	top25_string := ""
	for i := 0; i < 25; i++ {
		top25_string += fmt.Sprintf("%s: %d\n", words_freqs[i].word, words_freqs[i].freq)
	}
	return top25_string
}
