package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeRegisterAdmin = "register-admin"
	TypeAddStudent    = "add-student"
)

var (
	_ sdk.Msg = &RegisterAdminRequest{}
)

func NewRegisterAdminRequest(address string, name string) *RegisterAdminRequest {
	return &RegisterAdminRequest{
		Address: address,
		Name:    name,
	}
}

func (msg RegisterAdminRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg RegisterAdminRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{fromAddress}
}

func (msg RegisterAdminRequest) ValidateBasic() error {
	// if _, err := sdk.AccAddressFromBech32("hii"); err != nil {
	// 	return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	// }
	if msg.Address == "" {
		return Err1
	} else if msg.Name == "" {
		return Err2
	} else {
		return nil
	}

}

// Add Student
func NewAddStudentReq(accountAddr sdk.AccAddress, students []*Student) *AddStudentRequest {
	return &AddStudentRequest{
		Admin:    accountAddr.String(),
		Students: students,
	}
}

func (msg AddStudentRequest) GetSignBytes() []byte {
	return []byte{}
}

func (msg AddStudentRequest) Route() string {
	return RouterKey
}

func (msg AddStudentRequest) Type() string {
	return TypeAddStudent
}

func (msg AddStudentRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.Admin)
	// valAddr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg AddStudentRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	if msg.Admin == "" {
		return errors.New("admin address cannot be empty")
	} else if msg.Students == nil {
		return errors.New("students list cannot be empty, no students provided")
	}
	return nil
}
func NewApplyLeaveReq(accountAddr sdk.AccAddress, leaves []*Leave) *ApplyLeaveRequest {
	return &ApplyLeaveRequest{
		Admin:  accountAddr.String(),
		Leaves: leaves,
	}
}

func (msg ApplyLeaveRequest) GetSignBytes() []byte {
	return []byte{}
}

func (msg ApplyLeaveRequest) Route() string {
	return RouterKey
}

func (msg ApplyLeaveRequest) Type() string {
	return TypeAddStudent
}

func (msg ApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.Admin)
	// valAddr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg ApplyLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	if msg.Admin == "" {
		return errors.New("admin address cannot be empty")
	} else if msg.Leaves == nil {
		return errors.New("students list cannot be empty, no students provided")
	}
	return nil
}
func NewAcceptLeaveReq(accountAddr sdk.AccAddress, leaves []*Leave) *ApplyLeaveRequest {
	return &ApplyLeaveRequest{
		Admin:  accountAddr.String(),
		Leaves: leaves,
	}
}

func (msg AcceptLeaveRequest) GetSignBytes() []byte {
	return []byte{}
}

func (msg AcceptLeaveRequest) Route() string {
	return RouterKey
}

func (msg AcceptLeaveRequest) Type() string {
	return TypeAddStudent
}

func (msg AcceptLeaveRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.Admin)
	// valAddr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg AcceptLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	if msg.Admin == "" {
		return errors.New("admin address cannot be empty")
	}
	return nil
}
