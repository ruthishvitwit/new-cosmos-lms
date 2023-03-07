package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	Err1 = sdkerrors.Register(Modulename, 1, "Admin name cannot be empty")
	Err2 = sdkerrors.Register(Modulename, 2, "Admin address cannot be empty")
)
