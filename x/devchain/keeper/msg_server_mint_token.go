package keeper

import (
	"context"

	"devchain/x/devchain/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintToken(ctx context.Context, msg *types.MsgMintToken) (*types.MsgMintTokenResponse, error) {

	creator, err := k.addressCodec.StringToBytes(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid creator address")
	}

	k
	coin := sdk.NewCoin(msg.Denom, math.NewIntFromUint64(msg.Amount))
	if err := coin.Validate(); err != nil {
		return nil, errorsmod.Wrap(err, "invalid coin")
	}

	coins := sdk.NewCoins(coin)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return nil, errorsmod.Wrap(err, "failed to mint coins")
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, coins); err != nil {
		return nil, errorsmod.Wrap(err, "failed to send coins to creator")
	}
	
	return &types.MsgMintTokenResponse{}, nil
}
