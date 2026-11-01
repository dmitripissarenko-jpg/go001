package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(BarServer)
	log.Fatal(http.ListenAndServe(":9000", handler))
}
