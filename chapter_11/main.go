package main

import "fmt"

func main() {
	//var foo int
	//var bar *int
	//
	//bar = &foo
	//
	//*bar = 2
	//
	//foo = 3
	//
	//fmt.Printf("%v %T \n", foo, foo)
	//fmt.Printf("%v %T \n", *bar, bar)

	//var a int = 1
	//var b *int = &a
	//var c **int = &b
	//var x int = *b
	//fmt.Println("a = ", a)  // 1
	//fmt.Println("&a = ", &a) // 内存A
	//fmt.Println("*&a = ", *&a) // 1
	//fmt.Println("b = ", b) // 内存A
	//fmt.Println("&b = ", &b) // 内存B
	//fmt.Println("*&b = ", *&b)// 内存A
	//fmt.Println("*b = ", *b)//1
	//fmt.Println("c = ", c)// 内存B
	//fmt.Println("*c = ", *c)// 内存A
	//fmt.Println("&c = ", &c)// 内存C
	//fmt.Println("*&c = ", *&c)// 内存B
	//fmt.Println("**c = ", **c)//1
	//fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)//1
	//fmt.Println("x = ", x)//1

	var a = 3
	var b = 4
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a *int, b *int) {
	//x := *a
	//*a = *b
	//*b = x
	*a, *b = *b, *a
}
