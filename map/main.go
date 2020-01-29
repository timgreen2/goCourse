package main

import "fmt"

func main() {

	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	// delete(colours, "red")
	fmt.Println(colours)
	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("hex code for", colour, "is", hex)
	}
}
