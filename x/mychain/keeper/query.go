package keeper

import (
	"myChain/x/mychain/types"
)

var _ types.QueryServer = Keeper{}
