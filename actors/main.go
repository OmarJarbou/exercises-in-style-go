package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("file path is required!")
		os.Exit(1)
	}
	data_storage_manager := DataStorageManager{}
	stop_words_manager := StopWordsManager{}
	word_freq_manager := WordFreqManager{}
	word_freq_controller := WordFreqController{}

	data_storage_manager.InitializeDataStorageManager()
	stop_words_manager.InitializeStopWordManager()
	word_freq_manager.InitializeWordFrequencyManager()
	word_freq_controller.initializeWordFrequencyController()

	init_msg1 := []interface{}{"init"}
	send(init_msg1, stop_words_manager.Queue)

	init_msg2 := []interface{}{"init", os.Args[1]}
	send(init_msg2, data_storage_manager.Queue)

	run_msg := []interface{}{"run", data_storage_manager, stop_words_manager, word_freq_manager}
	send(run_msg, word_freq_controller.Queue)

	for !data_storage_manager.stop && !stop_words_manager.stop && !word_freq_manager.stop && !word_freq_controller.stop {
		// busy wait
	}
}

func send(message []interface{}, queue chan []interface{}) {
	_, ok := message[0].(string)
	if !ok {
		log.Fatal("message tag can only be string")
	}
	queue <- message
}
