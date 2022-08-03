package migration_01_to_11

import (
	"github.com/brsuite/bronlog"
)

// log is a logger that is initialized as disabled.  This means the package will
// not perform any logging by default until a logger is set.
var log = bronlog.Disabled

// UseLogger uses a specified Logger to output package logging info.
func UseLogger(logger bronlog.Logger) {
	log = logger
}
