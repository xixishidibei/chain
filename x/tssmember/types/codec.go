package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	tsstypes "github.com/bandprotocol/chain/v2/x/tss/types"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgCreateGroup{}, "tssmember/CreateGroup")
	legacy.RegisterAminoMsg(cdc, &MsgReplaceGroup{}, "tssmember/ReplaceGroup")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateGroupFee{}, "tssmember/UpdateGroupFee")
	legacy.RegisterAminoMsg(cdc, &MsgRequestSignature{}, "tssmember/RequestSignature")
	legacy.RegisterAminoMsg(cdc, &MsgActivate{}, "tssmember/Activate")
	legacy.RegisterAminoMsg(cdc, &MsgHealthCheck{}, "tssmember/HealthCheck")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "tssmember/UpdateParams")

	cdc.RegisterConcrete(&tsstypes.TextSignatureOrder{}, "tss/TextSignatureOrder", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGroup{},
		&MsgReplaceGroup{},
		&MsgUpdateGroupFee{},
		&MsgRequestSignature{},
		&MsgActivate{},
		&MsgHealthCheck{},
		&MsgUpdateParams{},
	)
	registry.RegisterInterface(
		"tss.v1beta1.Content",
		(*tsstypes.Content)(nil),
		&tsstypes.TextSignatureOrder{},
	)
}

// RegisterRequestSignatureTypeCodec registers an external signature request type defined
// in another module for the internal ModuleCdc. This allows the MsgRequestSignature
// to be correctly Amino encoded and decoded.
//
// NOTE: This should only be used for applications that are still using a concrete
// Amino codec for serialization.
func RegisterSignatureOrderTypeCodec(o interface{}, name string) {
	amino.RegisterConcrete(o, name, nil)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	sdk.RegisterLegacyAminoCodec(amino)
}
