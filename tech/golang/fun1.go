package main

import (
	"fmt"
)

type fn func(int)

var v int

func t1(f fn) {
	fmt.Println("pre func ", v)
	f(1)
	fmt.Println("after func ", v)
}

func fx(x int) {
	fmt.Println("x")
}

func main() {
	fmt.Println("test")

	var f1 = func(i int) {
		fmt.Println("y")
		v++
	}

	t1(f1)

	t1(f1)

	t1(fx)
}
