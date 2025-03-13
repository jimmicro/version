package version

import (
	"fmt"
	"os"
	"runtime"
)

const (
	unknown  = "unknown"
	trueStr  = "true"
	falseStr = "false"
)

var (
	BuildTime        = ""
	GitTag           = ""
	EnableCmd string = trueStr
)

func init() {
	if EnableCmd != trueStr {
		return
	}

	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Println(Version())
		os.Exit(0)
	}
}

func Version() string {
	if GitTag == "" {
		GitTag = unknown
	}
	if BuildTime == "" {
		BuildTime = unknown
	}
	return fmt.Sprintf(`{"gitTag":"%s","buildTime":"%s","goVersion":"%s"}`, GitTag, BuildTime, runtime.Version())
}
