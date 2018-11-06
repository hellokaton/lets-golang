# Go 基本数据类型

Go 语言有以下几种基础类型：

- 布尔类型：bool，占1字节
- 整型：int8、byte（uint8）、int16、int、uint、uintptr等
- 浮点类型：float32、float64
- 复数类型：complex64、complex128
- 字符串：string
- 字符类型：rune（int32）
- 错误类型：error

还有下面的复合类型：

- 指针（pointer）
- 数组（array）
- 切片（slice）
- 字典（map）
- 通道（chan）
- 结构体（struct）
- 接口（interface）

## 布尔类型

- true
- false

## 整数

- int8（-128 -> 127）
- int16（-32768 -> 32767）
- int32（-2,147,483,648 -> 2,147,483,647）
- int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）

## 无符号整数

- uint8（0 -> 255）
- uint16（0 -> 65,535）
- uint32（0 -> 4,294,967,295）
- uint64（0 -> 18,446,744,073,709,551,615）

## 浮点型（IEEE-754 标准）

- float32（+- 1e-45 -> +- 3.4 * 1e38）
- float64（+- 5 * 1e-324 -> 107 * 1e308）

> IEEE二进制浮点数算术标准（IEEE 754）是20世纪80年代以来最广泛使用的浮点数运算标准，为许多CPU与浮点运算器所采用。这个标准定义了表示浮点数的格式（包括负零-0）与反常值（denormal number）），一些特殊数值（无穷（Inf）与非数值（NaN）），以及这些数值的“浮点数运算符”；它也指明了四种数值舍入规则和五种例外状况（包括例外发生的时机与处理方式）。
> 
> https://zh.wikipedia.org/wiki/IEEE_754

**复数类型**

1. complex64
32 位实数和虚数
2. complex128
64 位实数和虚数

```go
var value1 complex64 // 由2个float32构成的复数类型
value1 = 3.2 + 12i
value2 := 3.2 + 12i        // value2是complex128类型
value3 := complex(3.2, 12) // value3结果同value2
var value4 complex128 = 3.2 + 12i
```

## 一些示例

```go
// Bool类型
var my_bool bool = true

// 字符串类型
var my_string string = "hello, world!"

// int: 有符号整形，根据系统架构自动判断是int8，int16，int32还是int64
// 比如当前系统是64位系统，则为int64
var my_int_min int = -9223372036854775808
var my_int_max int = 9223372036854775807

// int8: 有符号 8 位整型 (-128 到 127)
var my_int8_min int8 = -128
var my_int8_max int8 = 127

// int16: 有符号 16 位整型 (-32768 到 32767)
var my_int16_min int16 = -32768
var my_int16_max int16 = 32767

// int32: 有符号 32 位整型 (-2147483648 到 2147483647)
var my_int32_min int32 = -2147483648
var my_int32_max int32 = 2147483647

// int64: 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
var my_int64_min int64 = -9223372036854775808
var my_int64_max int64 = 9223372036854775807

// uint: 无符号整形，根据系统架构自动判断是uint8，uint16，uint32还是uint64
// 比如当前系统是64位系统，则为uint64
var my_uint_min uint = 0
var my_uint_max uint = 18446744073709551615

// uint8: 无符号 8 位整型 (0 到 255)
var my_uint8_min uint8 = 0
var my_uint8_max uint8 = 255

// uint16: 无符号 16 位整型 (0 到 65535)
var my_uint16_min uint16 = 0
var my_uint16_max uint16 = 65535

// uint32: 无符号 32 位整型 (0 到 4294967295)
var my_uint32_min uint32 = 0
var my_uint32_max uint32 = 4294967295

// uint64: 无符号 64 位整型 (0 到 18446744073709551615)
var my_uint64_min uint64 = 0
var my_uint64_max uint64 = 18446744073709551615

// uintptr: 无符号整型，用于存放一个指针，可以足够保存指针的值的范围。和uint范围相同，根据系统架构自动判断
var my_uintptr_min uintptr = 0
var my_uintptr_max uintptr = 18446744073709551615

// byte: uint8的别名
var my_byte_min byte = 0
var my_byte_max byte = 255

// rune: int32的别名。代表1个unicode码
var my_rune_min rune = -2147483648
var my_rune_max rune = 2147483647

// float32: 单精度浮点数，在C语言里等同于float
// float64: 双精读浮点数，在C语言里等同于double
// 如果不写类型，则为float64（暂时不知道这个是根据系统架构判断还是默认就是float64）
// 从结果可以看出:
//   float32只能容纳8位数字（包括小数点前后数字，不包括小数点，超过8位会将四舍五入保留8位）
//   float64可以容纳比较多的数字（具体暂时还没测），而且这种双精度我也一直没搞懂，很复杂
//   当符合要求时候会自动用科学计数法来表示，要注意
var my_float32 float32 = 10086.141592653
var my_float64 float64 = 10086.141592653
var my_float_auto = 10086.141592653
```