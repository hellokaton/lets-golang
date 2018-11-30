# Go 中的命令 

## Go 命令概览

```bash
» go
Go is a tool for managing Go source code.
Usage:
	go command [arguments]
The commands are:
	build       compile packages and dependencies
	clean       remove object files
	doc         show documentation for package or symbol
	env         print Go environment information
	bug         start a bug report
	fix         run go tool fix on packages
	fmt         run gofmt on package sources
	generate    generate Go files by processing source
	get         download and install packages and dependencies
	install     compile and install packages and dependencies
	list        list packages
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         run go tool vet on packages
Use "go help [command]" for more information about a command.
Additional help topics:
	c           calling between Go and C
	buildmode   description of build modes
	filetype    file types
	gopath      GOPATH environment variable
	environment environment variables
	importpath  import path syntax
	packages    description of package lists
	testflag    description of testing flags
	testfunc    description of testing functions
Use "go help [topic]" for more information about that topic.
```

## go build

```bash
usage: go build [-o output] [-i] [build flags] [packages]
```


```bash
go build
go build .
go build hello.go
```

**编译指定包**

```bash
go build github.com/biezhi/stringutil
```

**编译所有包**

```bash
go build github.com/biezhi/stringutil/...
```

**查看环境信息**

```bash
» go env
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/biezhi/Library/Caches/go-build"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/biezhi/course/golang"
GORACE=""
GOROOT="/usr/local/go"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0
```

`GOOS` 是目标操作系统，它的值为：

- darwin
- freebsd
- linux
- windows
- android
- dragonfly
- netbsd
- openbsd
- plan9
- solaris

`GOARCH` 指的是目标处理器的架构，目前支持的有：

- arm
- arm64
- 386
- amd64
- ppc64
- ppc64le
- mips64
- mips64le
- s390x

**交叉编译**

```bash
GOOS=linux GOARCH=amd64 go build github.com/biezhi/hello
```

> 具体组合参考 https://golang.org/doc/install/source#environment

更多关于go build的用户可以通过以下命令查看:

```bash
go help build
```

## go clean

```bash
usage: go clean [-i] [-r] [-n] [-x] [build flags] [packages]
```


```bash
go help clean
```

## go run

```bash
➜  ~ go help run
usage: go run [build flags] [-exec xprog] gofiles... [arguments...]
Run compiles and runs the main package comprising the named Go source files.
A Go source file is defined to be a file ending in a literal ".go" suffix.
By default, 'go run' runs the compiled binary directly: 'a.out arguments...'.
If the -exec flag is given, 'go run' invokes the binary using xprog:
	'xprog a.out arguments...'.
If the -exec flag is not given, GOOS or GOARCH is different from the system
default, and a program named go_$GOOS_$GOARCH_exec can be found
on the current search path, 'go run' invokes the binary using that program,
for example 'go_nacl_386_exec a.out arguments...'. This allows execution of
cross-compiled programs when a simulator or other execution method is
available.
For more about build flags, see 'go help build'.
```

**接收参数**

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input Args:",os.Args[1])
}
```

## go install

```bash
➜  hello go help install
usage: go install [build flags] [packages]
Install compiles and installs the packages named by the import paths,
along with their dependencies.
```

## go get

**下载GO依赖库**

```bash
go get github.com/biezhi/moe
```

`go get` 支持大多数版本控制系统(VCS)，如Git。

**更新依赖**

```bash
go get -u github.com/biezhi/moe
```

**查看进度**

```bash
go get -v github.com/biezhi/moe
```

关于go get 命令的更多用法，可以使用如下命令查看:

```bash
go help get
```

## go fmt

```go
func main() {
fmt.Println("Input Args:",os.Args[1])
}
```

```bash
» gofmt -h  
usage: gofmt [flags] [path ...]
  -cpuprofile string
    	write cpu profile to this file
  -d	display diffs instead of rewriting files
  -e	report all errors (not just the first 10 on different lines)
  -l	list files whose formatting differs from gofmt's
  -r string
    	rewrite rule (e.g., 'a[b:len(a)] -> a[b:]')
  -s	simplify code
  -w	write result to (source) file instead of stdout
```

`go fmt` 为我们统一了代码风格，这样我们在整个团队协作中发现，所有代码都是统一的，像一个人写的一样。所以我们的代码在提交到git库之前，一定要使用go fmt进行格式化，现在也有很多编辑器也可以在保存的时候，自动帮我们格式化代码。

## go vet

1. `Printf` 这类的函数调用时，类型匹配了错误的参数。
2. 定义常用的方法时，方法签名错误。
3. 错误的结构标签。
4. 没有指定字段名的结构字面量。

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Printf(" 哈哈",3.14)
}
```

```bash
usage: go vet [-n] [-x] [build flags] [packages]
```

## go test

该命令用于Go的单元测试，它也是接受一个包名作为参数，如果没有指定，使用当前目录。
go test运行的单元测试必须符合go的测试要求。

1. 写有单元测试的文件名，必须以_test.go结尾。
2. 测试文件要包含若干个测试函数。
3. 这些测试函数要以Test为前缀，还要接收一个*testing.T类型的参数。

```go
package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(1,2) == 3 {
		t.Log("1+2=3")
	}
	if Add(1,1) == 3 {
		t.Error("1+1=3")
	}
}
```

更多关于go test命令的使用，请通过如下命令查看。

```bash
go help test
```

以上这些，主要时介绍的go这个开发工具常用的命令，熟悉了之后可以帮助我们更好的开发编码。

其他 `go` 命令的介绍，比如 `package` 是什么等等，可以直接使用 `go help [topic]` 命令查看。

```bash
Additional help topics:
	c           calling between Go and C
	buildmode   description of build modes
	filetype    file types
	gopath      GOPATH environment variable
	environment environment variables
	importpath  import path syntax
	packages    description of package lists
	testflag    description of testing flags
	testfunc    description of testing functions
Use "go help [topic]" for more information about that topic.
```
