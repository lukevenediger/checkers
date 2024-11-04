package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	checkers "github.com/lukevenediger/checkers"
)

type checkersTorramServer struct {
	k Keeper
}

var _ checkers.CheckersTorramServer = checkersTorramServer{}

// NewMsgServerImpl returns a message server implementation
func NewMsgServerImpl(keeper Keeper) checkers.CheckersTorramServer {
	return &checkersTorramServer{
		k: keeper,
	}
}

func (ms checkersTorramServer) CheckersCreateGm(
	goCtx context.Context,
	msg *checkers.ReqCheckersTorram,
) (*checkers.ResCheckersTorram, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate game ID length
	lenIndex := len([]byte(msg.Index))
	if lenIndex > checkers.MaxIndexLength {
		return nil, checkers.ErrIndexTooLong
	}
	if lenIndex == 0 {
		return nil, checkers.ErrEmptyIndex
	}

	// Can't create a new game for an existing index
	hasGame, err := ms.k.StoredGames.Has(ctx, msg.Index)
	if err != nil {
		panic(err)
	}
	if hasGame {
		return nil, checkers.ErrDuplicateGameIndex
	}

	// Create a new checkers game
	storedGame := checkers.NewStoredGame(msg.Black, msg.Red, ctx.BlockTime())

	// Stored game must validate
	if err := storedGame.Validate(); err != nil {
		return nil, err
	}
	// Pushing to storage must succeed
	if err := ms.k.StoredGames.Set(ctx, msg.Index, storedGame); err != nil {
		return nil, err
	}

	return &checkers.ResCheckersTorram{}, nil
}

func (ms checkersTorramServer) ForfeitGm(
	goCtx context.Context,
	msg *checkers.ReqForfeitGm,
) (*checkers.ResForfeitGm, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, err := ms.k.StoredGames.Get(ctx, msg.Index)
	if err != nil {
		return nil, err
	}

	// Only a game that's ready or in progress can be forfeited
	if !storedGame.IsReadyOrInProgress() {
		return nil, checkers.ErrActionNotAllowedForGameState
	}

	// Only the player whose turn it is can forfeit
	playerTurnAddress, err := storedGame.GetTurnAddress()
	if err != nil {
		return nil, err
	}
	if playerTurnAddress.String() != msg.Forfeiter {
		return nil, checkers.ErrNotPlayersTurn
	}

	// Update the game end time and state
	storedGame.EndGame(ctx.BlockTime(), checkers.GameState_GAME_STATE_FORFEITED)

	// Pushing to storage must succeed
	if err := ms.k.StoredGames.Set(ctx, msg.Index, storedGame); err != nil {
		return nil, err
	}

	return &checkers.ResForfeitGm{}, nil
}
