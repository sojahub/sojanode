package ibctransfer_test

import (
	app2 "github.com/Sojahub/sojanode/app"
	sctransfertypes "github.com/Sojahub/sojanode/x/ibctransfer/types"
	"testing"

	"github.com/Sojahub/sojanode/x/ethbridge/test"

	tokenregistrytest "github.com/Sojahub/sojanode/x/tokenregistry/test"
	tokenregistrytypes "github.com/Sojahub/sojanode/x/tokenregistry/types"
	transfertypes "github.com/cosmos/ibc-go/v2/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"

	"github.com/Sojahub/sojanode/x/ibctransfer/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestShouldConvertIncomingCoins(t *testing.T) {
	app, ctx, _ := tokenregistrytest.CreateTestApp(false)
	unitDenomEntry := tokenregistrytypes.RegistryEntry{
		Denom:     "ceth",
		Decimals:  18,
		UnitDenom: "ceth",
	}
	ibcRegistryEntry := tokenregistrytypes.RegistryEntry{
		Denom:     "ueth",
		Decimals:  10,
		UnitDenom: "ceth",
	}
	nonIBCRegistryEntry := tokenregistrytypes.RegistryEntry{
		Denom:    "cusdt",
		Decimals: 6,
	}
	app.TokenRegistryKeeper.SetToken(ctx, &unitDenomEntry)
	app.TokenRegistryKeeper.SetToken(ctx, &ibcRegistryEntry)
	app.TokenRegistryKeeper.SetToken(ctx, &nonIBCRegistryEntry)
	registry := app.TokenRegistryKeeper.GetRegistry(ctx)
	entry1, err := app.TokenRegistryKeeper.GetEntry(registry, "ueth")
	require.NoError(t, err)
	entry1c, err := app.TokenRegistryKeeper.GetEntry(registry, entry1.UnitDenom)
	require.NoError(t, err)
	require.True(t, entry1c.Decimals > entry1.Decimals)
	diff := uint64(entry1c.Decimals - entry1.Decimals)
	convAmount, err := helpers.ConvertIncomingCoins("1000000000000", diff)
	require.NoError(t, err)
	incomingDeduction := sdk.NewCoin("ueth", sdk.NewIntFromUint64(1000000000000))
	incomingAddition := sdk.NewCoin("ceth", convAmount)
	require.Equal(t, incomingDeduction.Denom, "ueth")
	require.Equal(t, incomingDeduction.Amount.String(), "1000000000000")
	require.Equal(t, incomingAddition.Denom, "ceth")
	require.Equal(t, incomingAddition.Amount.String(), "100000000000000000000")
	entry2, err := app.TokenRegistryKeeper.GetEntry(registry, "cusdt")
	require.NoError(t, err)
	_, err = app.TokenRegistryKeeper.GetEntry(registry, entry2.UnitDenom)
	require.Error(t, err)
}

func TestGetConvForIncomingCoins(t *testing.T) {
	app, ctx, _ := tokenregistrytest.CreateTestApp(false)
	ibcRegistryEntry := tokenregistrytypes.RegistryEntry{
		Denom:     "ueth",
		Decimals:  10,
		UnitDenom: "ceth",
	}
	unitDenomEntry := tokenregistrytypes.RegistryEntry{
		Denom:    "ceth",
		Decimals: 18,
	}
	app.TokenRegistryKeeper.SetToken(ctx, &unitDenomEntry)
	app.TokenRegistryKeeper.SetToken(ctx, &ibcRegistryEntry)
	registry := app.TokenRegistryKeeper.GetRegistry(ctx)
	entry1, err := app.TokenRegistryKeeper.GetEntry(registry, "ueth")
	require.NoError(t, err)
	_, err = app.TokenRegistryKeeper.GetEntry(registry, "ceth")
	require.NoError(t, err)
	entry1c, err := app.TokenRegistryKeeper.GetEntry(registry, entry1.UnitDenom)
	require.NoError(t, err)
	require.True(t, entry1c.Decimals > entry1.Decimals)
	diff := uint64(entry1c.Decimals - entry1.Decimals)
	convAmount, err := helpers.ConvertIncomingCoins("1000000000000", diff)
	require.NoError(t, err)
	incomingDeduction := sdk.NewCoin("ueth", sdk.NewIntFromUint64(1000000000000))
	incomingAddition := sdk.NewCoin("ceth", convAmount)
	intAmount, _ := sdk.NewIntFromString("100000000000000000000")
	require.Equal(t, incomingDeduction, sdk.NewCoin("ueth", sdk.NewInt(1000000000000)))
	require.Equal(t, incomingAddition, sdk.NewCoin("ceth", intAmount))
}

func TestIsRecvPacketAllowed(t *testing.T) {
	app, ctx, _ := tokenregistrytest.CreateTestApp(false)
	transferPacket := channeltypes.Packet{
		Sequence:           0,
		SourcePort:         "transfer",
		SourceChannel:      "channel-0",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-1",
		Data:               nil,
	}
	returningDenom := transfertypes.FungibleTokenPacketData{
		Denom: "transfer/channel-0/fury",
	}
	whitelistedDenom := transfertypes.FungibleTokenPacketData{
		Denom: "atom",
	}
	disallowedDenom := transfertypes.FungibleTokenPacketData{
		Denom: "transfer/channel-66/atom",
	}
	app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
		Denom:                "atom",
		Decimals:             6,
		IbcCounterpartyDenom: "",
		Permissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_IBCIMPORT},
	})
	app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
		Denom:                "ibc/44F0BAC50DDD0C83DAC9CEFCCC770C12F700C0D1F024ED27B8A3EE9DD949BAD3",
		Decimals:             6,
		IbcCounterpartyDenom: "",
		Permissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_IBCIMPORT},
	})
	app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
		Denom:                "ibc/A916425D9C00464330F8B333711C4A51AA8CF0141392E7E250371EC6D4285BF2",
		Decimals:             6,
		IbcCounterpartyDenom: "",
		Permissions:          []tokenregistrytypes.Permission{},
	})
	registry := app.TokenRegistryKeeper.GetRegistry(ctx)
	entry1, err := app.TokenRegistryKeeper.GetEntry(registry, "ibc/44F0BAC50DDD0C83DAC9CEFCCC770C12F700C0D1F024ED27B8A3EE9DD949BAD3")
	require.NoError(t, err)
	permitted1 := app.TokenRegistryKeeper.CheckEntryPermissions(entry1, []tokenregistrytypes.Permission{tokenregistrytypes.Permission_IBCIMPORT})
	require.Equal(t, permitted1, true)
	got := helpers.IsRecvPacketAllowed(ctx, app.TokenRegistryKeeper, transferPacket, whitelistedDenom, entry1)
	require.Equal(t, got, true)
	entry2, err := app.TokenRegistryKeeper.GetEntry(registry, "ibc/A916425D9C00464330F8B333711C4A51AA8CF0141392E7E250371EC6D4285BF2")
	require.NoError(t, err)
	permitted2 := app.TokenRegistryKeeper.CheckEntryPermissions(entry2, []tokenregistrytypes.Permission{tokenregistrytypes.Permission_IBCIMPORT})
	require.Equal(t, permitted2, false)
	got = helpers.IsRecvPacketAllowed(ctx, app.TokenRegistryKeeper, transferPacket, disallowedDenom, entry2)
	require.Equal(t, got, false)
	entry3, err := app.TokenRegistryKeeper.GetEntry(registry, "fury")
	require.Error(t, err)
	got = helpers.IsRecvPacketAllowed(ctx, app.TokenRegistryKeeper, transferPacket, returningDenom, entry3)
	require.Equal(t, got, true)
}

func TestExecConvForIncomingCoins(t *testing.T) {
	app, ctx, _ := tokenregistrytest.CreateTestApp(false)
	addrs, _ := test.CreateTestAddrs(2)
	packet := channeltypes.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-0",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-1",
	}
	returningData := transfertypes.FungibleTokenPacketData{
		Denom:    "transfer/channel-0/ueth",
		Receiver: addrs[0].String(),
		Amount:   "0",
	}
	nonReturningData := transfertypes.FungibleTokenPacketData{
		Denom:    "transfer/channel-1/ueth",
		Receiver: addrs[0].String(),
		Amount:   "0",
	}
	ibcRegistryEntry := tokenregistrytypes.RegistryEntry{
		Denom:     "ueth",
		Decimals:  10,
		UnitDenom: "ceth",
	}
	ibcRegistryEntry2 := tokenregistrytypes.RegistryEntry{
		Denom:       "ibc/C1061B25E69D71E96BED65B5652168F41927316D07D6B417A3A9774F94A4CB7A",
		Decimals:    10,
		UnitDenom:   "ceth",
		Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_IBCIMPORT},
	}
	unitDenomEntry := tokenregistrytypes.RegistryEntry{
		Denom:    "ceth",
		Decimals: 18,
	}
	app.TokenRegistryKeeper.SetToken(ctx, &unitDenomEntry)
	app.TokenRegistryKeeper.SetToken(ctx, &ibcRegistryEntry)
	app.TokenRegistryKeeper.SetToken(ctx, &ibcRegistryEntry2)
	mintedDenom := helpers.GetMintedDenomFromPacket(packet, returningData)
	registry := app.TokenRegistryKeeper.GetRegistry(ctx)
	mintedDenomEntry, err := app.TokenRegistryKeeper.GetEntry(registry, mintedDenom)
	require.NoError(t, err)
	allowed := helpers.IsRecvPacketAllowed(ctx, app.TokenRegistryKeeper, packet, returningData, mintedDenomEntry)
	require.Equal(t, allowed, true)
	convertToDenomEntry, err := app.TokenRegistryKeeper.GetEntry(registry, mintedDenomEntry.UnitDenom)
	require.NoError(t, err)
	err = helpers.ExecConvForIncomingCoins(ctx, app.BankKeeper, mintedDenomEntry, convertToDenomEntry, packet, returningData)
	require.NoError(t, err)
	mintedDenom = helpers.GetMintedDenomFromPacket(packet, nonReturningData)
	mintedDenomEntry, err = app.TokenRegistryKeeper.GetEntry(registry, mintedDenom)
	require.NoError(t, err)
	allowed = helpers.IsRecvPacketAllowed(ctx, app.TokenRegistryKeeper, packet, nonReturningData, mintedDenomEntry)
	require.Equal(t, allowed, true)
	convertToDenomEntry, err = app.TokenRegistryKeeper.GetEntry(registry, mintedDenomEntry.UnitDenom)
	require.NoError(t, err)
	err = helpers.ExecConvForIncomingCoins(ctx, app.BankKeeper, mintedDenomEntry, convertToDenomEntry, packet, nonReturningData)
	require.NoError(t, err)
}

func TestOnRecvPacketV2(t *testing.T) {
	addrs, _ := test.CreateTestAddrs(2)
	app, ctx, _ := tokenregistrytest.CreateTestApp(false)
	packet := channeltypes.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-0",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-1",
	}
	// Xfury which originated on Sojahub
	xFuryV2Amount := "10000000000"
	returningXfury := transfertypes.FungibleTokenPacketData{
		Denom:    "transfer/channel-0/xfury",
		Receiver: addrs[0].String(),
		Amount:   xFuryV2Amount,
	}
	ibcRegistryEntryXFury := tokenregistrytypes.RegistryEntry{
		Denom:     "xfury",
		Decimals:  10,
		UnitDenom: "fury",
	}
	// Adding registry entry for fury as fury is UnitDenom for xfury
	ibcRegistryEntryFury := tokenregistrytypes.RegistryEntry{
		Denom:     "fury",
		Decimals:  18,
		UnitDenom: "fury",
	}
	app.TokenRegistryKeeper.SetToken(ctx, &ibcRegistryEntryXFury)
	app.TokenRegistryKeeper.SetToken(ctx, &ibcRegistryEntryFury)
	mintedXFury := helpers.GetMintedDenomFromPacket(packet, returningXfury)
	registry := app.TokenRegistryKeeper.GetRegistry(ctx)
	mintedXFuryEntry, err := app.TokenRegistryKeeper.GetEntry(registry, mintedXFury)
	require.NoError(t, err)

	allowed := helpers.IsRecvPacketAllowed(ctx, app.TokenRegistryKeeper, packet, returningXfury, mintedXFuryEntry)
	require.Equal(t, allowed, true)
	convertToDenomEntry, err := app.TokenRegistryKeeper.GetEntry(registry, mintedXFuryEntry.UnitDenom)
	require.NoError(t, err)
	intAmount, ok := sdk.NewIntFromString(xFuryV2Amount)
	require.True(t, ok)
	err = app2.AddCoinsToAccount(sctransfertypes.ModuleName, app.BankKeeper, ctx, addrs[0], sdk.NewCoins(sdk.NewCoin("xfury", intAmount)))
	require.NoError(t, err)

	diff := uint64(convertToDenomEntry.Decimals - mintedXFuryEntry.Decimals)
	// This is the reduced precision xToken coming in , so we know for sure conversion to uint64 will not cause problems
	convAmount, err := helpers.ConvertIncomingCoins(xFuryV2Amount, diff)
	require.NoError(t, err)
	finalCoins := sdk.NewCoins(sdk.NewCoin(convertToDenomEntry.Denom, convAmount))
	escrowAddress := sctransfertypes.GetEscrowAddress(packet.GetDestPort(), packet.GetDestChannel())
	err = app2.AddCoinsToAccount(sctransfertypes.ModuleName, app.BankKeeper, ctx, escrowAddress, finalCoins)
	require.NoError(t, err)
	err = helpers.ExecConvForIncomingCoins(ctx, app.BankKeeper, mintedXFuryEntry, convertToDenomEntry, packet, returningXfury)
	require.NoError(t, err)
}

func TestIsRecvPacketReturning(t *testing.T) {
	packet := channeltypes.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-0",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-1",
	}
	returningData := transfertypes.FungibleTokenPacketData{
		Denom: "transfer/channel-0/atom",
	}
	nonReturningData := transfertypes.FungibleTokenPacketData{
		Denom: "transfer/channel-11/atom",
	}
	got := transfertypes.ReceiverChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), returningData.Denom)
	require.Equal(t, got, true)
	got = transfertypes.ReceiverChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), nonReturningData.Denom)
	require.Equal(t, got, false)
}

func TestGetMintedDenomFromPacket(t *testing.T) {
	packet := channeltypes.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-0",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-1",
	}
	returningData := transfertypes.FungibleTokenPacketData{
		Denom: "transfer/channel-0/atom",
	}
	returningData2 := transfertypes.FungibleTokenPacketData{
		Denom: "transfer/channel-0/fury",
	}
	nonReturningData := transfertypes.FungibleTokenPacketData{
		Denom: "transfer/channel-11/atom",
	}
	nonReturningData2 := transfertypes.FungibleTokenPacketData{
		Denom: "transfer/channel-11/fury",
	}
	nonReturningData3 := transfertypes.FungibleTokenPacketData{
		Denom: "transfer/channel-1/fury",
	}
	nonReturningData4 := transfertypes.FungibleTokenPacketData{
		Denom: "transfer/channel-1/atom",
	}
	returningDenom := helpers.GetMintedDenomFromPacket(packet, returningData)
	returningDenom2 := helpers.GetMintedDenomFromPacket(packet, returningData2)
	nonReturningDenom := helpers.GetMintedDenomFromPacket(packet, nonReturningData)
	nonReturningDenom2 := helpers.GetMintedDenomFromPacket(packet, nonReturningData2)
	nonReturningDenom3 := helpers.GetMintedDenomFromPacket(packet, nonReturningData3)
	nonReturningDenom4 := helpers.GetMintedDenomFromPacket(packet, nonReturningData4)
	require.Equal(t, "atom", returningDenom)
	require.Equal(t, "fury", returningDenom2)
	require.Equal(t, "ibc/611BB1D7CBB019DBA91690697697B3CB56335EBCFDD4573B9A11A34A20802940", nonReturningDenom)
	require.Equal(t, "ibc/66CEDBF4DD8DCC6F67D6BEFD7494CBB8456F77BD403BC4E59B7BE3ABB0E2A98E", nonReturningDenom2)
	require.Equal(t, "ibc/5FB4ED91304BD2AE8C19D497A6F223BAFE93836726EE5D4A726E0146EE74B479", nonReturningDenom3)
	require.Equal(t, "ibc/6D0449781D39534D032041B75F6C32DB251650083F7AC79C3975FFB7CDF7727F", nonReturningDenom4)
}
