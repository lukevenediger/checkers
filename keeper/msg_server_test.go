package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lukevenediger/checkers"
	"github.com/lukevenediger/checkers/keeper"
	mock_test "github.com/lukevenediger/checkers/test/mocks"
	testutil "github.com/lukevenediger/checkers/test/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// TODO: move this to common_test.go
type KeeperTestSuite struct {
	suite.Suite
	ctx          sdk.Context
	keeper       keeper.Keeper
	binaryCodec  mock_test.BinaryCodec
	addressCodec mock_test.Codec
	storeService mock_test.KVStoreService
	store        mock_test.KVStore
}

func (s *KeeperTestSuite) SetupTest() {
	_, addrs := testutil.PrivKeyAddressPairs(1)
	authority := addrs[0].String()

	s.binaryCodec = mock_test.BinaryCodec{}
	s.addressCodec = mock_test.Codec{}
	s.storeService = mock_test.KVStoreService{}
	s.store = mock_test.KVStore{}

	// Default optional expectations
	s.storeService.EXPECT().OpenKVStore(mock.Anything).Return(&s.store).Maybe()

	s.keeper = keeper.NewKeeper(&s.binaryCodec,
		&s.addressCodec,
		&s.storeService,
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

	_, addrs := testutil.PrivKeyAddressPairs(3)
	alice := addrs[0].String()
	bob := addrs[1].String()
	// eve := addrs[2].String()

	testCases := []struct {
		name    string
		reqMsg  *checkers.ReqCheckersTorram
		resMsg  *checkers.ResCheckersTorram
		wantErr string
		preHook func(ss *MsgServerTestSuite)
	}{
		{
			"success",
			&checkers.ReqCheckersTorram{
				Creator: alice,
				Index:   "game1",
				Black:   alice,
				Red:     bob,
			},
			&checkers.ResCheckersTorram{},
			"",
			nil,
		},
		{
			"fail: game already exists",
			&checkers.ReqCheckersTorram{
				Creator: alice,
				Index:   "game1",
				Black:   alice,
				Red:     bob,
			},
			nil,
			"game already exists at index game1",
			func(ss *MsgServerTestSuite) {
				s.store.EXPECT().Get([]byte("game1")).Return([]byte{}, nil)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()
			if tc.preHook != nil {
				tc.preHook(s)
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

			// Validate all expectations
			s.binaryCodec.AssertExpectations(s.T())
			s.addressCodec.AssertExpectations(s.T())
			s.storeService.AssertExpectations(s.T())
			s.store.AssertExpectations(s.T())
		})
	}
}
