package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "pong")
	})

	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "pong")
	})

	http.HandleFunc("hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello")
	})

	err := http.ListenAndServe(":9990", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	math.M()
}
