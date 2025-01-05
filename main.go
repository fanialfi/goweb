package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	message := "welcome"
	w.Write([]byte(message))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	message := "hello world"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/index", handleIndex)
	http.HandleFunc("/hello", handleHello)

	address := "localhost:3000"
	log.Printf("server started at %s\n", address)
	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error :", err.Error())
	}
}
