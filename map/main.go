package main

import "fmt"

func main() {
	// var colors map[string]string // Declare an empty map
	// colors := make(map[string]string) // Declare an empty map

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b5869",
	}

	colors["white"] = "#000000"

	delete(colors, "white")
	fmt.Println(colors)
}
