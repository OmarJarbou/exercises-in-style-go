package main

import (
	"fmt"

	datastorage "github.com/OmarJarbou/exercises-in-style-go/things/internal/data_storage"
	stopwords "github.com/OmarJarbou/exercises-in-style-go/things/internal/stop_words"
	structinfo "github.com/OmarJarbou/exercises-in-style-go/things/internal/struct_info"
	wordfrequency "github.com/OmarJarbou/exercises-in-style-go/things/internal/word_frequency"
)

type WordFrequencyContoller struct {
	structinfo.StructInfo
	data_storage_manager datastorage.DataStorageManager
	stop_words_manager   stopwords.StopWordsManager
	word_freq_manager    wordfrequency.WordFrequencyManager
}

func (wfc *WordFrequencyContoller) initializeWordFrequencyController() {
	wfc.data_storage_manager = datastorage.DataStorageManager{}
	wfc.stop_words_manager = stopwords.StopWordsManager{}
	wfc.word_freq_manager = wordfrequency.WordFrequencyManager{}

	wfc.data_storage_manager.InitializeDataStorageManager()
	wfc.stop_words_manager.InitializeStopWordManager()
	wfc.word_freq_manager.InitializeWordFrequencyManager()

	var si structinfo.StructInfo = wfc
	fmt.Println(structinfo.Info(si))
}

func (wfc *WordFrequencyContoller) run(stop_words_path, real_file_path string) {
	wfc.data_storage_manager.ReadRealFile(real_file_path)
	wfc.data_storage_manager.ExtractWords()

	wfc.stop_words_manager.ReadStopWordsFile(stop_words_path)
	wfc.stop_words_manager.NormalizeStopWords()
	wfc.stop_words_manager.AddAsciiCharsToStopWords()
	wfc.stop_words_manager.FillStopWordsMap()
	filtered_words := wfc.stop_words_manager.RemoveStopWordsFromASlice(wfc.data_storage_manager.GetWords())

	wfc.word_freq_manager.CalculateWordsFreqsForASlice(filtered_words)
	wfc.word_freq_manager.SortWordsBasedOnFreq()

	wfc.printWordsAndFreqs(wfc.word_freq_manager.GetWordsFreqs())
}

func (wfc *WordFrequencyContoller) printWordsAndFreqs(words_freqs []wordfrequency.WordFreq) {
	for i := 0; i < 25; i++ {
		fmt.Printf("%s: %d\n", words_freqs[i].Word, words_freqs[i].Freq)
	}
}
