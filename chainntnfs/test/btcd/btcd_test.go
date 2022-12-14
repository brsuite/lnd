//go:build dev
// +build dev

package brond_test

import (
	"testing"

	chainntnfstest "github.com/brsuite/lnd/chainntnfs/test"
)

// TestInterfaces executes the generic notifier test suite against a brond
// powered chain notifier.
func TestInterfaces(t *testing.T) {
	chainntnfstest.TestInterfaces(t, "brond")
}
