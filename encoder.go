package main

import "strings"

func encodeMessage(message, letterSplitter string) string {
	var output string
	message = strings.ToUpper(message)
	words := strings.Split(message, " ")
	for _, word := range words {
		word := encodeWord(word, letterSplitter)
		if word != "" {
			output += word
		}
	}
	return output
}

func encodeWord(word, letterSplitter string) string {
	var morse string
	for i := 0; i < len(word); i++ {
		code := morsemap[word[i:i+1]]
		if code != "" {
			morse += code + letterSplitter
		}
	}
	return morse
}

func decodeLetter(code string) string {
	return characterMap[code]
}

func decodeMessage(message, letterSplitter string) string {
	var decodedMessage string
	var letters = strings.Split(message, letterSplitter)
	for _, letter := range letters {
		decodedMessage += decodeLetter(strings.Trim(letter, " ")) + " "
	}
	return decodedMessage
}
