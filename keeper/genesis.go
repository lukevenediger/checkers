package keeper

import (
	"context"

	checkers "github.com/lukevenediger/checkers"
)

// InitGenesis initialises the model using genesis state
func (k *Keeper) InitGenesis(
	ctx context.Context,
	data *checkers.GenesisState,
) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	// Push all games into storage
	for _, indexedStoredGame := range data.IndexedStoredGameList {
		if err := k.StoredGames.Set(ctx, indexedStoredGame.Index, indexedStoredGame.StoredGame); err != nil {
			return err
		}
	}

	return nil
}

func (k *Keeper) ExportGenesis(
	ctx context.Context,
) (*checkers.GenesisState, error) {

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	var indexedStoredGames []checkers.IndexedStoredGame
	if err := k.StoredGames.Walk(ctx, nil, func(index string, storedGame checkers.StoredGame) (bool, error) {
		indexedStoredGames = append(indexedStoredGames, checkers.IndexedStoredGame{
			Index:      index,
			StoredGame: storedGame,
		})
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &checkers.GenesisState{
		Params:                params,
		IndexedStoredGameList: indexedStoredGames,
	}, nil
}
