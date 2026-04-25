package version

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
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

type versionInfo struct {
	gitTag    string
	buildTime string
}

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
	info := resolveVersionInfo(GitTag, BuildTime, debug.ReadBuildInfo)
	GitTag = info.gitTag
	BuildTime = info.buildTime
	return fmt.Sprintf(`{"gitTag":"%s","buildTime":"%s","goVersion":"%s"}`, info.gitTag, info.buildTime, runtime.Version())
}

func resolveVersionInfo(gitTag string, buildTime string, readBuildInfo func() (*debug.BuildInfo, bool)) versionInfo {
	if gitTag != "" && buildTime != "" {
		return versionInfo{
			gitTag:    gitTag,
			buildTime: buildTime,
		}
	}

	info, _ := readBuildInfo()
	if gitTag == "" {
		gitTag = buildVersion(info)
	}
	settings := buildSettings(info)
	if buildTime == "" {
		buildTime = settings["vcs.time"]
	}
	if gitTag == "" {
		gitTag = settings["vcs.revision"]
	}
	if gitTag == "" {
		gitTag = unknown
	}
	if buildTime == "" {
		buildTime = unknown
	}
	return versionInfo{
		gitTag:    gitTag,
		buildTime: buildTime,
	}
}

func buildVersion(info *debug.BuildInfo) string {
	if info == nil {
		return ""
	}
	if info.Main.Version == "" || info.Main.Version == "(devel)" {
		return ""
	}
	return info.Main.Version
}

func buildSettings(info *debug.BuildInfo) map[string]string {
	if info == nil {
		return nil
	}

	settings := make(map[string]string, len(info.Settings))
	for _, item := range info.Settings {
		settings[item.Key] = item.Value
	}
	return settings
}
