package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("using gorilla mux")
	sample()
	r := mux.NewRouter()
	r.HandleFunc("/", httphome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func sample() {
	fmt.Println("gorilla users")
}

func httphome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to pavan's golang</h1>"))
}
