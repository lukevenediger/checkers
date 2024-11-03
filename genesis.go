package checkers

// NewGenesisState creates a new instance of GenesisState with default values
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate ensures GenesisState has valid parameters
func (gs *GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	// To help find duplicate games
	unique := make(map[string]bool)

	// Traverse all stored games provided in genesis state
	for _, indexedStoredGame := range gs.IndexedStoredGameList {
		// Length must be non-zero and less than the max
		if length := len([]byte(indexedStoredGame.Index)); length < 1 || length > MaxIndexLength {
			return ErrIndexTooLong
		}
		// Cannot have multiple games with the same index
		if _, ok := unique[indexedStoredGame.Index]; ok {
			return ErrDuplicateAddress
		}
		// Stored game must be valid
		if err := indexedStoredGame.StoredGame.Validate(); err != nil {
			return err
		}
		// This is a valid game, add it to the unique map
		unique[indexedStoredGame.Index] = true
	}

	return nil
}
