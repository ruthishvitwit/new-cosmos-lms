package keeper

import (
	"clms/x/lms/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) QueryGetStudent(goCtx context.Context, req *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	result := k.GetStudents(ctx, req)
	//panic("eje")
	return &types.GetStudentResponse{Students: result}, nil
}
func (k Keeper) QueryGetLeaves(goCtx context.Context, req *types.GetLeavesRequest) (*types.GetLeavesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	result := k.GetLeaves(ctx, req)
	//panic("eje")
	return &types.GetLeavesResponse{Leaves: result}, nil
}
