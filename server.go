package main

import (
	"fmt"
	"net/http"
)

func BarServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bar")
}
