package datastorage

import (
	"log"
	"os"
)

type DataStorageManager struct {
	file_text string
	words     []string
}

func (dsm *DataStorageManager) InitializeDataStorageManager() {
	dsm.file_text = ""
	dsm.words = []string{}
}

func (dsm *DataStorageManager) ReadRealFile(path string) {
	real_file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	file_data := make([]byte, 10000000)
	_, err = real_file.Read(file_data)
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
