package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	CreatePost()
	// GetPostByID("101")
}

func GetPostByID(ID string) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + ID)
	if err != nil {
		fmt.Println("Failed to get post by ID: ", err)
		return
	}

	defer res.Body.Close()
	fmt.Println(res.Body)

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body: ", err)
		return
	}

	fmt.Println(string(bodyByte))
}

func CreatePost() {
	data := map[string]interface{}{
		"title":  "Test",
		"body":   "test test test test",
		"userId": 100,
	}

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Failed to marshal request data: ", err)
		return
	}

	res, err := http.Post("https://jsonplaceholder.typicode.com/posts", "Application/json", bytes.NewBuffer(requestJSON))
	defer res.Body.Close()
	fmt.Println(res.Body)

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body: ", err)
		return
	}

	fmt.Println(string(bodyByte))

	// req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJSON))
	// if err != nil {
	// 	fmt.Println("Failed to initialize create post client: ", err)
	// 	return
	// }

	// req.Header.Set("Content-type", "Application/json")

	// client := http.Client{}

	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Failed to create post: ", err)
	// 	return
	// }

	// defer res.Body.Close()
	// fmt.Println(res.Body)

	// bodyByte, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Failed to read response body: ", err)
	// 	return
	// }

	// fmt.Println(string(bodyByte))
}
