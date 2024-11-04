package checkers

import "cosmossdk.io/errors"

var (
	ErrIndexTooLong                 = errors.Register(ModuleName, 2, "index too long")
	ErrEmptyIndex                   = errors.Register(ModuleName, 3, "empty index not allowed")
	ErrDuplicateGameIndex           = errors.Register(ModuleName, 4, "duplicate game index")
	ErrInvalidBlack                 = errors.Register(ModuleName, 5, "black address is invalid: %s")
	ErrInvalidRed                   = errors.Register(ModuleName, 6, "red address is invalid: %s")
	ErrGameNotParseable             = errors.Register(ModuleName, 7, "game is not parseable")
	ErrNotPlayersTurn               = errors.Register(ModuleName, 8, "not player's turn")
	ErrActionNotAllowedForGameState = errors.Register(ModuleName, 9, "action not allowed for game state")
)
