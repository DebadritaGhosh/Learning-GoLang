package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Goooo................")
	// PerformGetReq()
	// performPostReqJsonReq()
	PerformPostFormReq()

}

func PerformGetReq() {
	const myURL = "http://localhost:5000/get"
	response, err := http.Get(myURL)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(string(content))

	//String builder
	var responseString strings.Builder

	byteCount, _ := responseString.Write((content))

	fmt.Println("Byte count is : ", byteCount)
	fmt.Println(responseString.String())
}

func performPostReqJsonReq() {
	const myURL = "http://localhost:5000/post"

	//Fake json payload
	reqBody := strings.NewReader(`
		{
			"coursename" : "Golang",
			"price" : 900,
			"platform" :"youtube.com"
		}
	`)
	response, _ := http.Post(myURL, "application/json", reqBody)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormReq() {
	const myURL = "http://localhost:5000/postform"

	//form data
	data := url.Values{}
	data.Add("firstname", "Debadrita")
	data.Add("lastname", "Ghosh")
	data.Add("email", "debadrita.ghosh@gmail.com")

	response, err := http.PostForm(myURL, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
