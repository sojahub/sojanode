//go:build TEST_INTEGRATION
// +build TEST_INTEGRATION

package testutil

import (
	"os"
	"testing"

	sojaapp "github.com/Sojahub/sojanode/app"
	"github.com/cosmos/cosmos-sdk/baseapp"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	encConfig := sojaapp.MakeTestEncodingConfig()
	cfg.InterfaceRegistry = encConfig.InterfaceRegistry
	cfg.Codec = encConfig.Marshaler
	cfg.TxConfig = encConfig.TxConfig
	cfg.AppConstructor = func(val network.Validator) servertypes.Application {
		return sojaapp.NewSojaApp(
			log.NewTMLogger(log.NewSyncWriter(os.Stdout)),
			dbm.NewMemDB(),
			nil,
			true,
			make(map[int64]bool),
			val.Dir,
			0,
			encConfig,
			sojaapp.EmptyAppOptions{},
			baseapp.SetTrace(true),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
		)
	}

	suite.Run(t, NewIntegrationTestSuite(cfg))
}
