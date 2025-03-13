# version

The "version" package is a utility that facilitates setting and retrieving the version during the compilation process. It provides a convenient way to manage and identify different versions of a software project. By using the "version" package, developers can automate the process of setting version numbers during the build and easily retrieve the current version when needed.

## Usage

```go
package main
import (
    "fmt"

    _ "github.com/jimmicro/version"
)

func main() {
    fmt.Println("hello world")
}

```

build with git tag build time and go version

```bash
GIT_TAG=$(git describe --tags --always --dirty="-dev")
BUILD_TIME=$(TZ='Asia/Shanghai' date "+%Y-%m-%d-%H-%M-%S")
LDFLAGS="-s -w  -X 'github.com/jimmicro/version.GitTag=$GIT_TAG' \
    -X 'github.com/jimmicro/version.BuildTime=$BUILD_TIME'"

go build -ldflags="$LDFLAGS" -o main main.go
```
