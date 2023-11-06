package main

import (
	"fmt"
	"math"
)

// An INterface variable is nil until initialised!!!

// var r io.Reader
// var b *bytes.Buffer

// r=b

// var r io.Reader: This declares a variable r of type io.Reader. At this point, r is a nil interface.
// var b *bytes.Buffer: This declares a pointer to a bytes.Buffer which is initialized to nil.
// r=b

// Here, we're assigning the b (which is a nil pointer to a bytes.Buffer) to r. bytes.Buffer implements the io.Reader interface, so this assignment is valid.

// After the assignment, r is an interface holding a nil bytes.Buffer value. While the value inside r is nil, r itself is not considered a nil interface anymore because it now has a type (*bytes.Buffer) associated with it.

// The distinction between a nil interface and an interface with a nil value can sometimes lead to unexpected results, especially when checking for nil values.

// Interface has 2 parts!

// -----> A value or pointer of some type
// -----> A pointer to type information so that the correct actual method can be identified!
/*
type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
	return fmt.Sprintf("%s PATH | %s err", e.path, e.err)
}

func XYZ(a int) error {
	return nil
}

func main() {
	var err error = XYZ(1)
	if err != nil { // When u place a nil concrete pointer to an interfacae, the interface is non nil
		fmt.Println("Print oops!!")
	} else {
		fmt.Println("OKKKK")
	}

}

// POINTER VS VALUE RECEIVER

type Coordinates struct {
	x, y float32
}

func (c *Coordinates) Add(x, y float32) {
	c.x, c.y = c.x+x, c.y+y
}

func (c Coordinates) AddOffSet(p Coordinates) (x float32, y float32) {
	x, y = c.x-p.x, c.y-p.y
	return
}

// The same method may not be bound to T and T*

// GO will automatically use * or & as needed!!!

p1 := new(Coordinates)
p2 := Coordinates{2,2}

p1.AddOffSet(p2) // same as (*p1).AddOffSet(p2)
p2.Add(3,4)  // same as (&p2).Add(3,4)

/*
                   |        POINTER           |      L-VALUE       |    R-VALUE
                   |                          |					   |
------------------------------------------------------------------------
POINTER RECEIVER   |                          |                    |
                   |              OK          |   OK &             | NOT OK
                   |                          |                    |
------------------------------------------------------------------------
                   |                          |                    |
VALUE RECEIVER	   |               OK*        |    OK              |  OK
				   |                          |                    |



var p Coordinates

p.Add(1,2)
Coordinates{1,2}.Add(2,3)  // NOT OK CAN'T TAKE ADDRESS!! Coordinates{1,2} (R) literal like thus will give a value

// If one method of a type takes a pointer receiver, then all its methods should take pointers*
// In general Objects of such types are not safe to copy

type Buffer struct{
	buf []byte
	off int
}

func (b *Buffer)ReadString(delim byte)(string,error){

}


// CURRYING FUNCTION


func Add(a,b int)int{
	return a+b
}

func AddToA(a int) func(int)int{
	return func(b int)int{
		return Add(a,b)
	}
}

AddTo1 := AddToA(1)
fmt.Println(Add(1,2)== AddTo1(2))  // True

// Currying takes a function and reduces its argument by 1  (ONe argument gets bound, and new function is returned!)
*/
// METHOD VALUES!

type Point struct {
	x, y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {
	p := Point{1, 1}
	q := Point{5, 4}
	//fmt.Println(p.Distance(q))
	DistanceFromP := p.Distance // CLosing over the value
	fmt.Printf("%T", DistanceFromP)
	fmt.Println(DistanceFromP(q))
}

// Empty Interfaces are empty. They are satisfied by anyone!
