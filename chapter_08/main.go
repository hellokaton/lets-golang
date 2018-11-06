package main

import "fmt"

var bar string = "bar"
var flag bool
var empty string
var score int

var (
	A int = 100
	B string
)

func main() {
	var msg string = "hello"
	fmt.Println(msg)

	foo := "hello"
	fmt.Println(foo)

	num := 1
	fmt.Println(num)

	fmt.Println("flag:", flag)
	fmt.Println("score:", score)
	fmt.Println("empty:", empty)

	// var i, j int
	// fmt.Println(i, j)

	// var i int, j int = 1, 2
	var i, j int = 1, 2
	fmt.Println(i, j)

	var i1, j1 = "aa", 2
	fmt.Println(i1, j1)

	var q string
	fmt.Println(q)
	fooFunc()

	foo2 := func() {
		// A = 102
		var A string = "200"
		bar := func() {
			A = "103"
		}
		bar()
	}
	// var A2 int = 1
	foo2()

	A3 := 100
	foo3 := func() {
		A3 := 101
		A3 = 102
		fmt.Println("foo3, A3:", A3)
	}
	foo3()
	fmt.Println("A3:", A3)
	fmt.Println("A:", A)
}

func fooFunc() {
	// fmt.Println(q)
}
