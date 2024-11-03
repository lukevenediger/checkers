package checkers

import (
	fmt "fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lukevenediger/checkers/rules"
)

// GetBlackAddress validates and returns the black player's address
func (sg StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, parseErr := sdk.AccAddressFromBech32(sg.Black)
	return black, errors.Wrapf(parseErr, ErrInvalidBlack.Error(), sg.Black)
}

// GetRedAddress validates and returns the red player's address
func (sg StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, parseErr := sdk.AccAddressFromBech32(sg.Red)
	return red, errors.Wrapf(parseErr, ErrInvalidBlack.Error(), sg.Red)
}

// ParseGame returns an active checkers game
func (sg StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(sg.Board)
	if errBoard != nil {
		return nil, errors.Wrap(errBoard, ErrGameNotParseable.Error())
	}
	board.Turn = rules.StringPieces[sg.Turn].Player
	if board.Turn.Color == "" {
		return nil, errors.Wrapf(fmt.Errorf("turn: %s", sg.Turn), ErrGameNotParseable.Error())
	}
	return board, nil
}

// Validate ensures players are valid and game is parseable
func (sg StoredGame) Validate() (err error) {
	_, err = sg.GetBlackAddress()
	if err != nil {
		return err
	}
	_, err = sg.GetRedAddress()
	if err != nil {
		return err
	}
	_, err = sg.ParseGame()
	return err
}
