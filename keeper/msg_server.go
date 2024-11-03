package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	checkers "github.com/lukevenediger/checkers"
	"github.com/lukevenediger/checkers/rules"
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
	if length := len([]byte(msg.Index)); length < 1 || length > checkers.MaxIndexLength {
		return nil, checkers.ErrIndexTooLong
	}
	// Can't create a new game for an existing index
	if _, err := ms.k.StoredGames.Get(ctx, msg.Index); err == nil || errors.Is(err, collections.ErrEncoding) {
		return nil, fmt.Errorf("game already exists at index %s", msg.Index)
	}

	newBoard := rules.New()
	storedGame := checkers.StoredGame{
		Board:            newBoard.String(),
		Turn:             rules.PieceStrings[newBoard.Turn],
		Black:            msg.Black,
		Red:              msg.Red,
		State:            checkers.GameState_GAME_STATE_READY,
		StartHeight:      ctx.BlockHeight(),
		LastActionHeight: ctx.BlockHeight(),
	}

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
