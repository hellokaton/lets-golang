# Go 中的包管理

## 包的命名

go 语言的包的命名，遵循简洁、小写、和 go 文件所在目录同名的原则，这样就便于我们引用，书写以及快速定位查找。

```go
package main

import "net/http"

func main() {
	http.ListenAndServe("127.0.0.1:8080",handler);
}
```

**以域名命名包**


```bash
package main

import "biezhi.me/utils" 
```

**使用 Github 账号命名包**

```go
package main

import "github.com/biezhi/utils"
```

这就是换成github.com命名的方式。

## main 包

当把一个go文件的包名声明为 main 时，就等于告诉go编译程序，我这个是一个可执行的程序，那么go编译程序就会尝试把它编译为一个二进制的可执行文件。

一个main的包，一定会包含一个main()函数，这种我们也不陌生，比如C和Java都有main()函数,它是一个程序的入口，没这个函数，程序就无法执行。

> 在go语言里，同时要满足main包和包含main()函数，才会被编译成一个可执行文件。

## 导入包

```go
import "fmt"
```

```go
import (
	"net/http"
	"fmt"
)
```

> 对于多于一个路径的包名，在代码中引用的时候，使用全路径最后一个包名作为引用的包名，比如net/http,我们在代码使用的是http，而不是net。

编译器会使用我们设置的这两个路径，再加上import导入的相对全路径来查找磁盘上的包，比如我们导入的fmt包，编译器最终找到的是/usr/local/go/fmt这个位置。

值得了解的是：对于包的查找，是有优先级的，编译器会优先在GOROOT里搜索，其次是GOPATH,一旦找到，就会马上停止搜索。如果最终都没找到，就报编译异常了。

## 远程包导入

```go
import "github.com/biezhi/moe"
```

1. 在 `GOPATH` 下搜索
2. go get 下载

## 命名导入

```go
package main

import (
	"fmt"
	myfmt "mylib/fmt"
)

func main() {
	fmt.Println()
	myfmt.Println()
}
```

```go
package main
import (
	_ "mylib/fmt"
)
```

## 包的init函数

**init 函数**

```go
package mysql

import (
	"database/sql"
)

func init() {
	sql.Register("mysql", &MySQLDriver{})
}
```

因为我们只是想执行这个mysql包的init方法，并不想使用这个包，所以我们在导入这个包的时候，需要使用_重命名包名，避免编译错误。

**静默导入**

```go
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

db, err := sql.Open("mysql", "user:password@/dbname")
```

