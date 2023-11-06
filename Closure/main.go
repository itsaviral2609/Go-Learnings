// Application of Closure!

package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func do(d func()) {
	d()
}

func main() {

	//f := fib()

	f, g := fib(), fib()
	// for x := f(); x < 100; x = f() {
	// 	fmt.Println(x)
	// }

	fmt.Println(f(), f(), f(), f(), f())
	fmt.Println(g(), g(), g(), g(), g())

	for i := 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
		do(v)
	}

	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		i2 := i // capturing closure!!
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i2, &i2)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}

}
