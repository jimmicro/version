# version

The "version" package is a utility that facilitates setting and retrieving the version during the compilation process. It provides a convenient way to manage and identify different versions of a software project. By using the "version" package, developers can automate the process of setting version numbers during the build and easily retrieve the current version when needed.

## Usage

```go
package main
import (
    "fmt"

    "github.com/jimyag/version-go"
)

func main() {
    if len(os.Args[]) >1 && os.Args[1] == "-v" {
        fmt.Println("version: ", version.Version())    
    }
    fmt.Println("hello world")
}

```

build with

```bash
go build -ldflags="-X github.com/jimyag/version.version=1.0.0"
```

if you use git release/tag you can build with

```bash
go build -ldflags="-X github.com/jimyag/version.version=$(git describe --abbrev=0 --tags --always)"
````
