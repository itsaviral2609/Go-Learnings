package main

import "fmt"

// type Employee struct {
// 	mu   sync.Mutex
// 	Name string
// }

// func do(emp *Employee) {
// 	emp.mu.Lock()

// 	defer emp.mu.Unlock()
// }

func main() {
	// any struct with a Mutex must be passed by refernce!

	// value returned by range in for loop is always a copy!!!

	// for i,thing := range things{
	// 	// thing is a copy
	// 	.....
	// }

	// Use the index if you mutate the element!!

	// for i := range things {
	// 	things[i].which = xyz

	// 	// element will get mutated!
	// }

	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}} // slice of type array of 2 bytes
	a := [][]byte{}

	for _, item := range items {
		i := make([]byte, len(item))
		copy(i, item[:]) // make unique!!
		a = append(a, i)
	}
	fmt.Println(items)
	fmt.Println(a)

}
