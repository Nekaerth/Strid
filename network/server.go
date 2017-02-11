package main

import "net/http"

func main() {
	http.HandleFunc("/", echo)
	http.ListenAndServe(":8080", nil)
}

func echo(writer http.ResponseWriter, request *http.Request){
	writer.Write([]byte("hej du"))
}