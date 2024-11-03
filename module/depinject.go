package module

import (
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	modulev1 "github.com/lukevenediger/checkers/api/module/v1"
	"github.com/lukevenediger/checkers/keeper"
)

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements depinject.OnePerModuleType interface
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule tags this module as an App Module
func (am AppModule) IsAppModule() {}

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Cdc          codec.Codec
	StoreService store.KVStoreService
	AddressCodec address.Codec

	Config *modulev1.Module
}

type ModuleOutputs struct {
	depinject.Out

	Module appmodule.AppModule
	Keeper keeper.Keeper
}

// ProvideModule sets up the checkers module for injection into the app
func ProvideModule(in ModuleInputs) ModuleOutputs {
	authority := authtypes.NewModuleAddress("gov")
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	k := keeper.NewKeeper(in.Cdc, in.AddressCodec, in.StoreService, authority.String())
	m := NewAppModule(in.Cdc, k)

	return ModuleOutputs{
		Module: m,
		Keeper: k,
	}
}
