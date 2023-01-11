package keeper_test

import (
	"context"
	"testing"

	"github.com/alice/checkers/x/checkers/keeper"
	keepertest "github.com/alice/checkers/x/checkers/testutil/keeper"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
