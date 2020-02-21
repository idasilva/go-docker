package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/idasilva/go-dockerfile/helloworld"
	"log"
	"net/http"
)


func main() {

	router:= mux.NewRouter()
	router.HandleFunc("/hello",helloworld.Hello)

	fmt.Println("Server Listening")
	log.Fatal(http.ListenAndServe(":8080",router))

}