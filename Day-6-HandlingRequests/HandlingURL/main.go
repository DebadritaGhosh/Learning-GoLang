package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=kdj1235d"

func main() {
	fmt.Println("URL.........")

	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)   // Scheme = https
	fmt.Println(result.Host)     // Scheme = lco.dev:3000
	fmt.Println(result.Path)     // Scheme = /learn
	fmt.Println(result.Port())   // Scheme = 3000
	fmt.Println(result.RawQuery) // Scheme = coursename=reactjs&paymentid=kdj1235d

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n", qparams)

	fmt.Println(qparams["coursename"]) //[reactjs]
	fmt.Println(qparams["paymentid"])  //[kdj1235d]

	//Creating another URL

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=debadrita",
	}

	anotherURL := partsOfUrl.String()

	fmt.Println(anotherURL)
}
