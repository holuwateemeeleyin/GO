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

	//HTTP PUT Request
	client := &http.Client{}
	//create the request body
	putBody := []byte(`{"key":"updated value"}`)
	//create the pull request
	putReq, err := http.NewRequest(http.MethodPut, baseURL+"/put", bytes.NewBuffer(putBody))
	// set the request header
	putReq.Header.Set("Content-Type", "application/json")
	//Send the Request
	putResp, err := client.Do(putReq)
	if err != nil {
		fmt.Println("Error on PUT Request:", err)
		return
	}
	//close the response body
	defer putResp.Body.Close()
	//Read the response
	putRespBody, _ := io.ReadAll(putResp.Body)
	fmt.Println("PUT Response:", string(putRespBody))

	//HTTP DELETE Request
	//create the request
	delReq, err := http.NewRequest(http.MethodDelete, baseURL+"/delete", nil)
	//send the request
	delResp, err := client.Do(delReq)
	if err != nil {
		fmt.Println("Error on DELETE Request:", err)
		return
	}
	//close the connection
	defer delResp.Body.Close()
	// Read the Response
	delBodyResp, _ := io.ReadAll(delResp.Body)
	fmt.Println("DELETE Response:", string(delBodyResp))
}
