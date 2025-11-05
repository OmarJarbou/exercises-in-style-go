package main

import (
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

type StopWordsManager struct {
	ActiveWFObjects
	stop_words map[string]struct{}
}

func (swm *StopWordsManager) run() {
	for !swm.stop {
		message := <-swm.Queue
		if message[0] == "die" {
			swm.stop = true
		}
		swm.dispatch(message)
	}
}

func (swm *StopWordsManager) dispatch(message []interface{}) {
	switch message[0].(string) {
	case "init":
		{
			swm.ReadStopWordsFile("../stop_words.txt")
		}
	case "filter":
		{
			words, ok := message[1].([]string)
			if !ok {
				log.Fatal("when StopWordsManager recieves a filter message, 2nd message argument can only be a slice of strings, which are words to be filtered")
			}

			word_freq_controller, ok := message[2].(WordFreqController)
			if !ok {
				log.Fatal("when StopWordsManager recieves a filter message, 3rd message argument can only be a WordFreqController")
			}

			word_freq_manager, ok := message[3].(WordFreqManager)
			if !ok {
				log.Fatal("when StopWordsManager recieves a filter message, 4th message argument can only be a WordFreqManager")
			}

			filtered_words := swm.RemoveStopWordsFromASlice(words)

			words_msg := []interface{}{"words", filtered_words, word_freq_controller}
			send(words_msg, word_freq_manager.Queue)
		}
	}
}

func (swm *StopWordsManager) InitializeStopWordManager() {
	swm.stop_words = map[string]struct{}{}
	swm.Queue = make(chan []interface{})
	swm.name = reflect.TypeOf(&swm).Elem().Name()
	swm.stop = false
	go swm.run()
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
