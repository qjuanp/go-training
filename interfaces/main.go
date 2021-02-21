package main

import "fmt"

func main() {
	printEnGreeting(englishBot{})
	printSpGreeting(spanishBot{})
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) Greeting() string {
	return "Hi there!"
}

func (spanishBot) Greeting() string {
	return "Hola!"
}

func printEnGreeting(eb englishBot) {
	fmt.Println(eb.Greeting())
}

func printSpGreeting(sb spanishBot) {
	fmt.Println(sb.Greeting())
}
