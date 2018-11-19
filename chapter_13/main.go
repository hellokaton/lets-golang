package main

import "fmt"

func initSlice() {
	s1 := []int{1, 2, 3}
	fmt.Println(s1, len(s1), cap(s1))

	s2 := []int{0, 1, 2, 5: 100}
	fmt.Println(s2, len(s2), cap(s2))
}

func makeSlice() {
	s1 := make([]int, 3)
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 6)
	fmt.Println(s2, len(s2), cap(s2))
}

func emptySlice() {
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("len = %d, cap = %d, %#v\n", len(nilSlice), cap(nilSlice), nilSlice)
	fmt.Printf("len = %d, cap = %d, %#v\n", len(emptySlice), cap(emptySlice), emptySlice)
}

func arraySlice() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("expression\tslice\t\t\t\t\tlen\tcap\t")
	// 省略了 low
	fmt.Printf("data[:6:8]\t%v\t\t\t%d\t%d\n", data[:6:8], len(data[:6:8]), cap(data[:6:8]))
	// 省略了 high max
	fmt.Printf("data[5:]\t%v\t\t\t\t%d\t%d\n", data[5:], len(data[5:]), cap(data[5:]))
	// 省略了 low max
	fmt.Printf("data[:3]\t%v\t\t\t\t\t%d\t%d\n", data[:3], len(data[:3]), cap(data[:3]))
	// 全部省略
	fmt.Printf("data[:]\t\t%v\t%d\t%d\n", data[:], len(data[:]), cap(data[:]))
	s1 := data[:]
	s1[0] = 200
	fmt.Println(s1)
	fmt.Println(data)
}

func main() {
	reSlice()
}

func reSlice() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5]
	// 2 3 4
	fmt.Println(s1, len(s1), cap(s1))
	s1[1] = 100
	// 2 100 4
	s2 := s1[2:6]
	// 4 5 6 7
	fmt.Println("s2:", s2)
	s2[2] = 200
	fmt.Println(s)
}
