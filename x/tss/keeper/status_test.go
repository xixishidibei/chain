package keeper_test

import (
	"github.com/bandprotocol/chain/v2/pkg/tss/testutil"
	"github.com/bandprotocol/chain/v2/x/tss/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *KeeperTestSuite) TestSetInActive() {
	ctx, k := s.ctx, s.app.TSSKeeper
	s.SetupGroup(types.GROUP_STATUS_ACTIVE)
	address := sdk.AccAddress(testutil.TestCases[0].Group.Members[0].PubKey())

	k.SetInActive(ctx, address)

	status := k.GetStatus(ctx, address)
	s.Require().Equal(types.MEMBER_STATUS_INACTIVE, status.Status)
}

func (s *KeeperTestSuite) TestSetActive() {
	ctx, k := s.ctx, s.app.TSSKeeper
	s.SetupGroup(types.GROUP_STATUS_ACTIVE)
	address := sdk.AccAddress(testutil.TestCases[0].Group.Members[0].PubKey())

	// Success case
	err := k.SetActive(ctx, address)
	s.Require().NoError(err)

	status := k.GetStatus(ctx, address)
	s.Require().Equal(types.MEMBER_STATUS_ACTIVE, status.Status)

	// Failed case - penalty
	k.SetInActive(ctx, address)

	err = k.SetActive(ctx, address)
	s.Require().ErrorIs(err, types.ErrTooSoonToActivate)

	// Failed case - no member
	err = k.SetActive(ctx, address)
	s.Require().Error(err)
}
