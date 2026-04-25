package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	fmt.Println("ok:", ok)
	if !ok || bi == nil {
		return
	}

	fmt.Println("GoVersion:", bi.GoVersion)
	fmt.Println("Path:", bi.Path)
	fmt.Println("Main.Path:", bi.Main.Path)
	fmt.Println("Main.Version:", bi.Main.Version)
	fmt.Println("Main.Sum:", bi.Main.Sum)

	for _, s := range bi.Settings {
		switch s.Key {
		case "vcs", "vcs.revision", "vcs.time", "vcs.modified":
			fmt.Printf("%s=%s\n", s.Key, s.Value)
		}
	}
}
