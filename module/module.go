package module

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	checkers "github.com/lukevenediger/checkers"
	"github.com/lukevenediger/checkers/keeper"
)

var (
	_ module.AppModuleBasic = AppModule{}
	_ module.HasGenesis     = AppModule{}
	_ appmodule.AppModule   = AppModule{}
)

// ConsensusVersion defines the current version of the module's consensus
const ConsensusVersion = 1

type AppModule struct {
	cdc    codec.Codec
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
) AppModule {
	return AppModule{
		cdc:    cdc,
		keeper: keeper,
	}
}

// NewAppModuleBasic creates a new AppModuleBasic object
func NewAppModuleBasic(m AppModule) module.AppModuleBasic {
	return module.CoreAppModuleBasicAdaptor(m.Name(), m)
}

// Name returns the checkers module's name
func (AppModule) Name() string { return checkers.ModuleName }

// RegisterLegacyAminoCodec is present to meet the interface requirements but is not implemented
func (AppModule) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// noop - not supported
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the checkers module
func (AppModule) RegisterGRPCGatewayRoutes(
	clientCtx client.Context,
	mux *gwruntime.ServeMux,
) {
	if err := checkers.RegisterQueryHandlerClient(context.Background(), mux, checkers.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// RegisterInterfaces registers the module's interface types
func (AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	checkers.RegisterInterfaces(registry)
}

// ConsensusVersion returns the module's consensus version
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// RegisterServices registers a gRPC query service to handle module-specific queries
func (am AppModule) RegisterServices(cfg module.Configurator) {
	checkers.RegisterCheckersTorramServer(
		cfg.MsgServer(),
		keeper.NewMsgServerImpl(am.keeper),
	)
	checkers.RegisterQueryServer(
		cfg.QueryServer(),
		keeper.NewQueryServerImpl(am.keeper),
	)
}

// DefaultGenesis returns the default genesis state in raw bytes
func (AppModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(checkers.NewGenesisState())
}

// ValidateGenesis validates the provided genesis state
func (AppModule) ValidateGenesis(
	cdc codec.JSONCodec,
	_ client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	var data checkers.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", checkers.ModuleName, err)
	}
	return data.Validate()
}

// InitGenesis initialises the module's state from a given genesis state
func (am AppModule) InitGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
	data json.RawMessage,
) {
	var genesisState checkers.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	if err := am.keeper.InitGenesis(ctx, &genesisState); err != nil {
		panic(fmt.Sprintf("failed to initialise %s genesis state: %v", checkers.ModuleName, err))
	}
}

// ExportGenesis exports the genesis state in JSON format
func (am AppModule) ExportGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
) json.RawMessage {
	gs, err := am.keeper.ExportGenesis(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to export %s genesis state: %v", checkers.ModuleName, err))
	}

	return cdc.MustMarshalJSON(gs)
}
