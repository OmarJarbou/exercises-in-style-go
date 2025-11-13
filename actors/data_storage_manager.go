package main

import (
	"io"
	"log"
	"os"
)

type DataStorageManager struct {
	ActiveWFObjects
	file_text string
	words     []string
}

func (dsm *DataStorageManager) run() {
	for !dsm.stop {
		message := <-dsm.Queue
		if message[0] == "die" {
			dsm.stop = true
		}
		dsm.dispatch(message)
	}
}

func (dsm *DataStorageManager) dispatch(message []interface{}) {
	switch message[0].(string) {
	case "init":
		{
			file_path, ok := message[1].(string)
			if !ok {
				log.Fatal("in data storage manager initialization, 2nd message argument can only be string, which is the file path")
			}
			dsm.ReadRealFile(file_path)
		}
	case "send_word_freqs":
		{
			word_freq_controller, ok := message[1].(WordFreqController)
			if !ok {
				log.Fatal("when DataStorageManager recieves a send_word_freqs message, 2nd message argument can only be a WordFreqController")
			}

			stop_words_manager, ok := message[2].(StopWordsManager)
			if !ok {
				log.Fatal("when DataStorageManager recieves a send_word_freqs message, 3rd message argument can only be a StopWordsManager")
			}

			word_freq_manager, ok := message[3].(WordFreqManager)
			if !ok {
				log.Fatal("when DataStorageManager recieves a send_word_freqs message, 4th message argument can only be a WordFreqManager")
			}

			dsm.ExtractWords()

			filter_msg := []interface{}{"filter", dsm.GetWords(), word_freq_controller, word_freq_manager}
			send(filter_msg, stop_words_manager.Queue)
		}
	}
}

func (dsm *DataStorageManager) InitializeDataStorageManager() {
	dsm.file_text = ""
	dsm.words = []string{}
	dsm.Queue = make(chan []interface{})
	dsm.name = "data_storage_manager"
	dsm.stop = false
	go dsm.run()
}

func (dsm *DataStorageManager) ReadRealFile(path string) {
	real_file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	file_data, err := io.ReadAll(real_file)
	if err != nil {
		log.Fatal(err)
	}
	dsm.file_text = string(file_data)
}

func (dsm *DataStorageManager) ExtractWords() {
	temp_word := ""
	for _, char := range dsm.file_text {
		// only consider alphanumaric chars in words
		if (string(char) >= "a" && string(char) <= "z") || (string(char) >= "A" && string(char) <= "Z") || (string(char) >= "0" && string(char) <= "9") {
			temp_word += string(char)
			continue
		}
		// if the char read is not an alphabet then save the word
		if temp_word == "" {
			continue
		}
		dsm.words = append(dsm.words, temp_word)
		temp_word = ""
	}
}

func (dsm *DataStorageManager) GetWords() []string {
	return dsm.words
}
