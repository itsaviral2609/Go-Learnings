package main

import (
	"encoding/json"
	"fmt"
)

type album1 struct {
	X int `json:"foo"`
}

type album2 struct {
	X int `json:"bar"`
}

// type Employee struct {
// 	Name   string
// 	Number int
// 	Boss   *Employee
// 	Hired  time.Time
// }

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}

func main() {
	// c := map[string]*Employee{}

	// c["Aviral"] = &Employee{ // You can remove *Employee as Employee but there is a gotcha that is  when you will do &c (pass address, will result in error)
	// 	Name:   "Aviral", // as map can rearrange internally as it is a hash table internally which is dynamic. THis address can turn to be bogus
	// 	Number: 1,
	// 	Boss:   nil,
	// 	Hired:  time.Now(),
	// }

	// c["Lamine"] = &Employee{
	// 	Name:   "Matt",
	// 	Number: 2,
	// 	Boss:   c["Aviral"],
	// 	Hired:  time.Now(),
	// }

	// c["Lamine"].Number++ // Will result in same error talked above  -----> cannot assign to struct field c["Lamine"].Number in map

	// e.Name = "Aviral"
	// e.Number = 2
	// e.Hired = time.Now()
	// fmt.Printf("%T %+[1]v\n", c["Aviral"])
	// fmt.Printf("%T %+[1]v\n", c["Lamine"])

	var album = struct {
		title  string
		artist string
		year   int
		copies int
	}{
		"Choo loo",
		"Local train",
		2016,
		1000000,
	}

	var palbum = &album
	fmt.Println(album, palbum)

	// We can't resign named struct to one another like this

	// a1 := album1{
	// 	"Make u mine",
	// }

	// a2 := album2{
	// 	"7 years",
	// }

	// //	a1 = a2 ////  ./main.go:69:7: cannot use a2 (variable of type album2) as album1 value in assignment
	// a1 = album1(a2) // this works
	// fmt.Println(a1, a2)

	// Two struct types are compatible if:
	// ---> the fields have same type and names
	// ---> in the same order
	// ---> and the same tags

	// Anonymous structs!

	v1 := struct {
		X int `json:"foo"`
	}{1}

	v2 := struct {
		X int `json:"foo"`
	}{2}

	v1 = v2
	fmt.Println(v1)

	// Types with different user-declared names are never compatible!!
	// Named struct types are convertible if they are compatible!

	// From go 1.18+ we can have type conversion even if we have tag difference

	d1 := album1{1}
	d2 := album2{2}

	d1 = album1(d2)

	fmt.Println(d1)

	// Passing structs: Structs are passed as value unless a pointer is used!!!!

	// Empty structs

	// var isEmpty map[string]struct{} // A set type instead of bool

	// done := make(chan struct{})	// a very cheap channel type

	// Empty slice has the address of empty struct. That's what it differentiates from nil slice

	// Struct tags and JSON

	r := Response{Page: 1, Words: []string{"up", "in", "out"}}
	j, _ := json.Marshal(r)

	var t2 Response

	_ = json.Unmarshal(j, &t2)

	fmt.Println(string(j))
	fmt.Printf("%#v\n", r)
	fmt.Printf("%#v\n", t2)

}
