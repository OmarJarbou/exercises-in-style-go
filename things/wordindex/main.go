package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	word_index_controller := WordIndexController{}
	word_index_controller.InitializeWordIndexController()

	if len(os.Args) < 3 {
		log.Fatal("Three argument required: go run 'path_to_project' 'path_to_file' 'lines_per_page'")
		return
	}
	lines_per_page, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Error Converting lines per page:", err)
		return
	}

	err = word_index_controller.run(os.Args[1], lines_per_page)
	if err != nil {
		log.Fatal(err)
		return
	}
}
