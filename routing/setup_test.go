package routing

import (
	"testing"

	"github.com/brsuite/lnd/kvdb"
)

func TestMain(m *testing.M) {
	kvdb.RunTests(m)
}
