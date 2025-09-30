package main

import (
	"bufio"
	"errors"
	"os"
)

func Partition(path string, num_of_lines int) ([]string, error) {
	chunks := []string{}

	file, err := os.Open(path)
	if err != nil {
		return chunks, errors.New("error while openning the file: " + err.Error())
	}
	defer file.Close()

	chunk := ""
	i := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chunk += line + " "
		if i == num_of_lines {
			chunks = append(chunks, chunk)
			chunk = ""
			i = 1
		}
		i++
	}
	// Add last chunk if not empty
	if chunk != "" {
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}
