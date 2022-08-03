package chainntnfs

import (
	"github.com/brsuite/bronlog"
	"github.com/brsuite/lnd/build"
)

// Log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var Log bronlog.Logger

// The default amount of logging is none.
func init() {
	UseLogger(build.NewSubLogger("NTFN", nil))
}

// DisableLog disables all library log output.  Logging output is disabled
// by default until UseLogger is called.
func DisableLog() {
	UseLogger(bronlog.Disabled)
}

// UseLogger uses a specified Logger to output package logging info.
// This should be used in preference to SetLogWriter if the caller is also
// using bronlog.
func UseLogger(logger bronlog.Logger) {
	Log = logger
}