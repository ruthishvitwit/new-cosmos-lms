package keeper

import (
	"clms/x/lms/types"
	"fmt"
	"log"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) RegisterAdmin(ctx sdk.Context, RegisterAdmin *types.RegisterAdminRequest) error {

	if RegisterAdmin.Name == "" {
		return types.Err1
	} else if RegisterAdmin.Address == "" {
		return types.Err2
	} else {
		store := ctx.KVStore(k.storeKey)
		marshalRegisterAdmin, err := k.cdc.Marshal(RegisterAdmin)
		if err != nil {
			log.Fatal(err)
		}
		store.Set(types.AdminKey(RegisterAdmin.Address), marshalRegisterAdmin)
		return nil
	}
}

func (k Keeper) GetAdmin(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	fmt.Println(types.AdminKey(id))
	v := store.Get(types.AdminKey(id))
	fmt.Println(v)
}

func (k Keeper) AddStudent(ctx sdk.Context, addStudent *types.AddStudentRequest) string {
	students := addStudent.Students
	store := ctx.KVStore(k.storeKey)
	marshalAddStudent, err := k.cdc.Marshal(students)
	if err != nil {
		log.Fatal(err)
	}
	store.Set(types.StudentKey(students.Address), marshalAddStudent)
	return "Students Added Successfully"
}
func (k Keeper) ApplyLeave(ctx sdk.Context, applyleave *types.ApplyLeaveRequest) string {
	leave := applyleave.Leaves
	store := ctx.KVStore(k.storeKey)

	leaveid := store.Get(types.LeaveCounterKey(applyleave.Leaves.Address))
	leaveId, _ := strconv.Atoi(string(leaveid))
	if leaveid == nil {
		leaveId = 0

	}
	leaveId++
	leave.Leaveid = strconv.Itoa(leaveId)
	store.Set(types.LeaveCounterKey(applyleave.Leaves.Address), []byte(strconv.Itoa(leaveId)))
	marshalaplylv, err := k.cdc.Marshal(leave)
	if err != nil {
		log.Fatal(err)
	}
	store.Set(types.LeaveKey(applyleave.Leaves.Address, leaveId), marshalaplylv)

	return " leave applied Successfully"
}
func (k Keeper) GetStudents(ctx sdk.Context, getStudents *types.GetStudentRequest) []*types.Student {
	store := ctx.KVStore(k.storeKey)
	var students []*types.Student
	itr := sdk.KVStorePrefixIterator(store, types.SKey)
	for ; itr.Valid(); itr.Next() {
		var t types.Student
		k.cdc.Unmarshal(itr.Value(), &t)
		students = append(students, &t)
	}
	return students
}
func (k Keeper) GetaStudent(ctx sdk.Context, getaStudent *types.GetaStudentRequest) *types.Student {
	store := ctx.KVStore(k.storeKey)
	v := store.Get(types.StudentKey(getaStudent.Id))
	var id types.Student
	k.cdc.Unmarshal(v, &id)
	return &id
}
func (k Keeper) GetaStudentleave(ctx sdk.Context, getaStudent *types.GetaStudentRequest) *types.Student {
	store := ctx.KVStore(k.storeKey)
	v := store.Get(types.StudentKey(getaStudent.Id))
	var id types.Student
	k.cdc.Unmarshal(v, &id)
	return &id
}
func (k Keeper) GetLeaves(ctx sdk.Context, getLeaves *types.GetLeavesRequest) []*types.Leave {
	store := ctx.KVStore(k.storeKey)

	var leaves []*types.Leave
	itr := sdk.KVStorePrefixIterator(store, types.LKey)
	for ; itr.Valid(); itr.Next() {
		var t types.Leave
		k.cdc.Unmarshal(itr.Value(), &t)
		leaves = append(leaves, &t)
	}
	return leaves
}
func (k Keeper) AcceptLeave(ctx sdk.Context, StatusD *types.AcceptLeaveRequest) string {
	stud := StatusD
	store := ctx.KVStore(k.storeKey)
	id, _ := strconv.Atoi(string(stud.LeaveId))
	v := store.Get(types.LeaveKey(stud.Adress, id))
	var ki types.Leave
	k.cdc.Unmarshal(v, &ki)
	var ap types.Leave
	fmt.Println("d ", ki, "sagrag", ki.Address)
	ap.Address = ki.Address
	ap.Leaveid = ki.Leaveid
	ap.Reason = ki.Reason
	ap.Status = stud.Status
	ap.Sfrom = ki.Sfrom
	ap.To = ki.To
	fmt.Println(ap)
	store.Delete(types.LeaveKey(stud.Adress, id))
	c, _ := k.cdc.Marshal(&ap)
	store.Set(types.LeaveKey(stud.Adress, id), c)
	marshalAddStudent, err := k.cdc.Marshal(stud)
	if err != nil {
		log.Fatal(err)
	}
	store.Set(types.AcceptLeaveKey(stud.Adress), marshalAddStudent)
	return "Students Added Successfully"
}
