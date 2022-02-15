package main

import (
	"APIwithDb/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Mongodb api......")
	fmt.Print("Server is getting started...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}
