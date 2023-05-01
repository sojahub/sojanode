package v42

import (
	v039ethbridge "github.com/Sojahub/sojanode/x/ethbridge/legacy/v39"
	"github.com/Sojahub/sojanode/x/ethbridge/types"
)

func Migrate(state v039ethbridge.GenesisState) *types.GenesisState {
	return &types.GenesisState{
		CethReceiveAccount: state.CethReceiverAccount.String(),
		PeggyTokens:        state.PeggyTokens,
	}
}
