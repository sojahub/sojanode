package keeper_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	sojaapp "github.com/Sojahub/sojanode/app"
	clptest "github.com/Sojahub/sojanode/x/clp/test"
	clptypes "github.com/Sojahub/sojanode/x/clp/types"
	marginkeeper "github.com/Sojahub/sojanode/x/margin/keeper"
	"github.com/Sojahub/sojanode/x/margin/test"
	"github.com/Sojahub/sojanode/x/margin/types"
	tokenregistrytypes "github.com/Sojahub/sojanode/x/tokenregistry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeeper_Errors(t *testing.T) {
	_, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	require.NotNil(t, marginKeeper)
}

func TestKeeper_SetMTP(t *testing.T) {
	table := []struct {
		name      string
		mtp       types.MTP
		err       error
		errString error
	}{
		{
			name:      "missed defining asset",
			mtp:       types.MTP{},
			errString: errors.New("no asset specified: mtp invalid"),
		},
		{
			name:      "define asset but no address",
			mtp:       types.MTP{CollateralAsset: "xxx"},
			errString: errors.New("no address specified: mtp invalid"),
		},
		{
			name:      "define asset and address but no position",
			mtp:       types.MTP{CollateralAsset: "xxx", Address: "xxx"},
			errString: errors.New("no position specified: mtp invalid"),
		},
		{
			name: "define asset and address with long position",
			mtp:  types.MTP{CollateralAsset: "xxx", Address: "xxx", Position: types.Position_LONG},
		},
		{
			name: "define asset and address with short position",
			mtp:  types.MTP{CollateralAsset: "xxx", Address: "xxx", Position: types.Position_SHORT},
		},
	}

	for _, tt := range table {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, _, marginKeeper := initKeeper(t)
			got := marginKeeper.SetMTP(ctx, &tt.mtp)

			if tt.errString != nil {
				require.EqualError(t, got, tt.errString.Error())
			} else if tt.err == nil {
				require.NoError(t, got)
			} else {
				require.ErrorIs(t, got, tt.err)
			}
		})
	}
}

func TestKeeper_GetMTP(t *testing.T) {
	t.Run("get MTP from a store key that exists", func(t *testing.T) {
		ctx, app, marginKeeper := initKeeper(t)
		want := addMTPKey(t, ctx, app, marginKeeper, "ceth", "xxx", "xxx", types.Position_LONG, 1, sdk.NewDec(20))
		got, err := marginKeeper.GetMTP(ctx, want.Address, 1)

		fmt.Println(want)
		fmt.Println(got)

		require.NoError(t, err)
		require.Equal(t, got, want)
	})

	t.Run("fails when store keys does not exist", func(t *testing.T) {
		ctx, _, marginKeeper := initKeeper(t)
		_, got := marginKeeper.GetMTP(ctx, "xxx", 0)

		require.ErrorIs(t, got, types.ErrMTPDoesNotExist)
	})
}

func TestKeeper_GetMTPIterator(t *testing.T) {
	ctx, app, marginKeeper := initKeeper(t)
	want := addMTPKey(t, ctx, app, marginKeeper, "ceth", "xxx", "xxx", types.Position_LONG, 1, sdk.NewDec(20))
	iterator := marginKeeper.GetMTPIterator(ctx)
	bytesValue := iterator.Value()
	var got types.MTP
	types.ModuleCdc.MustUnmarshal(bytesValue, &got)

	require.Equal(t, got, want)
}

func TestKeeper_GetMTPs(t *testing.T) {
	ctx, app, marginKeeper := initKeeper(t)
	key1 := addMTPKey(t, ctx, app, marginKeeper, "ceth", "xxx", "key1", types.Position_LONG, 1, sdk.NewDec(20))
	key2 := addMTPKey(t, ctx, app, marginKeeper, "ceth", "xxx", "key2", types.Position_LONG, 1, sdk.NewDec(20))
	want := []*types.MTP{&key1, &key2}
	got, _, err := marginKeeper.GetMTPs(ctx, nil)
	require.NoError(t, err)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestKeeper_DestroyMTP(t *testing.T) {
	t.Run("key does not exist", func(t *testing.T) {
		ctx, _, marginKeeper := initKeeper(t)
		got := marginKeeper.DestroyMTP(ctx, "xxx", 1)

		require.ErrorIs(t, got, types.ErrMTPDoesNotExist)
	})
	t.Run("key exists", func(t *testing.T) {
		ctx, app, marginKeeper := initKeeper(t)
		mtp := addMTPKey(t, ctx, app, marginKeeper, "ceth", "xxx", "xxx", types.Position_LONG, 1, sdk.NewDec(20))
		got := marginKeeper.DestroyMTP(ctx, mtp.Address, 1)

		require.NoError(t, got)
	})
}

func TestKeeper_ClpKeeper(t *testing.T) {
	_, _, marginKeeper := initKeeper(t)
	clp := marginKeeper.ClpKeeper()
	require.NotNil(t, clp)
}

func TestKeeper_BankKeeper(t *testing.T) {
	_, _, marginKeeper := initKeeper(t)
	bank := marginKeeper.BankKeeper()
	require.NotNil(t, bank)
}

func TestKeeper_CLPSwap(t *testing.T) {
	asset := clptypes.Asset{Symbol: "fury"}
	pool := clptypes.Pool{
		ExternalAsset:        &asset,
		NativeAssetBalance:   sdk.NewUint(1000000000),
		NativeLiabilities:    sdk.NewUint(1000000000),
		ExternalCustody:      sdk.NewUint(1000000000),
		ExternalAssetBalance: sdk.NewUint(1000000000),
		ExternalLiabilities:  sdk.NewUint(1000000000),
		NativeCustody:        sdk.NewUint(1000000000),
		PoolUnits:            sdk.NewUint(1),
		Health:               sdk.NewDec(1),
	}

	custodySwapTests := []struct {
		name       string
		denom      string
		decimals   int64
		to         string
		sentAmount sdk.Uint
		err        error
	}{
		{
			name:       "denom not registered does not throw an error as it does not rely on tokenregistry",
			denom:      "unregistred_denom",
			decimals:   18,
			to:         "xxx",
			sentAmount: sdk.NewUint(100),
		},
		{
			name:       "invalid sent amount",
			denom:      "fury",
			decimals:   18,
			to:         "xxx",
			sentAmount: sdk.NewUint(0),
			err:        clptypes.ErrAmountTooLow,
		},
		{
			name:       "no token adjustment and non-fury target asset does not throw an error as swap result under the available pool balance",
			denom:      "fury",
			decimals:   18,
			to:         "xxx",
			sentAmount: sdk.NewUint(1000),
		},
		{
			name:       "no token adjustment and fury target asset does not throw an error as swap result under the available pool balance",
			denom:      "fury",
			decimals:   18,
			to:         "fury",
			sentAmount: sdk.NewUint(1000),
		},
		{
			name:       "token adjustment and non-fury target asset does not throw an error as swap result under the available pool balance",
			denom:      "fury",
			decimals:   9,
			to:         "xxx",
			sentAmount: sdk.NewUint(1000),
		},
		{
			name:       "token adjustment and fury target asset does not throw an error as swap result under the available pool balance",
			denom:      "fury",
			decimals:   9,
			to:         "fury",
			sentAmount: sdk.NewUint(1000),
		},
	}

	for _, tt := range custodySwapTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, app, marginKeeper := initKeeper(t)

			app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
				Denom:       tt.denom,
				Decimals:    tt.decimals,
				Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			})

			_, got := marginKeeper.CLPSwap(ctx, tt.sentAmount, tt.to, pool)

			if tt.err == nil {
				require.NoError(t, got)
			} else {
				require.ErrorIs(t, got, tt.err)
			}
		})
	}
}

func TestKeeper_Borrow(t *testing.T) {
	asset := clptypes.Asset{Symbol: "fury"}
	pool := clptypes.Pool{
		ExternalAsset:        &asset,
		NativeAssetBalance:   sdk.NewUint(1000000000),
		NativeLiabilities:    sdk.NewUint(1000000000),
		ExternalCustody:      sdk.NewUint(1000000000),
		ExternalAssetBalance: sdk.NewUint(1000000000),
		ExternalLiabilities:  sdk.NewUint(1000000000),
		NativeCustody:        sdk.NewUint(1000000000),
		PoolUnits:            sdk.NewUint(1),
		Health:               sdk.NewDec(1),
	}

	borrowTests := []struct {
		name             string
		denom            string
		decimals         int64
		to               string
		address          string
		collateralAmount sdk.Uint
		custodyAmount    sdk.Uint
		leverage         sdk.Dec
		health           sdk.Dec
		fundedAccount    bool
		err              error
		errString        error
	}{
		{
			name:             "wrong address",
			denom:            "unregistered_denom",
			decimals:         18,
			to:               "fury",
			address:          "xxx",
			collateralAmount: sdk.NewUint(10000),
			custodyAmount:    sdk.NewUint(1000),
			leverage:         sdk.NewDec(1),
			health:           sdk.NewDec(1),
			errString:        errors.New("decoding bech32 failed: invalid bech32 string length 3"),
		},
		{
			name:             "not enough fund",
			denom:            "unregistered_denom",
			decimals:         18,
			to:               "fury",
			address:          "did:fury:s1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v",
			collateralAmount: sdk.NewUint(10000),
			custodyAmount:    sdk.NewUint(1000),
			leverage:         sdk.NewDec(1),
			health:           sdk.NewDec(1),
			errString:        errors.New("user does not have enough balance of the required coin"),
		},
		{
			name:             "denom not registered does not throw error as it does not rely on token registry",
			denom:            "unregistered_denom",
			decimals:         18,
			to:               "fury",
			address:          "did:fury:s1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v",
			collateralAmount: sdk.NewUint(10000),
			custodyAmount:    sdk.NewUint(1000),
			leverage:         sdk.NewDec(1),
			health:           sdk.NewDec(1),
			fundedAccount:    true,
		},
		{
			name:             "not enough balance required to swap",
			denom:            "fury",
			decimals:         18,
			to:               "fury",
			address:          "did:fury:s1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v",
			collateralAmount: sdk.NewUint(1000000000000000),
			custodyAmount:    sdk.NewUint(1000000000000000),
			leverage:         sdk.NewDec(1),
			health:           sdk.NewDec(1),
			errString:        errors.New("user does not have enough balance of the required coin"),
		},
		{
			name:             "invalid address",
			denom:            "fury",
			decimals:         9,
			to:               "xxx",
			address:          "xxx",
			collateralAmount: sdk.NewUint(10000),
			custodyAmount:    sdk.NewUint(1000),
			leverage:         sdk.NewDec(1),
			health:           sdk.NewDec(1),
			errString:        errors.New("decoding bech32 failed: invalid bech32 string length 3"),
		},
		{
			name:             "insufficient funds",
			denom:            "fury",
			decimals:         9,
			to:               "xxx",
			address:          "did:fury:s1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v",
			collateralAmount: sdk.NewUint(10000),
			custodyAmount:    sdk.NewUint(1000),
			leverage:         sdk.NewDec(1),
			health:           sdk.NewDec(1),
			errString:        errors.New("user does not have enough balance of the required coin"),
		},
		{
			name:             "account funded",
			denom:            "fury",
			decimals:         9,
			to:               "fury",
			address:          "did:fury:s1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v",
			collateralAmount: sdk.NewUint(1000),
			custodyAmount:    sdk.NewUint(1000),
			leverage:         sdk.NewDec(1),
			health:           sdk.NewDec(1),
			fundedAccount:    true,
			err:              nil,
		},
	}

	for _, tt := range borrowTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, app, marginKeeper := initKeeper(t)

			app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
				Denom:       tt.denom,
				Decimals:    tt.decimals,
				Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			})

			var address string

			if tt.fundedAccount {
				signer := clptest.GenerateAddress(clptest.AddressKey1)
				address = signer.String()
				nativeCoin := sdk.NewCoin(clptypes.NativeSymbol, sdk.Int(sdk.NewUintFromString("10000")))
				err := sojaapp.AddCoinsToAccount(types.ModuleName, app.BankKeeper, ctx, signer, sdk.NewCoins(nativeCoin))
				require.Nil(t, err)
			} else {
				address = tt.address
			}

			mtp := addMTPKey(t, ctx, app, marginKeeper, tt.to, "xxx", address, types.Position_LONG, 1, sdk.NewDec(20))

			got := marginKeeper.Borrow(ctx, tt.to, tt.collateralAmount, tt.custodyAmount, &mtp, &pool, tt.leverage)

			if tt.errString != nil {
				require.EqualError(t, got, tt.errString.Error())
			} else if tt.err == nil {
				require.NoError(t, got)
			} else {
				require.ErrorIs(t, got, tt.err)
			}
		})
	}
}

func TestKeeper_UpdatePoolHealth(t *testing.T) {
	asset := clptypes.Asset{Symbol: "fury"}
	pool := clptypes.Pool{
		ExternalAsset:                &asset,
		NativeAssetBalance:           sdk.NewUint(1000000000),
		NativeLiabilities:            sdk.NewUint(1000000000),
		ExternalCustody:              sdk.NewUint(1000000000),
		ExternalAssetBalance:         sdk.NewUint(1000000000),
		ExternalLiabilities:          sdk.NewUint(1000000000),
		UnsettledExternalLiabilities: sdk.ZeroUint(),
		UnsettledNativeLiabilities:   sdk.ZeroUint(),
		BlockInterestExternal:        sdk.ZeroUint(),
		BlockInterestNative:          sdk.ZeroUint(),
		NativeCustody:                sdk.NewUint(1000000000),
		PoolUnits:                    sdk.NewUint(1),
		Health:                       sdk.NewDec(1),
	}

	ctx, _, marginKeeper := initKeeper(t)

	err := marginKeeper.UpdatePoolHealth(ctx, &pool)
	require.Nil(t, err)
}

func TestKeeper_UpdateMTPHealth(t *testing.T) {
	asset := clptypes.Asset{Symbol: "fury"}
	pool := clptypes.Pool{
		ExternalAsset:        &asset,
		NativeAssetBalance:   sdk.NewUint(1000000000),
		NativeLiabilities:    sdk.NewUint(1000000000),
		ExternalCustody:      sdk.NewUint(1000000000),
		ExternalAssetBalance: sdk.NewUint(1000000000),
		ExternalLiabilities:  sdk.NewUint(1000000000),
		NativeCustody:        sdk.NewUint(1000000000),
		PoolUnits:            sdk.NewUint(1),
		Health:               sdk.NewDec(1),
	}

	updateMTPHealthTests := []struct {
		name                     string
		denom                    string
		decimals                 int64
		to                       string
		position                 types.Position
		collateralAmount         sdk.Uint
		custodyAmount            sdk.Uint
		liabilities              sdk.Uint
		interestUnpaidCollateral sdk.Uint
		health                   sdk.Dec
		err                      error
		errString                error
		err2                     error
		errString2               error
	}{
		{
			name:                     "denom not registered does not throw an error as it does rely on token registry",
			denom:                    "unregistred_denom",
			decimals:                 18,
			to:                       "xxx",
			collateralAmount:         sdk.NewUint(1000),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(1000),
			interestUnpaidCollateral: sdk.NewUint(1000),
			health:                   sdk.NewDec(1),
			position:                 types.Position_LONG,
		},
		//{
		//	name:                     "not enough received asset tokens to swap does not throw an error as swap result readjusted based on slippage",
		//	denom:                    "fury",
		//	decimals:                 18,
		//	to:                       "fury",
		//	collateralAmount:         sdk.NewUint(1000),
		//	custodyAmount:            sdk.NewUint(10000000000),
		//	liabilities:              sdk.NewUint(1000),
		//	interestUnpaidCollateral: sdk.NewUint(1000),
		//	health:                   sdk.NewDec(1),
		//	position:                 types.Position_LONG,
		//},
		{
			name:                     "swap with same asset",
			denom:                    "fury",
			decimals:                 18,
			to:                       "fury",
			collateralAmount:         sdk.NewUint(1000),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(1000),
			interestUnpaidCollateral: sdk.NewUint(1000),
			health:                   sdk.NewDec(1),
			position:                 types.Position_LONG,
		},
		{
			name:                     "swap with different asset",
			denom:                    "fury",
			decimals:                 9,
			to:                       "xxx",
			collateralAmount:         sdk.NewUint(1000),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(1000),
			interestUnpaidCollateral: sdk.NewUint(1000),
			health:                   sdk.NewDec(1),
			position:                 types.Position_LONG,
		},
		{
			name:                     "insufficient liabilities funds does not throw an error as swap result readjusted based on slippage",
			denom:                    "fury",
			decimals:                 18,
			to:                       "xxx",
			collateralAmount:         sdk.NewUint(1000),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(10000000000000),
			interestUnpaidCollateral: sdk.NewUint(1000),
			health:                   sdk.NewDec(1),
			position:                 types.Position_LONG,
		},
		{
			name:                     "mtp invalid",
			denom:                    "fury",
			decimals:                 18,
			to:                       "xxx",
			collateralAmount:         sdk.NewUint(0),
			custodyAmount:            sdk.NewUint(0),
			liabilities:              sdk.NewUint(0),
			interestUnpaidCollateral: sdk.NewUint(0),
			health:                   sdk.NewDec(1),
			err2:                     types.ErrMTPInvalid,
			position:                 types.Position_UNSPECIFIED,
		},
	}

	for _, tt := range updateMTPHealthTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, app, marginKeeper := initKeeper(t)

			app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
				Denom:       tt.denom,
				Decimals:    tt.decimals,
				Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			})

			mtp := addMTPKey(t, ctx, app, marginKeeper, tt.to, "xxx", "xxx", tt.position, 1, sdk.NewDec(20))
			mtp.CustodyAmount = tt.custodyAmount
			mtp.Liabilities = tt.liabilities
			mtp.CollateralAmount = tt.collateralAmount
			mtp.InterestUnpaidCollateral = tt.interestUnpaidCollateral

			_, got := marginKeeper.UpdateMTPHealth(ctx, mtp, pool)

			if tt.errString != nil {
				require.EqualError(t, got, tt.errString.Error())
			} else if tt.err == nil {
				require.NoError(t, got)
			} else {
				require.ErrorIs(t, got, tt.err)
			}

			got = mtp.Validate()
			if tt.errString2 != nil {
				require.EqualError(t, got, tt.errString2.Error())
			} else if tt.err2 == nil {
				require.NoError(t, got)
			} else {
				require.ErrorIs(t, got, tt.err2)
			}
		})
	}
}

func TestKeeper_TakeInCustody(t *testing.T) {
	asset := clptypes.Asset{Symbol: "fury"}

	t.Run("settlement asset and mtp asset is equal", func(t *testing.T) {
		ctx, app, marginKeeper := initKeeper(t)
		mtp := addMTPKey(t, ctx, app, marginKeeper, "fury", "xxx", "xxx", types.Position_LONG, 1, sdk.NewDec(20))

		pool := clptypes.Pool{
			ExternalAsset:        &asset,
			NativeAssetBalance:   sdk.NewUint(1000),
			NativeLiabilities:    sdk.NewUint(1000),
			ExternalCustody:      sdk.NewUint(1000),
			ExternalAssetBalance: sdk.NewUint(1000),
			ExternalLiabilities:  sdk.NewUint(1000),
			NativeCustody:        sdk.NewUint(1000),
			PoolUnits:            sdk.NewUint(1),
			Health:               sdk.NewDec(1),
		}

		got := marginKeeper.TakeInCustody(ctx, mtp, &pool)

		require.NoError(t, got)
	})

	t.Run("settlement asset and mtp asset is not equal", func(t *testing.T) {
		ctx, app, marginKeeper := initKeeper(t)
		mtp := addMTPKey(t, ctx, app, marginKeeper, "notfury", "xxx", "xxx", types.Position_LONG, 1, sdk.NewDec(20))

		pool := clptypes.Pool{
			ExternalAsset:        &asset,
			NativeAssetBalance:   sdk.NewUint(1000),
			NativeLiabilities:    sdk.NewUint(1000),
			ExternalCustody:      sdk.NewUint(1000),
			ExternalAssetBalance: sdk.NewUint(1000),
			ExternalLiabilities:  sdk.NewUint(1000),
			NativeCustody:        sdk.NewUint(1000),
			PoolUnits:            sdk.NewUint(1),
			Health:               sdk.NewDec(1),
		}

		got := marginKeeper.TakeInCustody(ctx, mtp, &pool)

		require.NoError(t, got)
	})
}

func TestKeeper_TakeOutCustody(t *testing.T) {
	asset := clptypes.Asset{Symbol: "fury"}

	t.Run("settlement asset and mtp asset is equal", func(t *testing.T) {
		ctx, app, marginKeeper := initKeeper(t)
		mtp := addMTPKey(t, ctx, app, marginKeeper, "fury", "xxx", "xxx", types.Position_LONG, 1, sdk.NewDec(20))

		pool := clptypes.Pool{
			ExternalAsset:        &asset,
			NativeAssetBalance:   sdk.NewUint(1000),
			NativeLiabilities:    sdk.NewUint(1000),
			ExternalCustody:      sdk.NewUint(1000),
			ExternalAssetBalance: sdk.NewUint(1000),
			ExternalLiabilities:  sdk.NewUint(1000),
			NativeCustody:        sdk.NewUint(1000),
			PoolUnits:            sdk.NewUint(1),
			Health:               sdk.NewDec(1),
		}

		got := marginKeeper.TakeOutCustody(ctx, mtp, &pool)

		require.NoError(t, got)
	})

	t.Run("settlement asset and mtp asset is not equal", func(t *testing.T) {
		ctx, app, marginKeeper := initKeeper(t)
		mtp := addMTPKey(t, ctx, app, marginKeeper, "notfury", "xxx", "xxx", types.Position_LONG, 1, sdk.NewDec(20))

		pool := clptypes.Pool{
			ExternalAsset:        &asset,
			NativeAssetBalance:   sdk.NewUint(1000),
			NativeLiabilities:    sdk.NewUint(1000),
			ExternalCustody:      sdk.NewUint(1000),
			ExternalAssetBalance: sdk.NewUint(1000),
			ExternalLiabilities:  sdk.NewUint(1000),
			NativeCustody:        sdk.NewUint(1000),
			PoolUnits:            sdk.NewUint(1),
			Health:               sdk.NewDec(1),
		}

		got := marginKeeper.TakeOutCustody(ctx, mtp, &pool)

		require.NoError(t, got)
	})
}

func TestKeeper_Repay(t *testing.T) {
	asset := clptypes.Asset{Symbol: "fury"}
	pool := clptypes.Pool{
		ExternalAsset:                &asset,
		NativeAssetBalance:           sdk.NewUint(1000000000),
		NativeLiabilities:            sdk.NewUint(1000000000),
		ExternalCustody:              sdk.NewUint(1000000000),
		ExternalAssetBalance:         sdk.NewUint(1000000000),
		ExternalLiabilities:          sdk.NewUint(1000000000),
		NativeCustody:                sdk.NewUint(1000000000),
		UnsettledExternalLiabilities: sdk.ZeroUint(),
		UnsettledNativeLiabilities:   sdk.ZeroUint(),
		BlockInterestExternal:        sdk.ZeroUint(),
		BlockInterestNative:          sdk.ZeroUint(),
		PoolUnits:                    sdk.NewUint(1),
		Health:                       sdk.NewDec(1),
	}

	repayTests := []struct {
		name                     string
		denom                    string
		decimals                 int64
		to                       string
		address                  string
		position                 types.Position
		collateralAmount         sdk.Uint
		custodyAmount            sdk.Uint
		liabilities              sdk.Uint
		interestUnpaidCollateral sdk.Uint
		health                   sdk.Dec
		repayAmount              sdk.Uint
		overrideAddress          string
		err                      error
		errString                error
		err2                     error
		errString2               error
	}{
		{
			name:                     "denom not registered does not throw an error as it does rely on token registry",
			denom:                    "unregistred_denom",
			decimals:                 18,
			to:                       "xxx",
			address:                  "xxx",
			collateralAmount:         sdk.NewUint(1000),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(1000),
			interestUnpaidCollateral: sdk.NewUint(1000),
			health:                   sdk.NewDec(1),
			repayAmount:              sdk.NewUint(1),
			position:                 types.Position_LONG,
		},
		{
			name:                     "cannot afford principle liability",
			denom:                    "fury",
			decimals:                 18,
			to:                       "xxx",
			address:                  "xxx",
			collateralAmount:         sdk.NewUint(0),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(1000),
			interestUnpaidCollateral: sdk.NewUint(1000),
			health:                   sdk.NewDec(1),
			repayAmount:              sdk.NewUint(0),
			position:                 types.Position_LONG,
		},
		{
			name:                     "v principle libarity; x excess liability",
			denom:                    "fury",
			decimals:                 18,
			to:                       "xxx",
			address:                  "xxx",
			collateralAmount:         sdk.NewUint(0),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(0),
			interestUnpaidCollateral: sdk.NewUint(1000),
			health:                   sdk.NewDec(1),
			repayAmount:              sdk.NewUint(0),
			position:                 types.Position_LONG,
		},
		{
			name:                     "can afford both",
			denom:                    "fury",
			decimals:                 18,
			to:                       "xxx",
			address:                  "xxx",
			collateralAmount:         sdk.NewUint(0),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(0),
			interestUnpaidCollateral: sdk.NewUint(0),
			health:                   sdk.NewDec(1),
			repayAmount:              sdk.NewUint(0),
			position:                 types.Position_LONG,
		},
		{
			name:                     "non zero return amount + fails because of wrong address",
			denom:                    "fury",
			decimals:                 18,
			to:                       "xxx",
			address:                  "xxx",
			collateralAmount:         sdk.NewUint(0),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(0),
			interestUnpaidCollateral: sdk.NewUint(0),
			health:                   sdk.NewDec(1),
			repayAmount:              sdk.NewUint(1000),
			position:                 types.Position_LONG,
			errString:                errors.New("decoding bech32 failed: invalid bech32 string length 3"),
		},
		{
			name:                     "non zero return amount",
			denom:                    "fury",
			decimals:                 18,
			to:                       "xxx",
			address:                  "did:fury:s1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v",
			collateralAmount:         sdk.NewUint(0),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(0),
			interestUnpaidCollateral: sdk.NewUint(0),
			health:                   sdk.NewDec(1),
			repayAmount:              sdk.NewUint(1000),
			position:                 types.Position_LONG,
			errString:                errors.New("0xxx is smaller than 1000xxx: insufficient funds"),
		},
		{
			name:                     "collateral and native assets are equal",
			denom:                    "fury",
			decimals:                 18,
			to:                       "fury",
			address:                  "xxx",
			collateralAmount:         sdk.NewUint(0),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(1000),
			interestUnpaidCollateral: sdk.NewUint(1000),
			health:                   sdk.NewDec(1),
			repayAmount:              sdk.NewUint(0),
			position:                 types.Position_LONG,
		},
		{
			name:                     "mtp not found",
			denom:                    "fury",
			decimals:                 18,
			to:                       "fury",
			address:                  "xxx",
			overrideAddress:          "yyy",
			collateralAmount:         sdk.NewUint(0),
			custodyAmount:            sdk.NewUint(1000),
			liabilities:              sdk.NewUint(1000),
			interestUnpaidCollateral: sdk.NewUint(1000),
			health:                   sdk.NewDec(1),
			repayAmount:              sdk.NewUint(0),
			position:                 types.Position_LONG,
			err:                      types.ErrMTPDoesNotExist,
		},
	}

	for _, tt := range repayTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, app, marginKeeper := initKeeper(t)

			app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
				Denom:       tt.denom,
				Decimals:    tt.decimals,
				Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			})

			mtp := addMTPKey(t, ctx, app, marginKeeper, tt.to, "xxx", tt.address, types.Position_LONG, 1, sdk.NewDec(20))
			mtp.CustodyAmount = tt.custodyAmount
			mtp.Liabilities = tt.liabilities
			mtp.CollateralAmount = tt.collateralAmount
			mtp.InterestUnpaidCollateral = tt.interestUnpaidCollateral
			if tt.overrideAddress != "" {
				mtp.Address = tt.overrideAddress
			}

			got := marginKeeper.Repay(ctx, &mtp, &pool, tt.repayAmount, false)

			if tt.errString != nil {
				require.EqualError(t, got, tt.errString.Error())
			} else if tt.err == nil {
				require.NoError(t, got)
			} else {
				require.ErrorIs(t, got, tt.err)
			}

			got = mtp.Validate()
			if tt.errString2 != nil {
				require.EqualError(t, got, tt.errString2.Error())
			} else if tt.err2 == nil {
				require.NoError(t, got)
			} else {
				require.ErrorIs(t, got, tt.err2)
			}
		})
	}
}

func TestKeeper_CheckMinLiabilities(t *testing.T) {
	ctx, _, marginKeeper := initKeeper(t)
	params := marginKeeper.GetParams(ctx)
	params.InterestRateMin = sdk.MustNewDecFromStr("0.00000001")
	marginKeeper.SetParams(ctx, &params)
	pool := clptypes.Pool{
		ExternalAsset:                &clptypes.Asset{Symbol: "cusdc"},
		NativeAssetBalance:           sdk.NewUintFromString("1000000000000000000000"),
		NativeLiabilities:            sdk.NewUint(0),
		ExternalCustody:              sdk.NewUint(0),
		ExternalAssetBalance:         sdk.NewUint(1000000000),
		ExternalLiabilities:          sdk.NewUint(0),
		NativeCustody:                sdk.NewUint(0),
		UnsettledExternalLiabilities: sdk.ZeroUint(),
		UnsettledNativeLiabilities:   sdk.ZeroUint(),
		BlockInterestExternal:        sdk.ZeroUint(),
		BlockInterestNative:          sdk.ZeroUint(),
		PoolUnits:                    sdk.NewUint(1),
		Health:                       sdk.NewDec(1),
		InterestRate:                 sdk.NewDec(1),
	}
	got := marginKeeper.CheckMinLiabilities(ctx, sdk.NewUint(200000000), sdk.OneDec(), pool, "fury")
	require.Nil(t, got)

	got = marginKeeper.CheckMinLiabilities(ctx, sdk.NewUint(10000000), sdk.OneDec(), pool, "fury")
	require.EqualError(t, got, "borrowed amount is too low")

	got = marginKeeper.CheckMinLiabilities(ctx, sdk.NewUint(20000000), sdk.NewDec(9), pool, "fury")
	require.Nil(t, got)

	got = marginKeeper.CheckMinLiabilities(ctx, sdk.NewUint(2000000), sdk.NewDec(9), pool, "fury")
	require.EqualError(t, got, "borrowed amount is too low")
}

func TestKeeper_CalcMTPInterestLiabilities(t *testing.T) {
	ctx, app, marginKeeper := initKeeper(t)

	mtp := addMTPKey(t, ctx, app, marginKeeper, "fury", "xxx", "xxx", types.Position_LONG, 1, sdk.NewDec(20))
	// calculation on epoch
	got := marginkeeper.CalcMTPInterestLiabilities(&mtp, sdk.NewDecWithPrec(1, 1), 0, 1)
	require.Equal(t, sdk.NewUint(1200), got)
	// calculation within epoch
	got = marginkeeper.CalcMTPInterestLiabilities(&mtp, sdk.NewDecWithPrec(1, 1), 3, 10)
	require.Equal(t, sdk.NewUint(1060), got)
}

func TestKeeper_CalcMTPInterestLiabilitiesOverflow(t *testing.T) { // test fails after fix to interest calc

	mtp := types.MTP{
		Address:                  "did:fury:s123",
		CollateralAsset:          "fury",
		CollateralAmount:         sdk.Uint{},
		Liabilities:              sdk.NewUintFromString("100"),
		InterestPaidCollateral:   sdk.ZeroUint(),
		InterestPaidCustody:      sdk.ZeroUint(),
		InterestUnpaidCollateral: sdk.NewUintFromString("45231284858326638837332416019018714005183587760015845327913118753091066265500"),
		CustodyAsset:             "ceth",
		CustodyAmount:            sdk.Uint{},
		Leverage:                 sdk.Dec{},
		MtpHealth:                sdk.Dec{},
		Position:                 types.Position_LONG,
		Id:                       1,
	}

	require.Panics(t, func() {
		marginkeeper.CalcMTPInterestLiabilities(&mtp, sdk.NewDec(3.0), 0, 1)
	}, "the code did not panic")
}

func TestKeeper_InterestRateComputation(t *testing.T) {
	interestRateComputationTests := []struct {
		name                 string
		denom                string
		decimals             int64
		interestRate         sdk.Dec
		interestRateIncrease sdk.Dec
		interestRateDecrease sdk.Dec
		interestRateMax      sdk.Dec
		err                  error
		errString            error
	}{
		{
			name:                 "interest rate change lesser than decrease and increase",
			denom:                "unregistred_denom",
			decimals:             18,
			interestRate:         sdk.NewDec(1),
			interestRateIncrease: sdk.NewDec(1),
			interestRateDecrease: sdk.NewDec(1),
			interestRateMax:      sdk.NewDec(5),
			err:                  nil,
		},
		{
			name:                 "interest rate change greater than increase",
			denom:                "unregistred_denom",
			decimals:             18,
			interestRate:         sdk.NewDec(10),
			interestRateIncrease: sdk.NewDec(1),
			interestRateDecrease: sdk.NewDec(1),
			interestRateMax:      sdk.NewDec(5),
			err:                  nil,
		},
		{
			name:                 "interest rate greater than rate max",
			denom:                "unregistred_denom",
			decimals:             18,
			interestRate:         sdk.NewDec(10),
			interestRateIncrease: sdk.NewDec(1),
			interestRateDecrease: sdk.NewDec(1),
			interestRateMax:      sdk.NewDec(0),
			err:                  nil,
		},
	}

	for _, tt := range interestRateComputationTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, app, marginKeeper := initKeeper(t)

			app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
				Denom:       tt.denom,
				Decimals:    tt.decimals,
				Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			})

			data := types.GenesisState{Params: &types.Params{
				LeverageMax:                              sdk.NewDec(10),
				InterestRateMax:                          tt.interestRateMax,
				InterestRateMin:                          sdk.NewDec(1),
				InterestRateIncrease:                     tt.interestRateIncrease,
				InterestRateDecrease:                     tt.interestRateDecrease,
				HealthGainFactor:                         sdk.NewDec(1),
				EpochLength:                              1,
				ForceCloseFundPercentage:                 sdk.NewDecWithPrec(1, 1),
				ForceCloseFundAddress:                    "did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777p07psd",
				IncrementalInterestPaymentFundPercentage: sdk.NewDecWithPrec(1, 1),
				IncrementalInterestPaymentFundAddress:    "did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777p07psd",
				IncrementalInterestPaymentEnabled:        false,
				PoolOpenThreshold:                        sdk.NewDecWithPrec(1, 1),
				RemovalQueueThreshold:                    sdk.NewDecWithPrec(1, 1),
				MaxOpenPositions:                         10000,
				Pools:                                    []string{},
				SqModifier:                               sdk.MustNewDecFromStr("10000000000000000000000000"),
				SafetyFactor:                             sdk.MustNewDecFromStr("1.05"),
			}}
			marginKeeper.InitGenesis(ctx, data)

			asset := clptypes.Asset{Symbol: "fury"}
			pool := clptypes.Pool{
				ExternalAsset:                &asset,
				NativeAssetBalance:           sdk.NewUint(1000000000),
				NativeLiabilities:            sdk.NewUint(1000000000),
				ExternalCustody:              sdk.NewUint(1000000000),
				ExternalAssetBalance:         sdk.NewUint(1000000000),
				ExternalLiabilities:          sdk.NewUint(1000000000),
				NativeCustody:                sdk.NewUint(1000000000),
				UnsettledExternalLiabilities: sdk.ZeroUint(),
				UnsettledNativeLiabilities:   sdk.ZeroUint(),
				BlockInterestExternal:        sdk.ZeroUint(),
				BlockInterestNative:          sdk.ZeroUint(),
				PoolUnits:                    sdk.NewUint(1),
				Health:                       sdk.NewDec(1),
				InterestRate:                 tt.interestRate,
			}

			_, got := marginKeeper.InterestRateComputation(ctx, pool)

			t.Logf("got %v", got)

			if tt.errString != nil {
				require.EqualError(t, got, tt.errString.Error())
			} else if tt.err == nil {
				require.NoError(t, got)
			} else {
				require.ErrorIs(t, got, tt.err)
			}
		})
	}
}

func TestWhitelist(t *testing.T) {
	ctx, _, marginKeeper := initKeeper(t)

	marginKeeper.WhitelistAddress(ctx, "did:fury:s123")
	is := marginKeeper.IsWhitelisted(ctx, "did:fury:s123")
	require.True(t, is)
	whitelist, _, err := marginKeeper.GetWhitelist(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, []string{"did:fury:s123"}, whitelist)
}

func TestSQBeginBlocker(t *testing.T) {
	ctx, _, marginKeeper := initKeeper(t)
	params := marginKeeper.GetParams(ctx)
	params.RemovalQueueThreshold = sdk.NewDec(1)
	marginKeeper.SetParams(ctx, &params)
	pool := clptypes.Pool{
		ExternalAsset: &clptypes.Asset{Symbol: "ceth"},
		Health:        sdk.NewDec(2),
	}

	marginKeeper.TrackSQBeginBlock(ctx, &pool)
	require.Equal(t, uint64(0), marginKeeper.GetSQBeginBlock(ctx, &pool))

	pool.Health = sdk.NewDecWithPrec(1, 1)
	marginKeeper.TrackSQBeginBlock(ctx.WithBlockHeight(1), &pool)
	require.Equal(t, uint64(1), marginKeeper.GetSQBeginBlock(ctx, &pool))

	pool.Health = sdk.NewDec(2)
	marginKeeper.TrackSQBeginBlock(ctx.WithBlockHeight(2), &pool)
	require.Equal(t, uint64(0), marginKeeper.GetSQBeginBlock(ctx, &pool))
}

func initKeeper(t testing.TB) (sdk.Context, *sojaapp.SojahubApp, types.Keeper) {
	ctx, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	require.NotNil(t, marginKeeper)
	return ctx, app, marginKeeper
}
func addMTPKey(t testing.TB, ctx sdk.Context, app *sojaapp.SojahubApp, marginKeeper types.Keeper, collateralAsset string, custodyAsset string, address string, position types.Position, id uint64, health sdk.Dec) types.MTP {
	storeKey := app.GetKey(types.StoreKey)
	store := ctx.KVStore(storeKey)
	key := types.GetMTPKey(address, id)

	newMTP := types.MTP{
		Id:                       id,
		Address:                  address,
		CollateralAsset:          collateralAsset,
		Liabilities:              sdk.NewUint(1000),
		InterestPaidCollateral:   sdk.ZeroUint(),
		InterestPaidCustody:      sdk.ZeroUint(),
		InterestUnpaidCollateral: sdk.NewUint(1000),
		CollateralAmount:         sdk.NewUint(1000),
		CustodyAsset:             custodyAsset,
		CustodyAmount:            sdk.NewUint(1000),
		Leverage:                 sdk.NewDec(10),
		MtpHealth:                health,
		Position:                 position,
	}

	store.Set(key, types.ModuleCdc.MustMarshal(&newMTP))

	return newMTP
}
