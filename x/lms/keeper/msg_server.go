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
func (k Keeper) ApplyLeaves(ctx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.ApplyLeave(sdkCtx, req)
	return &types.ApplyLeaveResponse{}, nil
}
func (k Keeper) LeaveAccept(ctx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.AcceptLeave(sdkCtx, req)
	return &types.AcceptLeaveResponse{}, nil
}
