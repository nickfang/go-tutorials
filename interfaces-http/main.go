package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// using reader interface to read the response body
	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, res.Body)

	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if res.StatusCode > 299 {
	// 	fmt.Println("Error: ", res.StatusCode)
	// 	os.Exit(1)
	// }
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("%s", body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
