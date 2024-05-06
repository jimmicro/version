package version

import (
	"fmt"
	"os"
	"runtime/debug"
)

const (
	unknown  = "unknown"
	devel    = "(devel)"
	trueStr  = "true"
	falseStr = "false"
)

var version = ""
var enableCmd string = falseStr

func init() {
	if enableCmd != trueStr {
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
	if !ok {
		return unknown
	}
	if v == devel {
		return unknown
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
