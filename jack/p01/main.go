package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type person struct {
	id   int
	name string
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: get <url>\n")
		os.Exit(1)
	}
	fmt.Printf("Arguments: \n%s\n", args[1:])
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("invalid url: %s\n", err)
	}

	response, err := http.Get(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	fmt.Printf("Http status code: %d\nHttp Body: %s", response.StatusCode, body)

	// var kv = make(map[int]string)
	// kv[1] = "jack"
	// kv[2] = "david"

	// jack := person{2, "jack"}
	// fmt.Println(kv, jack)

}
