// Package testutil provides utilities for test data generation, setup and other common testing needs.
// Some functions are gracefully borrowed from Niburu's test utils:
// https://github.com/NibiruChain/nibiru/blob/main/x/common/testutil/sample.go#L39
package testutil

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PrivKeyAddressPairs generates (deterministically) a total of n private keys
// and addresses.
func PrivKeyAddressPairs(n int) (keys []cryptotypes.PrivKey, addrs []sdk.AccAddress) {
	r := rand.New(rand.NewSource(12345)) // make the generation deterministic
	keys = make([]cryptotypes.PrivKey, n)
	addrs = make([]sdk.AccAddress, n)
	for i := 0; i < n; i++ {
		secret := make([]byte, 32)
		_, err := r.Read(secret)
		if err != nil {
			panic("Could not read randomness")
		}
		keys[i] = secp256k1.GenPrivKeyFromSecret(secret)
		addrs[i] = sdk.AccAddress(keys[i].PubKey().Address())
	}
	return
}

type TypeLatin struct {
	Letters    string
	CapLetters string
	Numbers    string
}

var Latin = TypeLatin{
	Letters:    "abcdefghijklmnopqrstuvwxyz",
	CapLetters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	Numbers:    "0123456789",
}

// RandLetters generates a random string of length n using the Latin alphabet.
func RandLetters(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = Latin.Letters[rand.Intn(len(Latin.Letters))]
	}
	return string(b)
}
