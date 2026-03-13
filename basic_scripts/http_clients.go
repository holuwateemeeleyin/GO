package main

import (
	"bytes"
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

	// Http POST Request:
	//create the body
	postBody := []byte(`{"key":"value"}`)
	//send the body
	postResp, err := http.Post(baseURL+"/post", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("Error on POST Request:", err)
	}
	//Close the connection
	defer postResp.Body.Close()
	//Read the response
	postRespBody, _ := io.ReadAll(postResp.Body)
	fmt.Println("POST Response:", string(postRespBody))
}
