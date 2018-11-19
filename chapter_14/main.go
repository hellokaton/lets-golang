package main

import "fmt"

func useSlice() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice[4])
	slice[0] = 100
	fmt.Println(slice)
	newSlice := append(slice, 6, 7, 8)
	fmt.Println(newSlice)
}

func copySlice() {
	data := [...]int{0, 1, 2, 3, 4, 9: 0}
	s := data[:2:3]
	fmt.Println("s:", s, len(s), cap(s))
	s = append(s, 233, 444, 1, 2, 3)
	fmt.Println("s:", s, len(s), cap(s))
	fmt.Println("data:", data, len(data), cap(data))
	fmt.Println(&s[0], &data[0])
}

func reCapSlice() {
	s := make([]int, 0, 1)
	c := cap(s)
	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d \n", c, n)
			c = n
		}
	}
}

func copySlice_() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[8:]
	s2 := data[:5]
	copy(s2, s)
	fmt.Println("s:", s)
	fmt.Println("s2:", s2)
	fmt.Println("data:", data)
}

func change(s []int) {
	s[0] = 100
}

func main() {
	slice := []int{0, 1, 2, 3}
	fmt.Println(slice)
	change(slice)
	fmt.Println(slice)
}
