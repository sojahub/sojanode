package keeper

import (
	clptypes "github.com/Sojahub/sojanode/x/clp/types"
	"github.com/Sojahub/sojanode/x/dispensation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

func (m Migrator) MigrateToVer2(ctx sdk.Context) error {
	m.keeper.SetMintController(ctx,
		types.MintController{TotalCounter: sdk.NewCoin(clptypes.GetSettlementAsset().Symbol, sdk.ZeroInt())})
	return nil
}
