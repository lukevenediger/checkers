# A Checkers Game Implementation for Cosmos SDK

* Module provides a way to play checkers on a Cosmos SDK app chain
* Demonstrates how to read, validate, and write data to the chain
* Demonstrates how to create test suites for queries and transactions

## About a Checkers Game
* A checkers game contains all the state for a game, including:
  * the game's start time
  * the game status
  * the checkers board with current pieces and positions
  * the addresses of the players

* a checkers game can be forfeited
  * only games that are in the Ready or Playing state may be forfeited
  * only the player who's turn it is may forfeit the game

* a game is completed, either through forfeit, win, or expiry
  * the end time is recorded upon completion

## Commands available

### Transactions:
* CheckersCreateGm
  * create a new Checkers game
  * takes in a game ID - this must be unique
  * takes in addresses of two players - they must not be the same
  * CLI: $chain_bin tx checkers-torram create $GAME_ID $PLAYER_1 $PLAYER_2
* ForfeitGame
  * forfeit an active game
  * this transaction must be signed by the player wishing to forfeit
  * the signer must be the player who's turn it is in the game
  * CLI: $chain_bin tx checkers-torram forfeit $GAME_ID

### Queries:
* GetGame
  * retrieves a game by its game ID
  * an empty result is returned if no game is found
  * CLI: $chain_bin query checkers-torram get-game $GAME_ID

# Contributions
* We'd love to get your contributions
* Please check out [CONTRIBUTING.md](CONTRIBUTING.md) for more information

# TODO List

## Storage improvements
* use a more optimal data type for storage keys (not strings)

## Quality improvements
* Create a test suite for stored-game.go
* Create a test suite for genesis.go
  * especially to confirm the duplicate game check

## CI Improvements
* Protect the main branch and disallow direct commits
  * everything must be submitted via a PR
