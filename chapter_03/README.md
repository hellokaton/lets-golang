# Hello Go 程序 

## 包路径

标准库中的包有给定的短路径，比如 `"fmt"` 和 `"net/http"`。 对于你自己的包，你必须选择一个基本路径，来保证它不会与将来添加到标准库， 或其它扩展库中的包相冲突。

如果你将你的代码放到了某处的源码库，那就应当使用该源码库的根目录作为你的基本路径。 例如，若你在 GitHub 上有账户 `github.com/biezhi` 那么它就应该是你的基本路径。

注意，在你能构建这些代码之前，无需将其公布到远程代码库上。只是若你某天会发布它， 这会是个好习惯。在实践中，你可以选择任何路径名，只要它对于标准库和更大的Go生态系统来说， 是唯一的就行。

> fmt包含有格式化 I/O 函数，类似于 C 语言的 printf 和 scanf。格式字符串的规则来源于 C 但更简单一些。

我们将使用 `github.com/biezhi` 作为基本路径。在你的工作空间里创建一个目录， 我们将源码存放到其中：

```bash
$ mkdir -p $GOPATH/src/github.com/biezhi/hello
```

接着，在该目录中创建名为 `hello.go` 的文件，其内容为以下Go代码：

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go")
}
```

现在你可以用 go 工具构建并安装此程序了：

```bash
$ go install github.com/biezhi/hello
```

注意，你可以在系统的任何地方运行此命令。go 工具会根据 GOPATH 指定的工作空间，在 `github.com/biezhi/hello` 包内查找源码。

若在从包目录中运行 `go install`，也可以省略包路径：

```bash
$ cd $GOPATH/src/github.com/biezhi/hello
$ go install
```

此命令会构建 hello 命令，产生一个可执行的二进制文件。 接着它会将该二进制文件作为 `hello`（在 Windows 下则为 `hello.exe`）安装到工作空间的 bin 目录中。 在我们的例子中为 `$GOPATH/bin/hello`，具体一点就是 $HOME/go/bin/hello。

go 工具只有在发生错误时才会打印输出，因此若这些命令没有产生输出， 就表明执行成功了。

现在，你可以在命令行下输入它的完整路径来运行它了：

```bash
$ $GOPATH/bin/hello
Hello, Go
```

若你已经将 `$GOPATH/bin` 添加到 `PATH` 中了，只需输入该二进制文件名即可：

```bash
$ hello
Hello, Go.
```

若你使用源码控制系统，那现在就该初始化仓库，添加文件并提交你的第一次更改了。 再次强调，这一步是可选的：你无需使用源码控制来编写Go代码。

```bash
$ cd $GOPATH/src/github.com/biezhi/hello
$ git init
Initialized empty Git repository in /Users/biezhi/course/golang/src/github.com/biezhi/hello/.git/
$ git add hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 1 insertion(+)
  create mode 100644 hello.go
```