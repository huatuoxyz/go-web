package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Profile ...
type Profile struct {
	Name    string
	Hobbies []string
}

const links = `
  <ul>
  <li><a href='/hello/dongri'>hello</a></li>
  <li><a href='/html'>html</a></li>
  <li><a href='/json'>json</a></li>
  </ul>`

func main() {
	http.HandleFunc("/", homeViewHandler)
	http.HandleFunc("/hello/", helloViewHandler)
	http.HandleFunc("/html", htmlViewHandler)
	http.HandleFunc("/json", jsonViewHandler)
	http.ListenAndServe(":8080", nil)
}

func homeViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, links)
}

func helloViewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path[7:])
}

func htmlViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<img src='http://golang.org/doc/gopher/appenginegophercolor.jpg' width=100>")
}

func jsonViewHandler(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Dongri Jin", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/json")
	w.Write(js)
}
