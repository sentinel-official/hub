package version

import (
	"fmt"
	"runtime"
)

var (
	Commit        = ""
	Version       = ""
	VendorDirHash = ""
	BuildTags     = ""
)

type versionInfo struct {
	SentinelSDK   string `json:"sentinel_sdk"`
	GitCommit     string `json:"commit"`
	VendorDirHash string `json:"vendor_hash"`
	BuildTags     string `json:"build_tags"`
	GoVersion     string `json:"go"`
}

func (v versionInfo) String() string {
	return fmt.Sprintf(`sentinel-sdk: %s
git commit: %s
vendor hash: %s
build tags: %s
%s`, v.SentinelSDK, v.GitCommit, v.VendorDirHash, v.BuildTags, v.GoVersion)
}

func newVersionInfo() versionInfo {
	return versionInfo{
		SentinelSDK:   Version,
		GitCommit:     Commit,
		VendorDirHash: VendorDirHash,
		BuildTags:     BuildTags,
		GoVersion:     fmt.Sprintf("go version %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH),
	}
}
