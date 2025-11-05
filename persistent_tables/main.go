package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/OmarJarbou/exercises-in-style-go/persistenttables/internal/database"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	dbQueries *database.Queries
	ctx       context.Context
}

func main() {
	config := Config{}

	if len(os.Args) < 2 {
		fmt.Println("Please provide the file path!")
		return
	}
	file_path := os.Args[1]
	fmt.Println("File path:", file_path)

	db := connectToDB("./term-frequency.db")
	defer db.Close()
	config.dbQueries = database.New(db)
	config.ctx = context.Background()

	stop_words_map := moveStopWordsToMap(addAsciiCharsToStopWords(normalizeStopWords(readStopWordsFile("../stop_words.txt"))))
	documentID := config.loadFileToDB(file_path, stop_words_map)

	words_freqs, err := config.dbQueries.CountWordsFrequencies(config.ctx, documentID)
	if err != nil {
		log.Fatal("faild to count words frequencies to db:", err)
		os.Exit(1)
	}
	for _, word_freq := range words_freqs {
		fmt.Printf("%s: %d\n", word_freq.Word, word_freq.Freq)
	}
}

func readStopWordsFile(path string) []string {
	stop_words_file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	stop_words_data, err := io.ReadAll(stop_words_file)
	if err != nil {
		log.Fatal(err)
	}

	stop_words := strings.Split(string(stop_words_data), ",")

	return stop_words
}

func normalizeStopWords(stop_words []string) []string {
	for i, stop_word := range stop_words {
		stop_words[i] = strings.ToLower(stop_word)
	}

	return stop_words
}

func addAsciiCharsToStopWords(stop_words []string) []string {
	ascii := make([]string, 0, 94)
	for i := 33; i <= 126; i++ {
		ascii = append(ascii, string(rune(i)))
	}
	stop_words = append(stop_words, ascii...)

	return stop_words
}

func moveStopWordsToMap(stop_words []string) map[string]struct{} {
	stop_words_map := map[string]struct{}{}
	for _, stop_word := range stop_words {
		stop_words_map[stop_word] = struct{}{}
	}

	return stop_words_map
}

func connectToDB(db_path string) *sql.DB {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal("faild to connect to db:", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping db:", err)
		os.Exit(1)
	}

	fmt.Println("Connected to SQLite database successfully")

	return db
}

func (config *Config) loadFileToDB(file_path string, stop_words_map map[string]struct{}) interface{} {
	path_elements := strings.Split(file_path, "/")
	file_name := path_elements[len(path_elements)-1]

	old_document, err := config.dbQueries.FindDocumentBYName(config.ctx, file_name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Document does NOT exist → proceed to add it")
			addDocumentParams := database.AddDocumentParams{
				ID:   uuid.New(),
				Name: file_name,
			}
			document, err := config.dbQueries.AddDocument(config.ctx, addDocumentParams)
			if err != nil {
				log.Fatal("failed to add document into db:", err)
				os.Exit(1)
			}

			file_data := readRealFile(file_path)

			temp_word := ""
			addWordParams := database.AddWordParams{
				ID:         uuid.New(),
				Value:      "",
				DocumentID: document.ID,
			}
			for _, byt := range file_data {
				// only consider alphabet chars in words
				if (string(byt) >= "a" && string(byt) <= "z") || (string(byt) >= "A" && string(byt) <= "Z") || (string(byt) >= "0" && string(byt) <= "9") {
					temp_word += string(byt)
					addCharacterParams := database.AddCharacterParams{
						ID:     uuid.New(),
						Value:  string(byt),
						WordID: addWordParams.ID,
					}
					_, err := config.dbQueries.AddCharacter(config.ctx, addCharacterParams)
					if err != nil {
						log.Fatal("failed to add character into db:", err)
						os.Exit(1)
					}
					continue
				}

				// if the char read is not an alphabet then save the word (only if it's not a stop word)
				addWordParams.Value = temp_word
				if addWordParams.Value == "" {
					continue
				}
				word_lower := strings.ToLower(addWordParams.Value)
				temp_word = ""

				isStopWord := false
				if _, found := stop_words_map[word_lower]; found {
					isStopWord = true
				}

				if !isStopWord {
					_, err := config.dbQueries.AddWord(config.ctx, addWordParams)
					if err != nil {
						log.Fatal("failed to add document into db:", err)
						os.Exit(1)
					}
					addWordParams.ID = uuid.New()
					addWordParams.Value = ""
				}
			}

			return document.ID
		} else {
			log.Fatal("failed to retrieve a document:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Document does exist → just return it's id")
	return old_document.ID
}

func readRealFile(path string) []byte {
	real_file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	file_data, err := io.ReadAll(real_file)
	if err != nil {
		log.Fatal(err)
	}

	return file_data
}
