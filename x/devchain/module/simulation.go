package devchain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	devchainsimulation "devchain/x/devchain/simulation"
	"devchain/x/devchain/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	devchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&devchainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgMintToken          = "op_weight_msg_devchain"
		defaultWeightMsgMintToken int = 100
	)

	var weightMsgMintToken int
	simState.AppParams.GetOrGenerate(opWeightMsgMintToken, &weightMsgMintToken, nil,
		func(_ *rand.Rand) {
			weightMsgMintToken = defaultWeightMsgMintToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintToken,
		devchainsimulation.SimulateMsgMintToken(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgBurnToken          = "op_weight_msg_devchain"
		defaultWeightMsgBurnToken int = 100
	)

	var weightMsgBurnToken int
	simState.AppParams.GetOrGenerate(opWeightMsgBurnToken, &weightMsgBurnToken, nil,
		func(_ *rand.Rand) {
			weightMsgBurnToken = defaultWeightMsgBurnToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnToken,
		devchainsimulation.SimulateMsgBurnToken(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
