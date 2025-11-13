package main

import (
	"log"
	"reflect"
	"sort"
)

type WordFreq struct {
	word string
	freq int
}

type WordFreqManager struct {
	ActiveWFObjects
	words_freqs     []WordFreq
	words_freqs_map map[string]int
}

func (wfm *WordFreqManager) run() {
	for !wfm.stop {
		message := <-wfm.Queue
		if message[0] == "die" {
			wfm.stop = true
		}
		wfm.dispatch(message)
	}
}

func (wfm *WordFreqManager) dispatch(message []interface{}) {
	switch message[0].(string) {
	case "words":
		{
			words, ok := message[1].([]string)
			if !ok {
				log.Fatal("when WordFreqManager recieves a filter message, 2nd message argument can only be a slice of strings, which are words to be counted and sorted")
			}

			word_freq_controller, ok := message[2].(WordFreqController)
			if !ok {
				log.Fatal("when WordFreqManager recieves a filter message, 3rd message argument can only be a WordFreqController")
			}

			wfm.CalculateWordsFreqsForASlice(words)
			wfm.SortWordsBasedOnFreq()

			top25_msg := []interface{}{"top25", wfm.words_freqs}
			send(top25_msg, word_freq_controller.Queue)
		}
	}
}

func (wfm *WordFreqManager) InitializeWordFrequencyManager() {
	wfm.words_freqs = []WordFreq{}
	wfm.words_freqs_map = map[string]int{}
	wfm.Queue = make(chan []interface{})
	wfm.name = reflect.TypeOf(&wfm).Elem().Name()
	wfm.stop = false
	go wfm.run()
}

func (wfm *WordFreqManager) GetWordsFreqs() []WordFreq {
	return wfm.words_freqs
}

func (wfm *WordFreqManager) IncrementWordFreq(word string) {
	wfm.words_freqs_map[word] = wfm.words_freqs_map[word] + 1
}

func (wfm *WordFreqManager) CalculateWordsFreqsForASlice(words []string) {
	for _, word := range words {
		wfm.IncrementWordFreq(word)
	}
}

func (wfm *WordFreqManager) SortWordsBasedOnFreq() {
	for word, freq := range wfm.words_freqs_map {
		wfm.words_freqs = append(wfm.words_freqs, WordFreq{
			word: word,
			freq: freq,
		})
	}

	sort.Slice(wfm.words_freqs, func(i, j int) bool {
		return wfm.words_freqs[i].freq > wfm.words_freqs[j].freq
	})
}
