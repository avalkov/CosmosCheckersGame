package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgRejectGame{}, "cosmoscheckersgame/RejectGame", nil)

	cdc.RegisterConcrete(&MsgPlayMove{}, "cosmoscheckersgame/PlayMove", nil)

	cdc.RegisterConcrete(&MsgCreateGame{}, "cosmoscheckersgame/CreateGame", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRejectGame{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPlayMove{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGame{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
