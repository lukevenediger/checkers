package keeper_test

import (
	"testing"
	"time"

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

func TestMsgServerTestSuite(t *testing.T) {
	suite.Run(t, new(MsgServerTestSuite))
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

type MsgServerTestSuite struct {
	KeeperTestSuite
	msgServer checkers.CheckersTorramServer
}

func (s *MsgServerTestSuite) SetupTest() {
	s.KeeperTestSuite.SetupTest()
	s.msgServer = keeper.NewMsgServerImpl(s.keeper)
}

func (s *MsgServerTestSuite) TestCheckersCreateGm() {

	_, addrs := testutil.PrivKeyAddressPairs(2)
	alice := addrs[0].String()
	bob := addrs[1].String()

	testCases := map[string]struct {
		reqMsg   *checkers.ReqCheckersTorram
		resMsg   *checkers.ResCheckersTorram
		wantErr  string
		preHook  func(ss *MsgServerTestSuite) error
		postHook func(ss *MsgServerTestSuite) error
	}{
		"success": {
			&checkers.ReqCheckersTorram{
				Creator: alice,
				Index:   "game1",
				Black:   alice,
				Red:     bob,
			},
			&checkers.ResCheckersTorram{},
			"",
			nil,
			func(ss *MsgServerTestSuite) error {
				// Check game state is Ready
				game, err := ss.keeper.StoredGames.Get(ss.ctx, "game1")
				if err != nil {
					return err
				}
				ss.Require().NotNil(game)
				ss.Require().NotEmpty(game.StartTime)
				ss.Require().Empty(game.EndTime)
				ss.Require().Equal(checkers.GameState_GAME_STATE_READY, game.State)
				return nil
			},
		},
		"fail: game already exists": {
			&checkers.ReqCheckersTorram{
				Creator: alice,
				Index:   "game1",
				Black:   alice,
				Red:     bob,
			},
			nil,
			"duplicate game index",
			func(ss *MsgServerTestSuite) error {
				return ss.keeper.StoredGames.Set(ss.ctx, "game1", checkers.StoredGame{})
			},
			nil,
		},
		"fail: index too long": {
			&checkers.ReqCheckersTorram{
				Creator: alice,
				Index:   testutil.RandLetters(checkers.MaxIndexLength + 1),
				Black:   alice,
				Red:     bob,
			},
			nil,
			"index too long",
			nil,
			nil,
		},
		"fail: empty index": {
			&checkers.ReqCheckersTorram{
				Creator: alice,
				Index:   "",
				Black:   alice,
				Red:     bob,
			},
			nil,
			"empty index not allowed",
			nil,
			nil,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			s.SetupTest()
			if tc.preHook != nil {
				err := tc.preHook(s)
				if err != nil {
					s.Failf("preHook failed: %v", err.Error())
				}
			}

			resp, err := s.msgServer.CheckersCreateGm(s.ctx, tc.reqMsg)

			if tc.wantErr != "" {
				s.Error(err)
				s.ErrorContains(err, tc.wantErr)
				return
			}

			s.NoError(err)

			// Compare the response
			s.Equal(tc.resMsg, resp)

			// Execute post hook
			if tc.postHook != nil {
				err := tc.postHook(s)
				if err != nil {
					s.Failf("postHook failed: %v", err.Error())
				}
			}
		})
	}
}

func (s *MsgServerTestSuite) TestForfeitGm() {

	_, addrs := testutil.PrivKeyAddressPairs(3)
	alice := addrs[0].String()
	bob := addrs[1].String()

	testCases := map[string]struct {
		reqMsg   *checkers.ReqForfeitGm
		resMsg   *checkers.ResForfeitGm
		wantErr  string
		preHook  func(ss *MsgServerTestSuite) error
		postHook func(ss *MsgServerTestSuite) error
	}{
		"success": {
			&checkers.ReqForfeitGm{
				Forfeiter: alice,
				Index:     "game1",
			},
			&checkers.ResForfeitGm{},
			"",
			func(ss *MsgServerTestSuite) error {
				game := checkers.NewStoredGame(alice, bob, time.Now())
				return ss.keeper.StoredGames.Set(ss.ctx, "game1", game)
			},
			func(ss *MsgServerTestSuite) error {
				// Check game state is Ready
				game, err := ss.keeper.StoredGames.Get(ss.ctx, "game1")
				if err != nil {
					return err
				}
				ss.Require().NotNil(game)
				ss.Require().NotEmpty(game.StartTime)
				ss.Require().NotEmpty(game.EndTime)
				ss.Require().Equal(checkers.GameState_GAME_STATE_FORFEITED, game.State)
				return nil
			},
		},
		"not player's turn": {
			&checkers.ReqForfeitGm{
				Forfeiter: bob,
				Index:     "game1",
			},
			nil,
			"not player's turn",
			func(ss *MsgServerTestSuite) error {
				game := checkers.NewStoredGame(alice, bob, time.Now())
				return ss.keeper.StoredGames.Set(ss.ctx, "game1", game)
			},
			nil,
		},
		"game not found": {
			&checkers.ReqForfeitGm{
				Forfeiter: alice,
				Index:     "game1",
			},
			nil,
			"not found: key 'game1'",
			nil,
			nil,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			s.SetupTest()
			if tc.preHook != nil {
				err := tc.preHook(s)
				if err != nil {
					s.Failf("preHook failed: %v", err.Error())
				}
			}

			resp, err := s.msgServer.ForfeitGm(s.ctx, tc.reqMsg)

			if tc.wantErr != "" {
				s.Error(err)
				s.ErrorContains(err, tc.wantErr)
				return
			}

			s.NoError(err)

			// Compare the response
			s.Equal(tc.resMsg, resp)

			// Execute post hook
			if tc.postHook != nil {
				err := tc.postHook(s)
				if err != nil {
					s.Failf("postHook failed: %v", err.Error())
				}
			}

		})
	}
}
