# A Checkers Game Implementation for Cosmos SDK

## Usage
* Explain the commands available

## Installation
* Details on how to install this module

# TODO List

## Game Rules
* Ensure a player can't play against themselves
  * Red and Black wallets must be different

## Storage improvements
* use a more optimal data type for storage keys (not strings)

## Quality improvements
* Create a test suite for stored-game.go
* Add tests for the various validations performed
  * such as dupicate games in genesis
  * such as game ID length out of range

## CI Improvements
* Create a base image that can be used for tests - saves downloading all packages on every pipeline run
