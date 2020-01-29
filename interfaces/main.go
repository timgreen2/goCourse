package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting...
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	// VERY custom logic generating a spanish greeting...
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
