# 编写第一个 Go 库 

让我们编写一个库，并让 hello 程序来使用它。

同样，第一步还是选择包路径（我们将使用 `github.com/biezhi/stringutil`） 并创建包目录：

```bash
$ mkdir $GOPATH/src/github.com/biezhi/stringutil
```

接着，在该目录中创建名为 `reverse.go` 的文件，内容如下：

```bash
// stringutil 包含有用于处理字符串的工具函数。
package stringutil

// Reverse 将其实参字符串以符文为单位左右反转。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

现在用 `go build` 命令来测试该包的编译：

```bash
$ go build github.com/biezhi/stringutil
```

当然，若你在该包的源码目录中，只需执行：

```bash
$ go build
```

即可。这不会产生输出文件。想要输出的话，必须使用 go install 命令，它会将包的对象放到工作空间的 pkg 目录中。

确认 `stringutil` 包构建完毕后，修改原来的 hello.go 文件（它位于 `$GOPATH/src/github.com/biezhi/hello`）去使用它：

```bash
package main

import (
	"fmt"

	"github.com/biezhi/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
```

无论是安装包还是二进制文件，go 工具都会安装它所依赖的任何东西。 因此当我们通过

```bash
$ go install github.com/biezhi/hello
```

来安装 hello 程序时，`stringutil` 包也会被自动安装。

运行此程序的新版本，你应该能看到一条新的，反向的信息：

```bash
$ hello
Hello, Go!
```

做完上面这些步骤后，你的工作空间应该是这样的：

```bash
bin/
	hello                 # 可执行命令
pkg/
	linux_amd64/          # 这里会反映出你的操作系统和架构
		github.com/biezhi/
			stringutil.a  # 包对象
src/
	github.com/biezhi/
		hello/
			hello.go      # 命令源码
		stringutil/
			reverse.go    # 包源码
```

注意 `go install` 会将 `stringutil.a` 对象放到 `pkg/linux_amd64` 目录中，它会反映出其源码目录。 这就是在此之后调用 go 工具，能找到包对象并避免不必要的重新编译的原因。 `linux_amd64` 这部分能帮助跨平台编译，并反映出你的操作系统和架构。

Go的可执行命令是静态链接的；在运行Go程序时，包对象无需存在。

