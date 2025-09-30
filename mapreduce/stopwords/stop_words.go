package stopwords

import (
	"errors"
	"io"
	"os"
	"strings"
)

type StopWordsManager struct {
	stop_words     []string
	stop_words_map map[string]struct{}
}

func (swm *StopWordsManager) GetStopWords() (map[string]struct{}, error) {
	swm.stop_words = []string{}
	swm.stop_words_map = map[string]struct{}{}
	err := swm.readStopWords()
	if err != nil {
		return swm.stop_words_map, err
	}
	swm.normalizeStopWords()
	swm.addAsciiCharsToStopWords()
	swm.moveStopWordsToMap()

	return swm.stop_words_map, err
}

func (swm *StopWordsManager) readStopWords() error {
	file, err := os.Open("../stop_words.txt")
	if err != nil {
		return errors.New("error while openning stop file: " + err.Error())
	}
	defer file.Close()

	var stop_words_data []byte
	stop_words_data, err = io.ReadAll(file)
	if err != nil {
		return errors.New("error while reading stop file: " + err.Error())
	}

	swm.stop_words = strings.Split(string(stop_words_data), ",")
	return nil
}

func (swm *StopWordsManager) normalizeStopWords() {
	for i, stop_word := range swm.stop_words {
		swm.stop_words[i] = strings.ToLower(stop_word)
	}
}

func (swm *StopWordsManager) addAsciiCharsToStopWords() {
	ascii := make([]string, 0, 94)
	for i := 33; i <= 126; i++ {
		ascii = append(ascii, string(rune(i)))
	}
	swm.stop_words = append(swm.stop_words, ascii...)
}

func (swm *StopWordsManager) moveStopWordsToMap() {
	for _, stop_word := range swm.stop_words {
		swm.stop_words_map[stop_word] = struct{}{}
	}
}
