package main

import "fmt"

// every map key must be of the same type
// every map value must be of the same type
// maps do not have dot notation

func main() {
	// this is a nil map and cannot be used
	var colors1 = make(map[string]string)

	colors2 := make(map[int]string)
	colors2[10] = "#ffffff"
	colors2[20] = "#000000"
	delete(colors2, 10)

	fmt.Println(colors1)
	fmt.Println(colors2)
	fmt.Println()

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
