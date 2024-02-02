package keeper

import (
	"github.com/foxytanuki/mychain/x/mychain/types"
)

var _ types.QueryServer = Keeper{}
