package channeldb

import (
	"github.com/brsuite/bronlog"
	"github.com/brsuite/lnd/build"
	mig "github.com/brsuite/lnd/channeldb/migration"
	"github.com/brsuite/lnd/channeldb/migration12"
	"github.com/brsuite/lnd/channeldb/migration13"
	"github.com/brsuite/lnd/channeldb/migration16"
	"github.com/brsuite/lnd/channeldb/migration24"
	"github.com/brsuite/lnd/channeldb/migration_01_to_11"
	"github.com/brsuite/lnd/kvdb"
)

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log bronlog.Logger

func init() {
	UseLogger(build.NewSubLogger("CHDB", nil))
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
	log = logger
	mig.UseLogger(logger)
	migration_01_to_11.UseLogger(logger)
	migration12.UseLogger(logger)
	migration13.UseLogger(logger)
	migration16.UseLogger(logger)
	migration24.UseLogger(logger)
	kvdb.UseLogger(logger)
}
