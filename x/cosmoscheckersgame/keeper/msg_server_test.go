package keeper

import (
	"context"
	"testing"

	"github.com/avalkov/CosmosCheckersGame/x/cosmoscheckersgame/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	keeper, ctx := setupKeeper(t)
	return NewMsgServerImpl(*keeper), sdk.WrapSDKContext(ctx)
}
