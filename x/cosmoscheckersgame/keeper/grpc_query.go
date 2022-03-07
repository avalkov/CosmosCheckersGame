package keeper

import (
	"github.com/avalkov/CosmosCheckersGame/x/cosmoscheckersgame/types"
)

var _ types.QueryServer = Keeper{}
