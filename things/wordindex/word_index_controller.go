package main

import (
	"fmt"
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

	err = wic.word_index_manager.ExtractWordsAndIndexes(lines, lines_per_page)
	if err != nil {
		return err
	}
	wic.word_index_manager.SortListAlphabetically()
	words_indexes := wic.word_index_manager.GetWordsIndexes()

	for _, w := range words_indexes {
		if w.Occurrences < 100 {
			fmt.Println(w.Word, "(#Occurrences: "+strconv.Itoa(w.Occurrences)+") - ", w.Indexes)
		}
	}

	return nil
}
