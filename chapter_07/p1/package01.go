package p1

import "fmt"

func Foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func init() {
	fmt.Println("init")
}
