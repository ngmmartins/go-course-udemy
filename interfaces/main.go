package main

import "fmt"

type Bot interface {
	GetGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	PrintGreeting(eb)
	PrintGreeting(sb)
}

func PrintGreeting(b Bot) {
	fmt.Println(b.GetGreeting())
}

func (EnglishBot) GetGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (SpanishBot) GetGreeting() string {
	return "Hola!"
}
