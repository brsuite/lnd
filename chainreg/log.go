package chainreg

import (
	"github.com/brsuite/bronlog"
	"github.com/brsuite/lnd/build"
)

// Subsystem defines the logging code for this subsystem.
const Subsystem = "CHRE"

// log is a logger that is initialized with the bronlog.Disabled logger.
var log bronlog.Logger

// The default amount of logging is none.
func init() {
	UseLogger(build.NewSubLogger(Subsystem, nil))
}

// DisableLog disables all logging output.
func DisableLog() {
	UseLogger(bronlog.Disabled)
}

// UseLogger uses a specified Logger to output package logging info.
func UseLogger(logger bronlog.Logger) {
	log = logger
}
