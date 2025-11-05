package main

import (
	"fmt"
	"log"
	"os"

	"github.com/OmarJarbou/exercises-in-style-go/mapreduce/stopwords"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("you should enter the name of the file")
	}

	chunks, err := Partition(os.Args[1], 200)
	if err != nil {
		log.Fatal(err.Error())
	}

	stop_words_manager := stopwords.StopWordsManager{}
	stop_words, err := stop_words_manager.GetStopWords()
	if err != nil {
		log.Fatal(err.Error())
	}

	splits := Map(splitWords, chunks, stop_words)
	words_freqs := Reduce(countWords, splits)

	for i := 0; i < 25; i++ {
		fmt.Printf("%s: %d\n", words_freqs[i].word, words_freqs[i].freq)
	}
}
