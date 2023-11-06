package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s []int
	t := []int{}
	u := make([]int, 5)
	v := make([]int, 0, 5)

	fmt.Printf("%d, %d, %T,%t,%#[3]v\n", len(s), cap(s), s, s == nil) // nil slice      0, 0, []int,true,[]int(nil)
	fmt.Printf("%d, %d, %T,%t,%#[3]v\n", len(t), cap(t), t, t == nil) // empty slice    0, 0, []int,false,[]int{}
	fmt.Printf("%d, %d, %T,%t,%#[3]v\n", len(u), cap(u), u, u == nil) // empty slice    5, 5, []int,false,[]int{0, 0, 0, 0, 0}
	fmt.Printf("%d, %d, %T,%t,%#[3]v\n", len(v), cap(v), v, v == nil) //               0, 5, []int,false,[]int{}

	// Slices (Maps too!) encoding differently in JSON when nil!

	var a []int
	b := []int{}

	j, _ := json.Marshal(a)
	fmt.Println(string(j)) // null

	k, _ := json.Marshal(b)
	fmt.Println(string(k)) // []

	// define length of slice as 0 with a capacity defined so when u will append any element, u will get the element at first place!

	// It is perfectly ok to append to a nil slice!! // var s []int -------> we can append in this nil slice

	// Appending to a nil slice != inserting into a nil map!!!!

	// Length vs capacity in slice!

	x := [...]int{1, 2, 3}
	y := x[0:1]
	z := y[:2]

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

	fmt.Println("capacity of z: ", cap(z))
	fmt.Println("capacity of z: ", len(z))

	d := x[0:1:1] //  [i:j:k] len: i-j cap: j-k

	fmt.Println("len of d: ", len(d))
	fmt.Println("cap of d: ", cap(d))

	// If u use 2 index slice operator [:], capacity of slice is equal to the capacity of underlying array!!

	fmt.Printf("x[%p] = %v\n", &x, x)
	fmt.Printf("y[%p] = %[1]v\n", y)
	fmt.Printf("z[%p] = %[1]v\n", z) // There is one more space in z as its capacity is 3 so 6 is getting appended in the same memory location array(underlying)
	z = append(z, 6)
	fmt.Printf("z[%p] = %[1]v\n", z)

}
