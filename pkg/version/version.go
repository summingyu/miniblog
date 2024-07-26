// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"
)

var (
	// GitVersion 是语义化的版本号.
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate 是 ISO8601 格式的构建时间, $(date -u +'%Y-%m-%dT%H:%M:%SZ') 命令的输出.
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit 是 Git 的 SHA1 值，$(git rev-parse HEAD) 命令的输出.
	GitCommit = "$Format:%H$"
	// GitTreeState 代表构建时 Git 仓库的状态，可能的值有：clean, dirty.
	GitTreeState = ""
)

type Info struct {
	GitVersion   string `json:"gitVersion"`
	BuildDate    string `json:"buildDate"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info Info) String() string {
	if s, err := info.Text(); err == nil {
		return string(s)
	}
	return info.GitVersion
}

func (info Info) ToJSON() string {
	s, _ := json.Marshal(info)
	return string(s)
}

func (info Info) Text() ([]byte, error) {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("gitVersion:", info.GitVersion)
	table.AddRow("gitCommit:", info.GitCommit)
	table.AddRow("gitTreeState:", info.GitTreeState)
	table.AddRow("buildDate:", info.BuildDate)
	table.AddRow("goVersion:", info.GoVersion)
	table.AddRow("compiler:", info.Compiler)
	table.AddRow("platform:", info.Platform)

	return table.Bytes(), nil
}

func Get() Info {
	return Info{
		GitVersion:   GitVersion,
		BuildDate:    BuildDate,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
