package test

import (
	sojaapp "github.com/Sojahub/sojanode/app"
	admintypes "github.com/Sojahub/sojanode/x/admin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func CreateTestApp(isCheckTx bool) (*sojaapp.SojahubApp, sdk.Context, string) {
	sojaapp.SetConfig(false)
	app := sojaapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	initTokens := sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction)
	_ = sojaapp.AddTestAddrs(app, ctx, 6, initTokens)
	admin := sdk.AccAddress("addr1_______________")
	app.AdminKeeper.InitGenesis(ctx, admintypes.GenesisState{AdminAccounts: GetAdmins(admin.String())})
	return app, ctx, admin.String()
}

func GetAdmins(address string) []*admintypes.AdminAccount {
	return []*admintypes.AdminAccount{
		{
			AdminType:    admintypes.AdminType_TOKENREGISTRY,
			AdminAddress: address,
		},
	}
}
