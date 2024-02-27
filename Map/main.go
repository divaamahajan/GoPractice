package main

import "fmt"

// ------------------Map=> Dictionary -->
func main() {
	colors := map[string]string{
		"red":  "#ff0000",
		"blue": "#0000ff",
        "green": "#00ff00",
		"yellow": "#ffff00",
		"white": "#ffffff",
	}
	// printMap(colors)
	colors["black"] = "#000000"
	delete(colors, "green")
	printMap(colors)
	fmt.Println(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		m[k] = "hello,"
        fmt.Println(k, v, m[k])
    }
}

