package version

import (
	"fmt"
	"runtime"
)

var (
	Name       = ""
	ServerName = ""
	ClientName = ""
	Version    = ""
	Commit     = ""
	BuildTags  = ""
)

type Info struct {
	Name       string `json:"name" yaml:"name"`
	ServerName string `json:"server_name" yaml:"server_name"`
	ClientName string `json:"client_name" yaml:"client_name"`
	Version    string `json:"version" yaml:"version"`
	GitCommit  string `json:"commit" yaml:"commit"`
	BuildTags  string `json:"build_tags" yaml:"build_tags"`
	GoVersion  string `json:"go" yaml:"go"`
}

func NewInfo() Info {
	return Info{
		Name:       Name,
		ServerName: ServerName,
		ClientName: ClientName,
		Version:    Version,
		GitCommit:  Commit,
		BuildTags:  BuildTags,
		GoVersion:  fmt.Sprintf("go version %s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
	}
}

func (vi Info) String() string {
	return fmt.Sprintf(`%s: %s
git commit: %s
build tags: %s
%s`,
		vi.Name, vi.Version, vi.GitCommit, vi.BuildTags, vi.GoVersion,
	)
}
