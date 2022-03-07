package keeper

import (
	"context"
	"strconv"

	"github.com/avalkov/CosmosCheckersGame/x/cosmoscheckersgame/rules"
	"github.com/avalkov/CosmosCheckersGame/x/cosmoscheckersgame/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("NextGame not found")
	}
	newIndex := strconv.FormatUint(nextGame.IdValue, 10)

	game := rules.New()
	storedGame := types.StoredGame{
		Creator:   msg.Creator,
		Index:     newIndex,
		Game:      game.String(),
		Red:       msg.Red,
		Black:     msg.Black,
		MoveCount: 0,
	}

	storedGame.Turn = game.Turn.Color

	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SetStoredGame(ctx, storedGame)

	nextGame.IdValue++
	k.Keeper.SetNextGame(ctx, nextGame)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "cosmoscheckersgame"),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.StoredGameEventKey),
			sdk.NewAttribute(types.StoredGameEventCreator, msg.Creator),
			sdk.NewAttribute(types.StoredGameEventIndex, newIndex),
			sdk.NewAttribute(types.StoredGameEventRed, msg.Red),
			sdk.NewAttribute(types.StoredGameEventBlack, msg.Black),
		),
	)

	return &types.MsgCreateGameResponse{IdValue: newIndex}, nil
}
