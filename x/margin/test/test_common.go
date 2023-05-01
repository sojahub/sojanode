package test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sojaapp "github.com/Sojahub/sojanode/app"
)

func CreateTestApp(isCheckTx bool) (*sojaapp.SojahubApp, sdk.Context) {
	app := sojaapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	initTokens := sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction)
	_ = sojaapp.AddTestAddrs(app, ctx, 6, initTokens)
	return app, ctx
}

func CreateTestAppMargin(isCheckTx bool) (sdk.Context, *sojaapp.SojahubApp) {
	sojaapp.SetConfig((false))
	app := sojaapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	return ctx, app
}

func CreateTestAppMarginFromGenesis(isCheckTx bool, genesisTransformer func(*sojaapp.SojahubApp, sojaapp.GenesisState) sojaapp.GenesisState) (sdk.Context, *sojaapp.SojahubApp) {
	sojaapp.SetConfig(false)
	app := sojaapp.SetupFromGenesis(isCheckTx, genesisTransformer)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	return ctx, app
}
