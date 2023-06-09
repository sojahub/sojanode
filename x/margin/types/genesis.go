package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: &Params{
			LeverageMax:                              sdk.NewDec(2),
			HealthGainFactor:                         sdk.NewDec(1),
			InterestRateMin:                          sdk.NewDecWithPrec(5, 3),
			InterestRateMax:                          sdk.NewDec(3),
			InterestRateDecrease:                     sdk.NewDecWithPrec(1, 1),
			InterestRateIncrease:                     sdk.NewDecWithPrec(1, 1),
			ForceCloseFundPercentage:                 sdk.NewDecWithPrec(1, 1),
			ForceCloseFundAddress:                    "did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm",
			IncrementalInterestPaymentFundPercentage: sdk.NewDecWithPrec(1, 1),
			IncrementalInterestPaymentFundAddress:    "did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm",
			PoolOpenThreshold:                        sdk.NewDecWithPrec(1, 1),
			RemovalQueueThreshold:                    sdk.NewDecWithPrec(1, 1),
			EpochLength:                              1,
			MaxOpenPositions:                         10000,
			Pools:                                    []string{},
			SqModifier:                               sdk.MustNewDecFromStr("10000000000000000000000000"),
			SafetyFactor:                             sdk.MustNewDecFromStr("1.05"),
			IncrementalInterestPaymentEnabled:        true,
			ClosedPools:                              []string{},
			WhitelistingEnabled:                      false,
			FuryCollateralEnabled:                   true,
		},
	}
}
