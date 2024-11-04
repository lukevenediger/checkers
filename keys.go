package checkers

import "cosmossdk.io/collections"

const ModuleName = "checkers-torram"
const MaxIndexLength = 256
const StoreKey = "StoredGames"
const GameTimeLayout = "2006-01-02 15:04:05.999999999 +0000 UTC"

var (
	ParamsKey      = collections.NewPrefix("Params")
	StoredGamesKey = collections.NewPrefix(StoreKey + "/value/")
)
