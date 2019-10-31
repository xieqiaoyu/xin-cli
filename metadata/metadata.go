package metadata

import (
	"fmt"
)

var (
	// Name 项目名称
	Name = "xin-cli"
	// Version 项目版本号
	Version = "Unknown"
	// 项目的平台
	Platform = "Unknown"
)

const versionString = `
xin-cli (%s) on %s

---- Four words are enough -----

`

func ShowVersion() {
	fmt.Printf(versionString, Version, Platform)
}
