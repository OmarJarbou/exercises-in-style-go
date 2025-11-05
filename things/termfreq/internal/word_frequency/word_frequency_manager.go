package wordfrequency

import (
	"fmt"
	"reflect"
	"sort"

	structinfo "github.com/OmarJarbou/exercises-in-style-go/things/termfreq/internal/struct_info"
)

type WordFreq struct {
	Word string
	Freq int
}

type WordFrequencyManager struct {
	words_freqs     []WordFreq
	words_freqs_map map[string]int
}

func (wfm *WordFrequencyManager) Info(i interface{}) string {
	t := reflect.TypeOf(wfm)

	// If it's a pointer, get the element type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return "In " + t.Name() + " struct"
}

func (wfm *WordFrequencyManager) InitializeWordFrequencyManager() {
	wfm.words_freqs = []WordFreq{}
	wfm.words_freqs_map = map[string]int{}
	var si structinfo.StructInfo = wfm
	fmt.Println(si.Info(si))
}

func (wfm *WordFrequencyManager) GetWordsFreqs() []WordFreq {
	return wfm.words_freqs
}

func (wfm *WordFrequencyManager) IncrementWordFreq(word string) {
	wfm.words_freqs_map[word] = wfm.words_freqs_map[word] + 1
}

func (wfm *WordFrequencyManager) CalculateWordsFreqsForASlice(words []string) {
	for _, word := range words {
		wfm.IncrementWordFreq(word)
	}
}

func (wfm *WordFrequencyManager) SortWordsBasedOnFreq() {
	for word, freq := range wfm.words_freqs_map {
		wfm.words_freqs = append(wfm.words_freqs, WordFreq{
			Word: word,
			Freq: freq,
		})
	}

	sort.Slice(wfm.words_freqs, func(i, j int) bool {
		return wfm.words_freqs[i].Freq > wfm.words_freqs[j].Freq
	})
}
