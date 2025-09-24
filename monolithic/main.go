package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	word_freqs := map[string]int{}
	words := []string{}

	stop_words_file, err := os.Open("../stop_words.txt")
	if err != nil {
		log.Fatal(err)
	}

	// we user ReadAll here, but not Read, because we want it to read
	// only for the last byte in stop_words.txt, otherwise (if we used
	// regular read) if will continue reading all remaining zero bytes,
	// which will affect last stop word and will not be stored properly
	stop_words_data, err := io.ReadAll(stop_words_file)
	if err != nil {
		log.Fatal(err)
	}

	stop_words := strings.Split(string(stop_words_data), ",")
	// convert to lower
	for i, stop_word := range stop_words {
		stop_words[i] = strings.ToLower(stop_word)
	}
	// list of all printable ascii chars to be added to stop words
	ascii := make([]string, 0, 94)
	for i := 33; i <= 126; i++ {
		ascii = append(ascii, string(rune(i)))
	}
	stop_words = append(stop_words, ascii...)

	real_file, err := os.Open("../sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	file_data := make([]byte, 10000000)
	_, err = real_file.Read(file_data)
	if err != nil {
		log.Fatal(err)
	}

	// Loop over all chars/bytes in the file
	isStopWord := false
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

		for _, stop_word := range stop_words {
			if word_lower == stop_word {
				isStopWord = true
				break
			}
		}
		if !isStopWord {
			word_freqs[word]++
			if word_freqs[word] == 1 {
				words = append(words, word)
			}
		}
		isStopWord = false
	}

	// MAP IN GO IS UNORDERED; THATS WHY WE DO THIS STEP.
	// order words based on thier freqs (we don't care about order between same freqs)
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if word_freqs[words[j]] > word_freqs[words[i]] {
				words[i], words[j] = words[j], words[i]
			}
		}
	}

	// print words and thier freqs
	for i := 0; i < 25; i++ {
		fmt.Printf("%s: %d\n", words[i], word_freqs[words[i]])
	}
}
