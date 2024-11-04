package checkers

import (
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lukevenediger/checkers/rules"
)

// NewStoredGame creates a new stored game with a fresh board
func NewStoredGame(black, red string, startTime time.Time) StoredGame {
	newBoard := rules.New()
	return StoredGame{
		Board:     newBoard.String(),
		Turn:      rules.PieceStrings[newBoard.Turn],
		Black:     black,
		Red:       red,
		StartTime: startTime.UTC().Format(GameTimeLayout),
		EndTime:   "",
		State:     GameState_GAME_STATE_READY,
	}
}

// GetBlackAddress validates and returns the black player's address
func (sg StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, parseErr := sdk.AccAddressFromBech32(sg.Black)
	return black, errors.Wrapf(parseErr, ErrInvalidBlack.Error(), sg.Black)
}

// GetRedAddress validates and returns the red player's address
func (sg StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, parseErr := sdk.AccAddressFromBech32(sg.Red)
	return red, errors.Wrapf(parseErr, ErrInvalidRed.Error(), sg.Red)
}

func (sg StoredGame) GetTurnAddress() (turn sdk.AccAddress, err error) {
	player := rules.StringPieces[sg.Turn].Player
	if player == rules.RED_PLAYER {
		return sg.GetRedAddress()
	} else if player == rules.BLACK_PLAYER {
		return sg.GetBlackAddress()
	} else {
		return nil, errors.Wrapf(ErrGameNotParseable, "invalid turn: %s", sg.Turn)
	}
}

// ParseGame returns an active checkers game
func (sg StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(sg.Board)
	if errBoard != nil {
		return nil, errors.Wrap(errBoard, ErrGameNotParseable.Error())
	}
	board.Turn = rules.StringPieces[sg.Turn].Player
	if board.Turn.Color == "" {
		return nil, errors.Wrapf(ErrGameNotParseable, "invalid turn: %s", sg.Turn)
	}
	return board, nil
}

// Validate ensures players are valid and game is parseable
func (sg StoredGame) Validate() (err error) {
	bAddr, err := sg.GetBlackAddress()
	if err != nil {
		return err
	}
	rAddr, err := sg.GetRedAddress()
	if err != nil {
		return err
	}

	// Addresses cannot match
	if bAddr.Equals(rAddr) {
		return ErrCannotPlayAgainstSelf
	}

	_, err = sg.ParseGame()
	return err
}

func (sg *StoredGame) EndGame(endTime time.Time, state GameState) {
	sg.State = state
	sg.EndTime = endTime.UTC().Format(GameTimeLayout)
}

func (sg *StoredGame) IsReadyOrInProgress() bool {
	return sg.State == GameState_GAME_STATE_READY ||
		sg.State == GameState_GAME_STATE_PLAYING
}
