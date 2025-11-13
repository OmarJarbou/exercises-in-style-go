package main

func Map(splitWords func(string, map[string]struct{}, chan []string), chunks []string, stop_words map[string]struct{}) [][]string {
	splits := [][]string{}
	channels := [](chan []string){}
	for i, chunk := range chunks {
		ch := make(chan []string)
		channels = append(channels, ch)
		go splitWords(chunk, stop_words, channels[i])
	}
	for i, _ := range chunks {
		split := <-channels[i]
		splits = append(splits, split)
	}
	return splits
}
