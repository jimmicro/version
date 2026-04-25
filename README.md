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

Version info is resolved with this priority:

1. `debug.ReadBuildInfo().Main.Version` if it is available and not `(devel)`
2. `GitTag` injected by `-ldflags -X`
3. `debug.ReadBuildInfo()` `vcs.revision`
4. `unknown`

Build time is resolved with this priority:

1. `BuildTime` injected by `-ldflags -X`
2. `debug.ReadBuildInfo()` `vcs.time`
3. `unknown`

When the binary is built with VCS stamping enabled (default `-buildvcs=auto`),
Go may embed `vcs.revision`, `vcs.time`, and `vcs.modified` into the binary.

Note: if you need a specific git tag value from a local build, keep passing
`GitTag` with `-ldflags -X`. Go build info does not expose a dedicated
`vcs.tag` field.

```bash
GIT_TAG=$(git describe --tags --always --dirty="-dev")
BUILD_TIME=$(TZ='Asia/Shanghai' date "+%Y-%m-%d-%H-%M-%S")
LDFLAGS="-s -w  -X 'github.com/jimmicro/version.GitTag=$GIT_TAG' \
    -X 'github.com/jimmicro/version.BuildTime=$BUILD_TIME'"

go build -ldflags="$LDFLAGS" -o main main.go
```
