package keeper_test

import (
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkaddress "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	cosmostestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	checkers "github.com/lukevenediger/checkers"
	"github.com/lukevenediger/checkers/keeper"
	testutil "github.com/lukevenediger/checkers/util/test"
	"github.com/stretchr/testify/suite"
)

// Make sure that KeeperTestSuite implements the SetupTestSuite interface
var _ suite.SetupTestSuite = (*KeeperTestSuite)(nil)

// TODO: move this to common_test.go
type KeeperTestSuite struct {
	suite.Suite
	ctx          sdk.Context
	keeper       keeper.Keeper
	cdc          codec.BinaryCodec
	addressCodec address.Codec
	storeService store.KVStoreService
}

func (s *KeeperTestSuite) SetupTest() {
	_, addrs := testutil.PrivKeyAddressPairs(1)
	authority := addrs[0].String()

	key := storetypes.NewKVStoreKey(checkers.StoreKey)
	s.storeService = runtime.NewKVStoreService(key)
	testCtx := cosmostestutil.DefaultContextWithDB(s.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	s.ctx = testCtx.Ctx.WithBlockHeader(cmtproto.Header{Time: cmttime.Now()})
	s.addressCodec = sdkaddress.NewBech32Codec("cosmos")
	encCfg := moduletestutil.MakeTestEncodingConfig()
	s.cdc = encCfg.Codec

	// Default optional expectations
	s.keeper = keeper.NewKeeper(s.cdc,
		s.addressCodec,
		s.storeService,
		authority,
	)
}
