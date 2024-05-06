package version

import (
	"fmt"
	"os"
	"runtime/debug"
)

const (
	unknown = "unknown"
	devel   = "(devel)"
)

var version = unknown
var EnableCmd bool = false

func init() {
	if !EnableCmd {
		return
	}
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Println(os.Args[0], Version())
		os.Exit(0)
	}
}

func Version() string {
	if version != unknown {
		return version
	}
	v, ok := buildInfoVersion()
	if ok {
		return v
	}
	if v == devel {
		return version
	}
	return version
}

func buildInfoVersion() (string, bool) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", false
	}
	return info.Main.Version, true
}
