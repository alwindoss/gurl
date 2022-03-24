package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to the gURL")
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pass the short URL key")
	})
	r.HandleFunc("/{shortKey:[a-zA-Z0-9\\d-_]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortURL := vars["shortKey"]
		fmt.Printf("Short Key: %v\n", shortURL)

		redirectURL := fetchURL(shortURL)
		http.Redirect(w, r, redirectURL, http.StatusFound)
	})
	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", r)
}

func fetchURL(shortURL string) string {
	// TODO Until this is umplemented return a valid string
	longURL := "https://alwindoss.github.io/posts/os/freebsd/how-to-setup-staticip-on-freebsd/"
	return longURL
}

func validShortURL(url string) bool {
	fmt.Println("URL: ", url)
	pathArr := strings.Split(url, "/")
	fmt.Printf("pathArr[0]: %s\n", pathArr[0])
	fmt.Printf("pathArr[1]: %s\n", pathArr[1])
	fmt.Printf("pathArr[2]: %s\n", pathArr[2])
	fmt.Println(len(pathArr))
	return false
}
