package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", homePage)
	os.Setenv("PORT", "8080")
	fmt.Printf("listenting on PORT: %v", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprintf(res, "The homepage.")
}
