//nolint
package keeper_test

/* TODO update this for 42 branch

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/Sojahub/sojanode/x/dispensation/keeper"
	"github.com/Sojahub/sojanode/x/dispensation/test"
	"github.com/Sojahub/sojanode/x/dispensation/types"
	"github.com/Sojahub/sojanode/x/dispensation/types/legacy"
)

func TestUpgradeDistributionRecords(t *testing.T) {
	output := test.CreatOutputList(2, "100")

	legacyRecords := []legacy.DistributionRecord084{
		{
			ClaimStatus:                 1,
			DistributionName:            "first",
			RecipientAddress:            output[0].Address,
			Coins:                       sdk.NewCoins(sdk.NewCoin("fury", sdk.NewInt(100))),
			DistributionStartHeight:     2,
			DistributionCompletedHeight: 3,
		},
	}

	upgradedRecords := []types.DistributionRecord{
		{
			DistributionStatus:          1,
			DistributionName:            "first",
			DistributionType:            types.Airdrop,
			RecipientAddress:            output[0].Address,
			Coins:                       sdk.NewCoins(sdk.NewCoin("fury", sdk.NewInt(100))),
			DistributionStartHeight:     2,
			DistributionCompletedHeight: 3,
		},
	}

	var tt = []struct {
		name     string
		records  []legacy.DistributionRecord084
		upgraded []types.DistributionRecord
	}{
		{"success", legacyRecords, upgradedRecords},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sojaapp, ctx := test.CreateTestApp(false)

			for _, dr := range tc.records {
				sojaapp.DispensationKeeper.Set(ctx,
					types.GetDistributionRecordKey(dr.DistributionName, dr.RecipientAddress.String(), types.DistributionTypeUnknown.String()),
					sojaapp.Codec().MustMarshal(dr),
				)
			}

			keeper.UpgradeDistributionRecords(ctx, sojaapp.DispensationKeeper)

			var got []types.DistributionRecord
			iterator := sojaapp.DispensationKeeper.GetDistributionRecordsIterator(ctx)
			defer iterator.Close()
			for ; iterator.Valid(); iterator.Next() {
				var dr types.DistributionRecord
				bytesValue := iterator.Value()
				sojaapp.Codec().MustUnmarshal(bytesValue, &dr)
				got = append(got, dr)
			}

			require.Equal(t, tc.upgraded, got)
		})
	}
}

func TestUpgradeDistributions(t *testing.T) {
	legacyDistributions := []legacy.Distribution084{
		{
			DistributionName: "first",
			DistributionType: types.LiquidityMining,
		},
	}
	upgradedDistributions := []types.Distribution{
		{
			DistributionName: "first",
			DistributionType: types.LiquidityMining,
		},
	}
	var tt = []struct {
		name     string
		records  []legacy.Distribution084
		upgraded []types.Distribution
	}{
		{"success", legacyDistributions, upgradedDistributions},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sojaapp, ctx := test.CreateTestApp(false)

			for _, dr := range tc.records {
				sojaapp.DispensationKeeper.Set(ctx,
					types.GetDistributionsKey(dr.DistributionName, dr.DistributionType),
					sojaapp.Codec().MustMarshal(dr),
				)
			}

			keeper.UpgradeDistributions(ctx, sojaapp.DispensationKeeper)

			var got []types.Distribution
			iterator := sojaapp.DispensationKeeper.GetDistributionIterator(ctx)
			defer iterator.Close()
			for ; iterator.Valid(); iterator.Next() {
				var d types.Distribution
				bytesValue := iterator.Value()
				sojaapp.Codec().MustUnmarshal(bytesValue, &d)
				got = append(got, d)
			}

			require.Equal(t, tc.upgraded, got)
		})
	}
}

*/
