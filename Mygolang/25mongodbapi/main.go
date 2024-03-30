package main

import (
	"fmt"
	"log"
	"mongodbapi/router"
	"net/http"
)

func main() {
	fmt.Println("mangodbapi")
	r := router.Router()
	fmt.Println("serevr is started....")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("listening")

}
