package datastorage

import (
	"bufio"
	"errors"
	"os"
)

type DataStorageManager struct {
	lines []string
}

func (dsm *DataStorageManager) InitializeDataStorageManager() {
	dsm.lines = []string{}
}

func (dsm *DataStorageManager) GetLines() []string {
	return dsm.lines
}

func (dsm *DataStorageManager) ReadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.New("Error while opening file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dsm.lines = append(dsm.lines, line)
	}

	return nil
}

func (dsm *DataStorageManager) FilterLinesCharacters() {
	for i, line := range dsm.lines {
		filtered_line := ""
		for _, char := range line {
			if (string(char) >= "a" && string(char) <= "z") || (string(char) >= "A" && string(char) <= "Z") {
				filtered_line += string(char)
			} else {
				if len(filtered_line) > 0 {
					if string(filtered_line[len(filtered_line)-1]) != " " {
						filtered_line += " "
					}
				}
			}
		}
		dsm.lines[i] = filtered_line
	}
}
