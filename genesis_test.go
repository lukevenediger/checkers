package checkers_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lukevenediger/checkers"
	testutil "github.com/lukevenediger/checkers/util/test"
	"github.com/stretchr/testify/suite"
)

var _ suite.SetupTestSuite = (*genesisTestSuite)(nil)

type genesisTestSuite struct {
	suite.Suite
	genesisState checkers.GenesisState
	players      []sdk.AccAddress
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(genesisTestSuite))
}

func (s *genesisTestSuite) SetupTest() {
	s.genesisState = *checkers.NewGenesisState()
	_, addrs := testutil.PrivKeyAddressPairs(2)
	s.players = addrs
}

func (s *genesisTestSuite) TestValidate() {

	testCases := map[string]struct {
		games   []checkers.IndexedStoredGame
		wantErr string
	}{
		"no games": {
			games:   []checkers.IndexedStoredGame{},
			wantErr: "",
		},
		"one game": {
			games: []checkers.IndexedStoredGame{
				{
					"1",
					checkers.NewStoredGame(s.players[0].String(), s.players[1].String(), time.Now()),
				},
			},
			wantErr: "",
		},
		"same game added twice": {
			games: []checkers.IndexedStoredGame{
				{
					"1",
					checkers.NewStoredGame(s.players[0].String(), s.players[1].String(), time.Now()),
				},
				{
					"1",
					checkers.NewStoredGame(s.players[0].String(), s.players[1].String(), time.Now()),
				},
			},
			wantErr: "duplicate game index",
		},
		"index too long": {
			games: []checkers.IndexedStoredGame{
				{
					testutil.RandLetters(checkers.MaxIndexLength + 1),
					checkers.NewStoredGame(s.players[0].String(), s.players[1].String(), time.Now()),
				},
			},
			wantErr: "index too long",
		},
		"game is invalid (opponents are the same)": {
			games: []checkers.IndexedStoredGame{
				{
					"1",
					checkers.NewStoredGame(s.players[0].String(), s.players[0].String(), time.Now()),
				},
			},
			wantErr: "cannot play against self",
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			s.genesisState.IndexedStoredGameList = tc.games
			err := s.genesisState.Validate()

			if tc.wantErr != "" {
				s.Error(err)
				s.Contains(err.Error(), tc.wantErr)
				return
			}

			s.NoError(err)
		})
	}
}
