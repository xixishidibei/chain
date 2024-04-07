package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/pkg/tss"
)

const (
	// module name
	ModuleName = "bandtss"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName

	// RouterKey is the message route for the bandtss module
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the bandtss module
	QuerierRoute = ModuleName
)

var (

	// GlobalStoreKeyPrefix is the prefix for global primitive state variables.
	GlobalStoreKeyPrefix = []byte{0x00}
	// ParamsKeyPrefix is a prefix for keys that store bandtss's parameters
	ParamsKeyPrefix = []byte{0x01}
	// StatusStoreKeyPrefix is the prefix for status store.
	MemberStoreKeyPrefix = []byte{0x02}
	// GroupIDStoreKeyPrefix is the prefix for groupID store.
	GroupIDStoreKeyPrefix = []byte{0x03}
	// SigningStoreKeyPrefix is the prefix for bandtss Signing store.
	SigningStoreKeyPrefix = []byte{0x04}

	// SigningCountStoreKey is the key that keeps the total number of Signing.
	SigningCountStoreKey = append(GlobalStoreKeyPrefix, []byte("SigningCount")...)

	// CurrentGroupIDKey is the key for storing the current group ID under GroupIDStoreKeyPrefix.
	CurrentGroupIDStoreKey = append(GroupIDStoreKeyPrefix, []byte{0x00}...)
	// ReplacingGroupIDKey  is the key for storing the replacing group ID under GroupIDStoreKeyPrefix.
	ReplacingGroupIDStoreKey = append(GroupIDStoreKeyPrefix, []byte{0x01}...)

	// SigningInfoStoreKeyPrefix is the key for storing the bandtss signing info under SigningStoreKeyPrefix.
	SigningInfoStoreKeyPrefix = append(SigningStoreKeyPrefix, []byte{0x00}...)
	// SigningIDMappingStoreKey is the key for storing the bandtss signing ID under SigningStoreKeyPrefix.
	SigningIDMappingStoreKeyPrefix = append(SigningStoreKeyPrefix, []byte{0x01}...)
)

func MemberStoreKey(address sdk.AccAddress) []byte {
	return append(MemberStoreKeyPrefix, address...)
}

func SigningStoreKey(id SigningID) []byte {
	return append(SigningInfoStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(id))...)
}

func SigningIDMappingStoreKey(id tss.SigningID) []byte {
	return append(SigningIDMappingStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(id))...)
}
