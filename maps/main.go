package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	printMap(colors)
	delete(colors, "green")
	printMap(colors)

	var othercolors map[string]string
	printMap(othercolors)
	
	morecolors := make(map[string]string)
	printMap(morecolors)
	morecolors["white"] = "#ffffff"
	printMap(morecolors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key + " -> " + value)
	}
}
