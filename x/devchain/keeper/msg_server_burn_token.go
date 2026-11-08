package keeper

import (
	"context"

	"devchain/x/devchain/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BurnToken(ctx context.Context, msg *types.MsgBurnToken) (*types.MsgBurnTokenResponse, error) {
	creator, err := k.addressCodec.StringToBytes(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid creator address")
	}

	coin := sdk.NewCoin(msg.Denom, math.NewIntFromUint64(msg.Amount))
	if err := coin.Validate(); err != nil {
		return nil, errorsmod.Wrap(err, "invalid coin")
	}

	coins := sdk.NewCoins(coin)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, coins); err != nil {
		return nil, errorsmod.Wrap(err, "failed to send coins from creator to module")
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins); err != nil {
		return nil, errorsmod.Wrap(err, "failed to burn coins")
	}

	return &types.MsgBurnTokenResponse{}, nil
}
