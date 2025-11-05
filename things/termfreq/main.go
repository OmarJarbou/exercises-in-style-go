package main

func main() {
	word_freq_controller := WordFrequencyContoller{}
	word_freq_controller.initializeWordFrequencyController()
	word_freq_controller.run("../../stop_words.txt", "../../sample.txt")
}
