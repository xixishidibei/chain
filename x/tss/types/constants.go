package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func GetMsgGrants() []string {
	return []string{
		sdk.MsgTypeURL(&MsgSubmitDKGRound1{}),
		sdk.MsgTypeURL(&MsgSubmitDKGRound2{}),
		sdk.MsgTypeURL(&MsgComplain{}),
		sdk.MsgTypeURL(&MsgConfirm{}),
		sdk.MsgTypeURL(&MsgSubmitDEs{}),
		sdk.MsgTypeURL(&MsgSubmitSignature{}),
	}
}

const (
	AddrLen   = 20
	uint64Len = 8
)
