package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var s string = "hello world \n"
	var s2 string = `helle world \n`
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(len(s))
	fmt.Printf("%c\n", s[0])
	fmt.Printf("%c\n", s[len(s)-3])

	// 字符串的修改
	str := "hello world"
	strArr := []byte(str)
	strArr[0] = 'X'
	fmt.Println(str)
	fmt.Println(string(strArr))

	u := "王爵的技术小黑屋"
	us := []rune(u)
	us[0] = 'w'
	fmt.Println(u)
	fmt.Println(string(us))

	u2 := "王爵的技术小黑屋"
	us2 := []byte(u2)
	us2[0] = 'w'
	fmt.Println(u2)
	fmt.Println(string(us2))

	// 字符串拼接
	foo := "hello" + "world" + string(3)
	fmt.Println(foo)

	// 字符串的遍历
	str = "hello 世界"
	for i, n := 0, len(str); i < n; i++ {
		var ch2 byte = str[i]
		fmt.Printf("%d - %c , ", i, ch2)
	}
	fmt.Println("\n==================")
	for i, ch := range str {
		fmt.Printf("%d: %v, %c \n", i, ch, ch) // ch 是 rune类型
		// fmt.Printf("%c .", ch)
	}

	fmt.Println("是否以 hello 开头:", strings.HasPrefix(str, "hello"))
	fmt.Println("是否包含世界:", strings.Contains(str, "世界"))
	fmt.Println("o出现的位置:", strings.Index(str, "o"))

	// bytes.Buffer
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	fmt.Println(buffer.String())
}
