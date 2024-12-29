package keeper

import (
	"mylayer/x/mylayer/types"
)

var _ types.QueryServer = Keeper{}
