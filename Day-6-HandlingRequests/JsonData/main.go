package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json......")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"ReactJS Bootcamp", 299, "Youtube.com", "abc123", []string{"facebook", "js"}},
		{"Go Bootcamp", 199, "Youtube.com", "def123", []string{"backedn-dev", "GoLang"}},
		{"Angular Bootcamp", 365, "Youtube.com", "ghi123", nil},
	}

	//Package this data as JSON data
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename":"ReactJS Bootcamo",
		"Price":299,
		"website": "youtube.com",
		"tags" : ["web-dev","js"]
	}
	`)

	var courses course

	checkvalid := json.Valid(jsonDataFromWeb)
	if checkvalid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//Some cases where we just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", courses)

}
