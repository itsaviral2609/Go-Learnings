package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// NOTE: don't do this in real life
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// Add the handlers!

func (db database) List(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("Invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "Added %s with price %s \n", item, db[item])

}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("No such item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("Invalid price: %q", price)
		http.Error(w, msg, http.StatusNotFound) // 404
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "New Price %s for price %s \n", db[item], item)

}

func (db database) fetch(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("No such item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	fmt.Fprintf(w, "item %s has price %s \n", item, db[item])

}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("No such item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "item %s has been deleted whose price %s \n", item, db[item])

}

func main() {
	db := database{
		"shoes": 50,
		"socks": 20,
	}

	// Adding some routes
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/delete", db.delete)

	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/read", db.fetch)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
