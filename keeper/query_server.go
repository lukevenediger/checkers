package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	checkers "github.com/lukevenediger/checkers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ checkers.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the QueryServer interface
func NewQueryServerImpl(keeper Keeper) checkers.QueryServer {
	return queryServer{
		k: keeper,
	}
}

type queryServer struct {
	k Keeper
}

func (qs queryServer) GetGame(
	ctx context.Context,
	req *checkers.QueryGetGameRequest,
) (*checkers.QueryGetGameResponse, error) {
	game, err := qs.k.StoredGames.Get(ctx, req.Index)
	if err == nil {
		return &checkers.QueryGetGameResponse{
			Game: &game,
		}, nil
	}
	// If the game is not found, return an empty response
	if errors.Is(err, collections.ErrNotFound) {
		return &checkers.QueryGetGameResponse{
			Game: nil,
		}, nil
	}

	// If we're here then something went wrong
	return nil, status.Error(codes.Internal, err.Error())
}
