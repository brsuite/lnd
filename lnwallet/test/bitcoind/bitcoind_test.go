package brocoind_test

import (
	"testing"

	lnwallettest "github.com/brsuite/lnd/lnwallet/test"
)

// TestLightningWallet tests LightningWallet powered by brocoind against our
// suite of interface tests.
func TestLightningWallet(t *testing.T) {
	lnwallettest.TestLightningWallet(t, "brocoind")
}
