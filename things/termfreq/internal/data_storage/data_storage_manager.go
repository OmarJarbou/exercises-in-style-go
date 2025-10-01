package datastorage

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"

	structinfo "github.com/OmarJarbou/exercises-in-style-go/things/termfreq/internal/struct_info"
)

type DataStorageManager struct {
	structinfo.StructInfo
	file_text string
	words     []string
}

func (dsm *DataStorageManager) Info(i interface{}) string {
	t := reflect.TypeOf(i)

	// If it's a pointer, get the element type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return "In " + t.Name() + " struct"
}

func (dsm *DataStorageManager) InitializeDataStorageManager() {
	dsm.file_text = ""
	dsm.words = []string{}
	var si structinfo.StructInfo = dsm
	fmt.Println(si.Info(si))
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
