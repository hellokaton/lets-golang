# 变量 常量

## **var**

---

变量声明用 `var` 关键字，可以用在函数级别，也可以用在包级别，比如

```go
package main

var c, python, java bool    //这个就是在包级别

func main() {
  var i int               //这个就是在函数级别
}
```

**注意：**

`var` 多个变量时不能为每个变量设定类型（无论类型是否相同，无论是否有带初始值）比如：

| 语句 | 结果 |
| :-- | :-- |
| var i int, j int | 报错 |
| var i int, j int = 1, 2 | 报错 |
| var i string, j int | 报错 |
| var i string, j int = "a", 1 | 报错 |

var多个变量带初始值，如果类型不同，就不能指定变量类型，比如：

| 语句 | 结果 |
| :-- | :-- |
| var i, j = "aa", 2 | 正确 |
| var i, j int = "aa", 2 | 报错 |

var是否跟类型（包括var单个变量或者var多个变量）：

1. 跟类型：
	1. 带初始值: OK

		单个变量例子：`var i int = 1`

		多个变量例子：`var i, j int = 1, 2`

	2. 不带初始值: 根据类型自动设初始值，如int初始值为0，string初始值为""（空字符）

		单个变量例子：`var i int`，则 i 为 0

		多个变量例子：`var i, j int`，则 i 和 j 为 0

2. 不跟类型：
	1. 带初始值: 根据初始值自动判断类型，如 1 为 int 类型，"a" 为 string 类型

		单个变量例子：`var i = 1`，则 i 为 int 类型

		多个变量例子：

		1. `var i, j = 1, 2`，则 i 和 j 为 int 类型

		2. `var i, j = true, "a"`，则 i 为 bool 类型，j 为 string 类型

	2. 不带初始值: 报错

		单个变量例子：`var i`，会报错

		多个变量例子：`var i, j`，会报错

也可以这么声明

```text
var (
	A int = 100
	B string
)
```

**和 const 区别**

`var` 可以不赋值（会自动赋值初始值），而 `const` 必须赋值

## **:=**

---

`:=` 是短声明变量，也叫简洁赋值语句

函数内，`:=` 用在明确类型的地方（但不能显式指定类型），可以用于替代`var`定义。

函数外（即包级别）的每个语句都必须以关键字开始（`var`、`func`等），`:=` 不能使用在函数外，如：

```go
var k = 3	//OK
k := 3		//报错

func main() {
	k := 3			//OK
	var k int = 3	//OK
	k int := 3		//报错
	var k := 3		//报错
	var k int := 3	//报错
}
```

## 作用域

1. 在包级别定义的变量，是全局变量，即在main函数里对这个变量做的修改，或者在其他函数里对这个变量做的修改，都是真实改了这个变量的值
2. 在函数里（包括 main 函数或其他函数）定义的变量，只在自己函数里有效，若在其他函数里对这个变量做修改，会提示 undefined
3. 可以存在变量名相同，但作用域不同的多个变量。（个人觉得太晦涩，不建议这么用）
4. 函数内部的函数（即函数值），有前后顺序要求

**例1. 包级别的变量能否被函数引用?**

```go
var A int = 1

func foo() {
  fmt.Println(A)
}

func main() {
  foo()
}
```

**例2. 包级别的函数是否有顺序要求?**

```go
var A int = 1

func main() {
  foo()
}

func foo() {
  fmt.Println(A)
}
```

**例3. 一个函数里的变量能否被包级别的另一个函数引用?**

```go
func foo() {
  fmt.Println(A)
}

func main() {
  var A int = 1
  foo()
}
```

**例4. 一个函数里的变量能否被内嵌的函数值引用?**

```go
func main() {
    var A int = 1
    foo := func () {
        fmt.Println(A)
    }
    foo()
}
```

**例5. 函数值是否有顺序要求?**

```go
func main() {
    foo := func () {
        fmt.Println(A)
    }
    var A int = 1
    foo()
}
```

**例6. 包级别的变量能否在函数内部的函数值里引用?**

```go
var A int = 1
func main() {
    foo := func () {
        A = 2
    }
    foo()
    fmt.Println(A)
}
```

**例7. 函数值嵌套函数值，还能否引用外部变量?**

```go
func main() {
    A := 1
    foo := func () {
        bar := func () {
            A = 2
        }
        bar()
    }
    foo()
    fmt.Println(A)
}
```

**例8. 是否可以存在不同作用域，但名字相同的变量(本例十分经典)?**

```go
func main() {
    A := 1
    foo := func () {
        A = 2
        A := 3
        A = 4
        fmt.Println(A)
    }
    foo()
    fmt.Println(A)
}
```

**例9. 本例与例8相似，若一开始就声明同名的变量名，那么还能否引用外部同名变量?**

```go
func main() {
    A := 1
    foo := func () {
        A := 3
        A = 2
        fmt.Println(A)
    }
    foo()
    fmt.Println(A)
}
```

## 常量

**声明**

---

用 `const` 声明，无论在包级别还是函数级别。不能用 `:=`

常量只能是字符串、布尔、数字类型（整数、浮点数、复数）的值

```go
const foo = "bar"
```

上面声明一个常量，常量名叫foo，值是"bar"，foo的类型自动推导为string

也可以`#!go const foo string = "bar"`来直接指定类型

注意%T是查看类型，即string还是int，而看不出是常量还是变量

也可以这么声明

```text
const (
	A int = 100
	B = "hello"
)
```

**和 var 的区别**

`var` 可以不赋值（会自动赋值初始值），而 `const` 必须赋值

