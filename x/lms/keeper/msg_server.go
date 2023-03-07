package keeper

import (
	"clms/x/lms/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

type msgServer struct {
	Keeper
	types.UnimplementedMsgServer
}

func (k Keeper) AdminRegister(ctx context.Context, req *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.RegisterAdmin(sdkCtx, req)
	return &types.RegisterAdminResponse{}, nil
}
func (k Keeper) AddStudents(ctx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.AddStudent(sdkCtx, req)
	return &types.AddStudentResponse{}, nil
}
