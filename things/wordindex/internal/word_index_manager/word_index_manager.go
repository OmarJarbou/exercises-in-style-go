package wordindexmanager

import (
	"errors"
	"sort"
	"strings"
)

type WordIndex struct {
	Word        string
	Occurrences int
	Indexes     []int
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

func (wim *WordindexManager) ExtractWordsAndIndexes(lines []string, lines_per_page int) error {
	if lines_per_page < 1 {
		return errors.New("lines per page cannot be less than 1")
	}

	words_indexes_map := map[string]map[int]struct{}{}
	words_occurences_map := map[string]int{}

	index := 1
	line_number := 1
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		words := strings.Split(trimmed, " ")
		for _, word := range words {
			if word != "" {
				if _, ok := words_indexes_map[word]; !ok {
					words_indexes_map[word] = map[int]struct{}{}
				}
				words_indexes_map[word][index] = struct{}{}
				words_occurences_map[word]++
			}
		}
		if line_number == lines_per_page {
			index++
			line_number = 0
		}
		line_number++
	}

	for word, indexes := range words_indexes_map {
		indexes_slice := []int{}
		for index := range indexes {
			indexes_slice = append(indexes_slice, index)
		}
		sort.Ints(indexes_slice)

		wim.words_indexes = append(wim.words_indexes, WordIndex{
			Word:        word,
			Occurrences: words_occurences_map[word],
			Indexes:     indexes_slice,
		})
	}

	return nil
}

func (wim *WordindexManager) SortListAlphabetically() {
	sort.Slice(wim.words_indexes, func(i, j int) bool {
		return wim.words_indexes[i].Word < wim.words_indexes[j].Word
	})
}
