package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	checkers "github.com/lukevenediger/checkers"
)

// Keeper defines the Checkers module's keeper
type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	// authority address is allowed to update params via MsgUpdateParams
	// and perform other privileged operations
	authority string

	// state management
	Schema      collections.Schema
	Params      collections.Item[checkers.Params]
	StoredGames collections.Map[string, checkers.StoredGame]
}

// NewKeeper creates a new instance of the keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	storeService storetypes.KVStoreService,
	authority string,
) Keeper {

	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,
		Params:       collections.NewItem(sb, checkers.ParamsKey, "params", codec.CollValue[checkers.Params](cdc)),
		StoredGames: collections.NewMap(sb,
			checkers.StoredGamesKey,
			"storedGames",
			collections.StringKey,
			codec.CollValue[checkers.StoredGame](cdc),
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority address
func (k Keeper) GetAuthority() string {
	return k.authority
}
