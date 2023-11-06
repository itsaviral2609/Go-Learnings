package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// const url = "http://jsonplaceholder.typicode.com"

type todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>`

// Server in Go

func handler(w http.ResponseWriter, r *http.Request) {
	const base = "http://jsonplaceholder.typicode.com"

	resp, err := http.Get(base + r.URL.Path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close() //else server will gonna run out of sockets!!

	// body, err := io.ReadAll(resp.Body)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// if resp.StatusCode == http.StatusOK {
	// 	body, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(-1)
	// 	}

	var item todo
	err = json.NewDecoder(resp.Body).Decode(&item)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(-1)
	// }
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.New("mine")
	tmpl.Parse(form)
	tmpl.Execute(w, item)

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
