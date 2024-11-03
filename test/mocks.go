package mocks

//go:generate mockery --name=BinaryCodec --srcpkg github.com/cosmos/cosmos-sdk/codec --outpkg mock_test
//go:generate mockery --name=Codec --srcpkg cosmossdk.io/core/address --outpkg mock_test
//go:generate mockery --name=KVStoreService --srcpkg cosmossdk.io/core/store --outpkg mock_test
//go:generate mockery --name=KVStore --srcpkg cosmossdk.io/core/store --outpkg mock_test
