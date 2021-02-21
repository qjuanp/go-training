package main

import "fmt"

func main() {
	// var colors map[string]string // Declare an empty map
	// colors := make(map[string]string) // Declare an empty map

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b5869",
		"white": "#ffffff",
		"black": "#000000",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("The color %s is represented by %s\n", color, hex)
	}
}
