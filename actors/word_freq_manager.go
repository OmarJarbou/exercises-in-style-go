package main

import (
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

func (wfm *WordFreqManager) dispatch(message []interface{}) {

}

func (wfm *WordFreqManager) InitializeWordFrequencyManager(ch chan []interface{}) {
	wfm.words_freqs = []WordFreq{}
	wfm.words_freqs_map = map[string]int{}
	wfm.queue = ch
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
