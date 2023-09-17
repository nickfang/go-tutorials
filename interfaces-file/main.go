package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	lw := logWriter{}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file) // without interface
	io.Copy(lw, file)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	// fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
