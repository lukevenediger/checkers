package keeper_test

import (
	"testing"
	"time"

	"github.com/lukevenediger/checkers"
	"github.com/lukevenediger/checkers/keeper"
	"github.com/stretchr/testify/suite"
)

// Make sure that QueryServerTestSuite implements the SetupTestSuite interface
var _ suite.SetupTestSuite = (*QueryServerTestSuite)(nil)

type QueryServerTestSuite struct {
	KeeperTestSuite
	queryServer checkers.QueryServer
	msgServer   checkers.CheckersTorramServer
}

func (s *QueryServerTestSuite) SetupTest() {
	s.KeeperTestSuite.SetupTest()
	s.queryServer = keeper.NewQueryServerImpl(s.keeper)
	s.msgServer = keeper.NewMsgServerImpl(s.keeper)
}

func TestQueryServerTestSuite(t *testing.T) {
	suite.Run(t, new(QueryServerTestSuite))
}

func (s *QueryServerTestSuite) TestGetGame() {

	referenceGame := checkers.NewStoredGame("alice", "bob", time.Now())

	testCases := map[string]struct {
		reqMsg  *checkers.QueryGetGameRequest
		resMsg  *checkers.QueryGetGameResponse
		wantErr string
		preHook func(ss *QueryServerTestSuite) error
	}{
		"success": {
			&checkers.QueryGetGameRequest{
				Index: "game1",
			},
			&checkers.QueryGetGameResponse{
				Game: &referenceGame,
			},
			"",
			func(ss *QueryServerTestSuite) error {
				return ss.keeper.StoredGames.Set(
					ss.ctx,
					"game1",
					referenceGame,
				)
			},
		},
		"not found": {
			&checkers.QueryGetGameRequest{
				Index: "game1",
			},
			&checkers.QueryGetGameResponse{},
			"",
			nil,
		},
	}

	for name, tc := range testCases {
		tc := tc
		s.Run(name, func() {

			// Run the preHook if there is one
			if tc.preHook != nil {
				err := tc.preHook(s)
				s.Require().NoError(err)
			}

			// Execute the get game query
			res, err := s.queryServer.GetGame(s.ctx, tc.reqMsg)

			if tc.wantErr != "" {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.wantErr)
				return
			}

			s.Require().NoError(err)
			s.Require().Equal(tc.resMsg, res)
		})
	}
}
