package main

import "fmt"

func change(a *[3]int) {
	a[1] = 233
}

func main() {
	arr1 := [...]int{1, 2, 3}

	//ptrArr := [3]*int{1:new(int)}
	//
	//fmt.Println("ptrArr[1]:", *ptrArr[1])
	change(&arr1)

	//fmt.Println("ptrArr[1]:", *ptrArr[1])
	//arr1[1] = 233
	fmt.Println(arr1)
	//
	//for i:=0; i<len(arr1);i++{
	//    fmt.Println("索引:", i , ", 值:", arr1[i])
	//}
	//
	//for i, v := range arr1 {
	//    fmt.Println("索引:", i , ", 值:", v)
	//}
	//
	//var arr2 = arr1
	//fmt.Println(arr2)

	//coder := [4]string{1:"biezhi", 3: "jack"}
	//fmt.Println(coder[0])
	//fmt.Println(coder[1])
	//fmt.Println(coder[2])
	//fmt.Println(coder[3])
	//
	//arr2 := [3][4]int{
	//    {0,1,2,3},
	//    {4,5,6,7},
	//    {4,5,6,7},
	//}
	//fmt.Println(arr2)
}
