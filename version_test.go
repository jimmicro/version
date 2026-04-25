package version

import (
	"runtime/debug"
	"testing"
)

func TestResolveVersionInfoUsesExplicitValues(t *testing.T) {
	info := resolveVersionInfo("v1.2.3", "2026-04-25 10:00:00", func() (*debug.BuildInfo, bool) {
		t.Fatal("build info should not be read when explicit values are set")
		return nil, false
	})

	if info.gitTag != "v1.2.3" {
		t.Fatalf("gitTag = %q, want %q", info.gitTag, "v1.2.3")
	}
	if info.buildTime != "2026-04-25 10:00:00" {
		t.Fatalf("buildTime = %q, want %q", info.buildTime, "2026-04-25 10:00:00")
	}
}

func TestResolveVersionInfoPrefersMainVersion(t *testing.T) {
	info := resolveVersionInfo("", "", func() (*debug.BuildInfo, bool) {
		return &debug.BuildInfo{
			Main: debug.Module{
				Version: "v1.2.3+dirty",
			},
			Settings: []debug.BuildSetting{
				{Key: "vcs.revision", Value: "abc123"},
				{Key: "vcs.time", Value: "2026-04-25T10:00:00Z"},
			},
		}, true
	})

	if info.gitTag != "v1.2.3+dirty" {
		t.Fatalf("gitTag = %q, want %q", info.gitTag, "v1.2.3+dirty")
	}
	if info.buildTime != "2026-04-25T10:00:00Z" {
		t.Fatalf("buildTime = %q, want %q", info.buildTime, "2026-04-25T10:00:00Z")
	}
}

func TestResolveVersionInfoFallsBackToBuildInfo(t *testing.T) {
	info := resolveVersionInfo("", "", func() (*debug.BuildInfo, bool) {
		return &debug.BuildInfo{
			Settings: []debug.BuildSetting{
				{Key: "vcs.revision", Value: "abc123"},
				{Key: "vcs.time", Value: "2026-04-25T10:00:00Z"},
			},
		}, true
	})

	if info.gitTag != "abc123" {
		t.Fatalf("gitTag = %q, want %q", info.gitTag, "abc123")
	}
	if info.buildTime != "2026-04-25T10:00:00Z" {
		t.Fatalf("buildTime = %q, want %q", info.buildTime, "2026-04-25T10:00:00Z")
	}
}

func TestResolveVersionInfoSkipsDevelMainVersion(t *testing.T) {
	info := resolveVersionInfo("", "", func() (*debug.BuildInfo, bool) {
		return &debug.BuildInfo{
			Main: debug.Module{
				Version: "(devel)",
			},
			Settings: []debug.BuildSetting{
				{Key: "vcs.revision", Value: "abc123"},
				{Key: "vcs.time", Value: "2026-04-25T10:00:00Z"},
			},
		}, true
	})

	if info.gitTag != "abc123" {
		t.Fatalf("gitTag = %q, want %q", info.gitTag, "abc123")
	}
}

func TestResolveVersionInfoUsesUnknownWhenNoMetadata(t *testing.T) {
	info := resolveVersionInfo("", "", func() (*debug.BuildInfo, bool) {
		return nil, false
	})

	if info.gitTag != unknown {
		t.Fatalf("gitTag = %q, want %q", info.gitTag, unknown)
	}
	if info.buildTime != unknown {
		t.Fatalf("buildTime = %q, want %q", info.buildTime, unknown)
	}
}
