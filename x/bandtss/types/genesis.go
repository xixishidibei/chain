package types

import "github.com/bandprotocol/chain/v3/pkg/tss"

// NewGenesisState creates a new GenesisState instance.
func NewGenesisState(
	params Params,
	members []Member,
	currentGroupID tss.GroupID,
) *GenesisState {
	return &GenesisState{
		Params:         params,
		Members:        members,
		CurrentGroupID: currentGroupID,
	}
}

// DefaultGenesisState returns the default genesis state.
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		DefaultParams(),
		[]Member{},
		0,
	)
}

// Validate performs basic validation of genesis data returning an
// error for any failed validation criteria.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	return nil
}
