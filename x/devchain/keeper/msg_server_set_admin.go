package keeper

import (
	"context"

	"devchain/x/devchain/types"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

func (k msgServer) SetAdmin(ctx context.Context, msg *types.MsgSetAdmin) (*types.MsgSetAdminResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid creator address")
	}

	isAdmin, err := k.IsAdmin(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to check if the creator is admin")
	}
	if !isAdmin {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the admin can set a new admin")
	}

	err = k.SetAdminAddress(ctx, msg.Address)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to set admin address")
	}

	return &types.MsgSetAdminResponse{}, nil
}
