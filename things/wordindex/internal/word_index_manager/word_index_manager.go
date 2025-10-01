package wordindexmanager

import (
	"sort"
	"strings"
)

type WordIndex struct {
	Word    string
	Indexes []int
}

type WordindexManager struct {
	words_indexes []WordIndex
}

func (wim *WordindexManager) InitializeWordindexManager() {
	wim.words_indexes = []WordIndex{}
}

func (wim *WordindexManager) GetWordsIndexes() []WordIndex {
	return wim.words_indexes
}

func (wim *WordindexManager) ExtractWordsAndIndexes(lines []string, lines_per_page int) {
	words_indexes_map := map[string][]int{}

	index := 1
	line_number := 1
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			if _, ok := words_indexes_map[word]; !ok {
				words_indexes_map[word] = []int{}
			}
			words_indexes_map[word] = append(words_indexes_map[word], index)
		}
		if line_number == lines_per_page {
			index++
			line_number = 1
		}
		line_number++
	}

	for word, indexes := range words_indexes_map {
		wim.words_indexes = append(wim.words_indexes, WordIndex{
			Word:    word,
			Indexes: indexes,
		})
	}
}

func (wim *WordindexManager) SortListAlphabetically() {
	sort.Slice(wim.words_indexes, func(i, j int) bool {
		return wim.words_indexes[i].Word < wim.words_indexes[j].Word
	})
}
