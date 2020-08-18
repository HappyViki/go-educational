package main

import (
	"fmt"

	"github.com/HappyViki/go-educational/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("oG olleH"))
	fmt.Println("Give me some words, and I'll tell you a story!")
	
	fmt.Print("Adjective: ")
	var adj string
	fmt.Scanln(&adj)

	fmt.Print("Noun: ")
	var noun string
	fmt.Scanln(&noun)

	fmt.Print("Verb: ")
	var verb string
	fmt.Scanln(&verb)

	fmt.Printf("\nOnce upon a time, there was a %v %v. Every day, %v took a %v around the block. One day, %v stopped. The end.\n", adj, noun, noun, verb, noun)
}
