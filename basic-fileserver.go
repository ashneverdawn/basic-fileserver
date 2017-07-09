package main

import (
	"net/http"
)

func main() {
    port := "8080"

	http.Handle("/", http.FileServer(http.Dir("public")))

    println("Listening on port " + port)
    err := http.ListenAndServe(":"+port, nil)
	if err != nil { println(err) }
}