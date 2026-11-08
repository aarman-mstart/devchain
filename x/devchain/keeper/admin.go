package keeper

import (
	"context" 
	errorsmod "cosmossdk.io/errors"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const DefaultAdminAddress = "cosmos1l9zg8wns3nxltycfyljp328c2wulwqct8mq62k"
func (k Keeper) GetAdminAddress(ctx context.Context) (string, error) {
	params, err := k.Params.Get(ctx)

	if err != nil {
		return "", errorsmod.Wrap(err, "failed to get admin address")
	}
	if params.Admin != "" {
		return params.Admin, nil
	}
	return DefaultAdminAddress, nil
}
func (k Keeper) SetAdminAddress(ctx context.Context, admin string) error {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return err
	}

	if admin == "" {
    return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "admin address cannot be empty")
	}

	params.Admin = admin
	if err := params.Validate(); err != nil {
		return err
	}

	if err := k.Params.Set(ctx, params); err != nil {
		return err
	}

	return nil
}

func (k Keeper) IsAdmin(ctx context.Context, address string) (bool, error) {
	admin, err := k.GetAdminAddress(ctx)
	if err != nil {
		return false, err
	}
	return address == admin, nil
}
