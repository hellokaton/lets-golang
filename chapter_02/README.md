# 搭建 Go 环境

## 安装

**系统需求**

gc 编译器支持以下操作系统及架构。在开始前，请确保你的系统满足这些需求。 若你的OS及架构不在此列表中，那么 gccgo 可能支持你的设置， 详情请访问设置并使用gccgo。


![IMAGE](quiver-image-url/F96609BD1616FC08AC1B66BA47D73EF2.jpg =897x166)

- 仅当你打算用cgo时才需要gcc。
- 只需为 Xcode安装命令行工具即可。 若你已经安装了Xcode 4.3+，只需从下载配置面板的组件标签内安装它即可。

下载此 [压缩包](https://golang.org/dl/) 并提取到 `/usr/local` 目录，在 `/usr/local/go` 中创建 Go目录。例如：

```bash
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

该压缩包的名称可能不同，这取决于你安装的Go版本和你的操作系统以及处理器架构。

（此命令必须作为root或通过 sudo 运行）

要将 `/usr/local/go/bin` 添加到 `PATH` 环境变量， 你需要将此行添加到你的 `/etc/profile`（全系统安装）或 `$HOME/.profile` 文件中：

**安装到指定位置**

Go二进制发行版假定它们会被安装到 `/usr/local/go` （或Windows下的 `c:\Go`）中，但也可将Go工具安装到不同的位置。 此时你必须设置 `GOROOT` 环境变量来指出它所安装的位置。

例如，若你将Go安装到你的home目录下，你应当将以下命令添加到 `$HOME/.profile` 文件中：

```bash
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

> 注：GOROOT 仅在安装到指定位置时才需要设置。

## 测试安装

```bash
go version
```

## 代码组织

### 概览

- Go程序员通常将所有的Go代码保存在一个_工作区_中。
- 工作区包含许多版本控制存储库 （例如，由Git管理）。
- 每个存储库都包含一个或多个包。
- 每个软件包由一个目录中的一个或多个Go源文件组成。
- 包的目录路径决定了它的导入路径。

请注意，这与其他编程环境不同，在这些编程环境中，每个项目都有一个单独的工作区，工作区与版本控制存储库紧密相关。

### 工作区

go 工具为公共代码仓库中维护的开源代码而设计。 无论你会不会公布代码，该模型设置工作环境的方法都是相同的。

Go代码必须放在工作空间内。它其实就是一个目录，其中包含三个子目录：

* src 目录包含Go的源文件，它们被组织成包（每个目录都对应一个包）
* pkg目录存放编译好的库文件, 主要是*.a文件
* bin 目录包含可执行命令

go 工具用于构建源码包，并将其生成的二进制文件安装到 pkg 和 bin 目录中。

src 子目录通常包会含多种版本控制的代码仓库（例如Git）， 以此来跟踪一个或多个源码包的开发。

以下例子展现了实践中工作空间的概念：

```bash
bin/
	streak                         # 可执行命令
	todo                           # 可执行命令
pkg/
	linux_amd64/
		code.google.com/p/goauth2/
			oauth.a                # 包对象
		github.com/nf/todo/
			task.a                 # 包对象
src/
	code.google.com/p/goauth2/
		.hg/                       # mercurial 代码库元数据
		oauth/
			oauth.go               # 包源码
			oauth_test.go          # 测试源码
	github.com/nf/
		streak/
		.git/                      # git 代码库元数据
			oauth.go               # 命令源码
			streak.go              # 命令源码
		todo/
		.git/                      # git 代码库元数据
			task/
				task.go            # 包源码
			todo.go                # 命令源码
```

此工作空间包含三个代码库（goauth2、streak 和 todo），两个命令（streak 和 todo） 以及两个库（oauth 和 task）。

命令和库从不同的源码包编译而来。稍后我们会对讨论它的特性。

## GOPATH

- GOPATH 环境变量指定了你的工作空间位置。它或许是你在开发Go代码时， 唯一需要设置的环境变量。

首先创建一个工作空间目录，并设置相应的 GOPATH。你的工作空间可以放在任何地方， 在此文档中我们使用 `$HOME/work`。注意，它绝对不能和你的Go安装目录相同。 （另一种常见的设置是 `GOPATH=$HOME`。）

```bash
export GOPATH=$HOME/course/golang
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

## 开发工具

- vscode
- Goland



