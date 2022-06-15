package main

import "fmt"

func main() {
	//different ways of decalring maps
	//example one
	colors := map[string]string {
		"red" : "#ff0000",
		"green": "#00FF00",
		"white" : "FFFFFF",
	}
	//example two - initialized without values
	// var moreColors map[string]string
	//example three - also initialized without values
	sroloc := make(map[string]string)
	//adding values to a map
	sroloc["white"] = "FFFFFF"
	//removing a value from a map
	delete(colors, "green")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

// maps vs structs: maps - keys and values must be of same type, struct can contain different types