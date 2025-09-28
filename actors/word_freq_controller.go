package main

import (
	"fmt"
	"log"
)

type WordFreqController struct {
	ActiveWFObjects
	data_storage_manager DataStorageManager
	stop_words_manager   StopWordsManager
	word_freq_manager    WordFreqManager
}

func (wfc *WordFreqController) run() {
	for !wfc.stop {
		message := <-wfc.Queue
		if message[0] == "die" {
			wfc.stop = true
		}
		wfc.dispatch(message)
	}
}

func (wfc *WordFreqController) dispatch(message []interface{}) {
	switch message[0].(string) {
	case "run":
		{
			data_storage_manager, ok := message[1].(DataStorageManager)
			if !ok {
				log.Fatal("when WordFreqController recieves a run message, 2nd message argument can only be a DataStorageManager")
			}
			wfc.data_storage_manager = data_storage_manager

			stop_words_manager, ok := message[2].(StopWordsManager)
			if !ok {
				log.Fatal("when WordFreqController recieves a run message, 3rd message argument can only be a StopWordsManager")
			}
			wfc.stop_words_manager = stop_words_manager

			word_freq_manager, ok := message[3].(WordFreqManager)
			if !ok {
				log.Fatal("when WordFreqController recieves a run message, 4th message argument can only be a WordFreqManager")
			}
			wfc.word_freq_manager = word_freq_manager

			send_word_freqs_msg := []interface{}{"send_word_freqs", *wfc, stop_words_manager, word_freq_manager}
			send(send_word_freqs_msg, data_storage_manager.Queue)
		}
	case "top25":
		{
			words_freqs, ok := message[1].([]WordFreq)
			if !ok {
				log.Fatal("when WordFreqController recieves a top25 message, 2nd message argument can only be a slice of WordFreq, which are words and corresponding to be printed")
			}

			wfc.printWordsAndFreqs(words_freqs)

			die_msg := []interface{}{"die"}
			send(die_msg, wfc.data_storage_manager.Queue)
			send(die_msg, wfc.stop_words_manager.Queue)
			send(die_msg, wfc.word_freq_manager.Queue)
			send(die_msg, wfc.Queue)
		}
	}
}

func (wfc *WordFreqController) initializeWordFrequencyController() {
	wfc.Queue = make(chan []interface{})
	wfc.name = "word_frequency_controller"
	wfc.stop = false
	go wfc.run()
}

func (wfc *WordFreqController) printWordsAndFreqs(words_freqs []WordFreq) {
	for i := 0; i < 25; i++ {
		fmt.Printf("%s: %d\n", words_freqs[i].word, words_freqs[i].freq)
	}
}
