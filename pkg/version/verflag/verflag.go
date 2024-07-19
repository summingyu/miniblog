package verflag

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/pflag"

	"github.com/summingyu/miniblog/internal/pkg/log"
	"github.com/summingyu/miniblog/pkg/version"
)

type versionValue int

const (
	VersionFalse versionValue = iota
	VersionTrue
	VersionRaw
)

const (
	strRawVersion   = "raw"
	versionFlagName = "version"
)

var versionFlag = Version(versionFlagName, VersionFalse, "Print version information and exit")

func (v *versionValue) IsBoolFlag() bool {
	return true
}

func (v *versionValue) Get() interface{} {
	return v
}

// String implements pflag.Value.
func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}
	return fmt.Sprintf("%v", bool(*v == VersionTrue))
}

// Type implements pflag.Value.
func (v *versionValue) Type() string {
	return "version"
}

func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}
	return err
}

// VersionVar 定义了一个具有指定名称和用法的标志
func VersionVar(p *versionValue, name string, value versionValue, usage string) {
	*p = value
	pflag.Var(p, name, usage)
	// `--version` 等价于 `--version=true`
	pflag.Lookup(name).NoOptDefVal = "true"
}

// Version 包装了 VersionVar 函数.
func Version(name string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, value, usage)
	return p
}

// AddFlags 在任意 FlagSet 上注册这个包的标志，这样它们指向与全局标志相同的值.
func AddFlags(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(versionFlagName))
}

func PrintAndExitIfRequested() {
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}
	log.Debugw(fmt.Sprintf("%#v\n", version.Get()))
}
