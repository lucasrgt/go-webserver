package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		_, err := fmt.Fprintf(writer, "ParseForm() err: %v", err)
		if err != nil {
			return
		}
	}

	_, err := fmt.Fprintf(writer, "POST request successful\n")
	if err != nil {
		return
	}

	name := req.FormValue("name")
	address := req.FormValue("address")

	_, err = fmt.Fprintf(writer, "Name = %s\n", name)
	if err != nil {
		return
	}

	_, err = fmt.Fprintf(writer, "Address = %s\n", address)
	if err != nil {
		return
	}
}

func helloHandler(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(writer, "method is not supported", http.StatusNotFound)
		return
	}

	_, err := fmt.Fprintf(writer, "hello!")
	if err != nil {
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("././static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
