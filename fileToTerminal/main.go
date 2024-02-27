package main

import (
	"io"
	"log"
	"os"
)

func main() {
    argsWithoutProg := os.Args[1:]	
	filepath := argsWithoutProg[0]

	file, err := os.Open(filepath) // For read access.
	// file is of type *File
	if err != nil {
		log.Fatal(err)
	}
	// io.Copy(io.Writer,io.Reader)
	io.Copy(os.Stdout,file)
	
}