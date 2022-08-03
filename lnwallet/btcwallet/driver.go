package bronwallet

import (
	"fmt"

	"github.com/brsuite/bronwallet/chain"
	"github.com/brsuite/lnd/blockcache"
	"github.com/brsuite/lnd/lnwallet"
)

const (
	walletType = "bronwallet"
)

// createNewWallet creates a new instance of bronwallet given the proper list of
// initialization parameters. This function is the factory function required to
// properly create an instance of the lnwallet.WalletDriver struct for
// bronwallet.
func createNewWallet(args ...interface{}) (lnwallet.WalletController, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("incorrect number of arguments to .New(...), "+
			"expected 2, instead passed %v", len(args))
	}

	config, ok := args[0].(*Config)
	if !ok {
		return nil, fmt.Errorf("first argument to brondnotifier.New is " +
			"incorrect, expected a *rpcclient.ConnConfig")
	}

	blockCache, ok := args[1].(*blockcache.BlockCache)
	if !ok {
		return nil, fmt.Errorf("second argument to brondnotifier.New is " +
			"incorrect, expected a *blockcache.BlockCache")
	}

	return New(*config, blockCache)
}

// init registers a driver for the bronwallet concrete implementation of the
// lnwallet.WalletController interface.
func init() {
	// Register the driver.
	driver := &lnwallet.WalletDriver{
		WalletType: walletType,
		New:        createNewWallet,
		BackEnds:   chain.BackEnds,
	}

	if err := lnwallet.RegisterWallet(driver); err != nil {
		panic(fmt.Sprintf("failed to register wallet driver '%s': %v",
			walletType, err))
	}
}
