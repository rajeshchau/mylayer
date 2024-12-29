package mylayer_test

import (
	"testing"

	keepertest "mylayer/testutil/keeper"
	"mylayer/testutil/nullify"
	mylayer "mylayer/x/mylayer/module"
	"mylayer/x/mylayer/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MylayerKeeper(t)
	mylayer.InitGenesis(ctx, k, genesisState)
	got := mylayer.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
