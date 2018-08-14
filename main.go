package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

func main() {
	messagePtr := flag.String("message", "", "the message to encode")
	morsePtr := flag.String("morse", "", "message in morse code to be decoded")
	letterSplitter := flag.String("delimiter", " ", "the delimiter for morse characters")
	flag.Parse()
	if *messagePtr == "" && *morsePtr == "" {
		log.Fatal(errors.New("No message to encode or decode. Provide one of -message or -morse"))
	}
	if *letterSplitter == "-" || *letterSplitter == "." {
		log.Fatal(errors.New("Letter Delimiter cannot be a \"-\" or \".\""))
	}
	if *messagePtr != "" {
		fmt.Println(encodeMessage(*messagePtr, *letterSplitter))
	}
	if *morsePtr != "" {
		fmt.Println(decodeMessage(*morsePtr, *letterSplitter))
	}
}
