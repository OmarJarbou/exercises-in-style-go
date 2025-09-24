package stopwords

import (
	"io"
	"log"
	"os"
	"strings"
)

type StopWordsManager struct {
	stop_words     []string
	stop_words_map map[string]struct{}
}

func (swm *StopWordsManager) InitializeStopWordManager() {
	swm.stop_words = []string{}
	swm.stop_words_map = map[string]struct{}{}
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
	swm.stop_words = strings.Split(string(stop_words_data), ",")
}

func (swm *StopWordsManager) NormalizeStopWords() {
	for i, stop_word := range swm.stop_words {
		swm.stop_words[i] = strings.ToLower(stop_word)
	}
}

func (swm *StopWordsManager) AddAsciiCharsToStopWords() {
	ascii_chars := make([]string, 0, 94)
	for i := 33; i <= 126; i++ {
		ascii_chars = append(ascii_chars, string(rune(i)))
	}
	swm.stop_words = append(swm.stop_words, ascii_chars...)
}

func (swm *StopWordsManager) FillStopWordsMap() {
	for _, stop_word := range swm.stop_words {
		swm.stop_words_map[stop_word] = struct{}{}
	}
}
func (swm *StopWordsManager) IsStopWord(word string) bool {
	_, found := swm.stop_words_map[word]
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
