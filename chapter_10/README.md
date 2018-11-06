# 字符串和字符

## 字符串

- Go 中的字符串根据需要占用 1 至 4 个字节
- 和 C/C++ 一样，Go 中的字符串是根据长度限定，而非特殊字符。
- str[i]的方式只对纯 ASCII 码的字符串有效。
- 和字符串有关的包：strings, strconv
- 解释字符串
- 非解释字符串

## strings 和 strconv 包

- 前缀和后缀: `strings.HasPrefix(s, prefix string) bool`
- 字符串包含关系: `strings.Contains(s, substr string) bool`
- 判断子字符串或字符在父字符串中出现的位置: `strings.Index(s, str string) int`
- 字符串替换: `strings.Replace(str, old, new, n) string`
- 统计字符串出现次数: `strings.Count(s, str string) int`
- 重复字符串: `strings.Repeat(s, count int) string`
- 修改字符串大小写: `strings.ToLower(s) string`
- 修剪字符串: `strings.TrimSpace(s)`
- 分割字符串: `strings.Split(s, sep)`
- 字符串与其它类型的转换: `strconv.Itoa(i int) string`

## bytes 包

- bytes 包和字符串包十分类似，而且它还包含一个十分有用的类型 Buffer。这是一个 bytes 的定长 buffer，提供 Read 和 Write 方法，因为读写不知道长度的 bytes 最好使用 buffer。
- Buffer 可以这样定义：var buffer bytes.Buffer 或者 new 出一个指针：var r *bytes.Buffer = new(bytes.Buffer) 或者通过函数：func NewBuffer(buf []byte) *Buffer，这就创建了一个 Buffer 对象并且用 buf 初始化好了；NewBuffer 最好用在从 buf 读取的时候使用。
- 通过 buffer 串联字符串：类似于 Java 的 StringBuilder 类 创建一个 Buffer，通过 buffer.WriteString(s) 方法将每个 string s 追加到后面，最后再通过 buffer.String() 方法转换为 string

## 字符类型

**byte和rune**

- byte（实际上是uint8的别名），代表 `UTF-8` 字符串的单个字节的值
- rune（int32），代表单个 `Unicode` 字符

