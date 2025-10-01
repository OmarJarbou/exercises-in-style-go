package main

import (
	"fmt"
	"sort"
	"strconv"

	datastorage "github.com/OmarJarbou/exercises-in-style-go/things/wordindex/internal/data_storage"
	wordindexmanager "github.com/OmarJarbou/exercises-in-style-go/things/wordindex/internal/word_index_manager"
)

type WordIndexController struct {
	data_storage_manager datastorage.DataStorageManager
	word_index_manager   wordindexmanager.WordindexManager
}

func (wic *WordIndexController) InitializeWordIndexController() {
	wic.data_storage_manager = datastorage.DataStorageManager{}
	wic.word_index_manager = wordindexmanager.WordindexManager{}

	wic.data_storage_manager.InitializeDataStorageManager()
	wic.word_index_manager.InitializeWordindexManager()
}

func (wic *WordIndexController) run(file_path string, lines_per_page int) error {
	err := wic.data_storage_manager.ReadFile(file_path)
	if err != nil {
		return err
	}
	wic.data_storage_manager.FilterLinesCharacters()
	lines := wic.data_storage_manager.GetLines()

	wic.word_index_manager.ExtractWordsAndIndexes(lines, lines_per_page)
	wic.word_index_manager.SortListAlphabetically()
	words_indexes := wic.word_index_manager.GetWordsIndexes()

	for _, w := range words_indexes {
		if len(w.Indexes) < 100 {
			// Use a map as a set to deduplicate
			unique := make(map[int]struct{})
			for _, index := range w.Indexes {
				unique[index] = struct{}{}
			}

			// Convert map keys back into a slice for sorting
			deduped := make([]int, 0, len(unique))
			for index := range unique {
				deduped = append(deduped, index)
			}

			sort.Ints(deduped)

			fmt.Println(w.Word, "(#Occurrences: "+strconv.Itoa(len(w.Indexes))+") - ", deduped)
		}
	}

	return nil
}
