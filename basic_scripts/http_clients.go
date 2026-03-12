package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// HTTP GET Request
	baseURL := "https://httpbin.org"
	getResp, err := http.Get(baseURL + "/get")
	if err != nil {
		fmt.Println("Error on GET Request:", err)
		return
	}
	defer getResp.Body.Close()
	getBody, err := io.ReadAll(getResp.Body)
	if err != nil {
		fmt.Println("Error reading the body:", err)
		return
	}
	fmt.Println("GET Response:", string(getBody))
}
