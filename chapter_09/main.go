package main

import (
	"fmt"
	"math"
)

func main() {
	var v1 bool
	v1 = true
	v2 := (1 == 2)
	fmt.Println(v1, v2)
	var b bool
	b = 1 != 0
	fmt.Println(b)

	var value1 int32
	value2 := 125 // 被推断为int类型
	value1 = int32(value2)
	fmt.Println(value1, value2)

	var i int32
	var j int64
	i, j = 1, 2
	if i == 1 || j == 2 {
		fmt.Println("==")
	}

	fmt.Println("=============")
	var fvalue1 float32
	fvalue1 = 122
	fvalue2 := 122.0
	fmt.Println(fvalue1, fvalue2)
	fmt.Printf("%T, %T \n", fvalue1, fvalue2)

	fmt.Println(math.Dim(1, 2))

	var cp1 complex64 = 12 + 10i
	fmt.Println(cp1)
}
