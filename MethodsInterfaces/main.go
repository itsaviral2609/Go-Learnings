package main

import (
	"fmt"
	"math"
)

// type Intslice []int

// func (k Intslice) String() string {
// 	var strs []string

// 	for _, v := range k {
// 		strs = append(strs, strconv.Itoa(v))
// 	}

// 	return "[" + strings.Join(strs, ";") + "]"
// }

// func main() {
// 	var v Intslice = []int{1, 2, 3, 4, 5}
// 	var s fmt.Stringer = v
// 	for i, x := range v {
// 		fmt.Printf("%d: , %d\n", i, x)
// 	}
// 	fmt.Printf("%T,%[1]v\n", v)
// 	fmt.Printf("%T,%[1]v\n", s)
// }

// type ByteCounter int

// func (b *ByteCounter) Write(p []byte) (int, error) {
// 	*b += ByteCounter(len(p)) //
// 	return len(p), nil
// }

// func main() {
// 	var c ByteCounter
// 	f1, _ := os.Open("a.txt")
// 	f2 := &c
// 	n, _ := io.Copy(f2, f1)
// 	fmt.Printf("Copied: %d, bytes\n", n)
// 	fmt.Println(c)

// }

// // io.ReadWriter interface

// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// // Combined above two

// type ReadWriter interface {
// 	Reader
// 	Writer
// }

type Point struct {
	X, Y float64
}

// type Path []Point

type Line struct {
	Begin, End Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.Begin.Y-l.End.Y)

}

// func (p Path) Distance() (sum float64) {
// 	for i := 1; i < len(p); i++ {
// 		sum += Line{p[i-1], p[i]}.Distance()
// 	}
// 	return sum
// }

// type Distancer interface {
// 	Distance() float64
// }

// func PrintDistance(d Distancer) {
// 	fmt.Println(d.Distance())
// }

func (l Line) ScaleBy(f float64) Line {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)
	return Line{l.Begin, Point{l.End.X, l.Begin.Y}}
}

func main() {
	side := Line{Point{1, 2}, Point{4, 6}}
	// perimeter := Path{{1, 1}, {5, 1}, {3, 5}, {1, 1}}
	s2 := side.ScaleBy(2)
	fmt.Println(s2.Distance())
	fmt.Println(Line{Point{1, 4}, Point{5, 6}}.ScaleBy(2).Distance())
	// PrintDistance(perimeter)
}

// Example to Demonstrate interface

// type Chargeable interface {
// 	charge()
// }

// func DoCharge(c Chargeable) {
// 	fmt.Println("Charging the given specimen")
// 	c.charge()
// }

// type Laptop struct{}

// func (s Laptop) charge() {
// 	fmt.Println("Laptop is getting charged!!!")
// }

// type Phone struct{}

// func (k Phone) charge() {
// 	fmt.Println("Phone is getting charged!!!")
// }

// type Ipad struct{}

// func (z Ipad) charge() {
// 	fmt.Println("Ipad is getting charged!!!")
// }

// func main() {
// 	laptop := Laptop{}
// 	DoCharge(laptop)

// 	phone := Phone{}
// 	DoCharge(phone)

// 	ipad := Ipad{}
// 	DoCharge(ipad)
// }
