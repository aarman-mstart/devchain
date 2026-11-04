package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewParams creates a new Params instance.
func NewParams(admin string) Params {
	return Params{Admin: admin}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams("")
}

// Validate validates the set of params.
func (p Params) Validate() error {
    if p.Admin == "" {
		return nil
	}
	if _, err := sdk.AccAddressFromBech32(p.Admin); err != nil {
		return errorsmod.Wrapf(err, "invalid admin address: %s", p.Admin)
	}
	return nil
}
