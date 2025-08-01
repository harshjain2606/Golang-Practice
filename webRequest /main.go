package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		fmt.Println("error in getting file", err)
		return
	}
	defer res.Body.Close()

	input, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error in reading input", err)
		return
	}
	fmt.Println(string(input))
}
