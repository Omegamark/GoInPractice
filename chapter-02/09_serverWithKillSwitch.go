package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	os.Setenv("PORT", "8080")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "The Homepage!!")
}
