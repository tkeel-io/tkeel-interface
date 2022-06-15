package version

import (
	"fmt"
	"os"
	"runtime"
)

// Version
const _Version = "VERSION"

var Version = "edge"

// GitCommit The git commit that was compiled. This will be filled in by the compiler.
var (
	GitCommit string
	GitBranch string
)

// GitVersion The main version number that is being run at the moment.
var GitVersion string

// BuildDate The build datetime at the moment.
var BuildDate = ""

// GoVersion The go compiler version.
var GoVersion = runtime.Version()

// OsArch The system info.
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)

func init() {
	if ver := os.Getenv(_Version); ver != "" {
		Version = ver
	}
}
