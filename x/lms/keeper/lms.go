package keeper

import (
	"clms/x/lms/types"
	"fmt"
	"log"

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
	for _, stud := range students {
		marshalAddStudent, err := k.cdc.Marshal(stud)
		if err != nil {
			log.Fatal(err)
		}
		store.Set(types.StudentKey(stud.Address), marshalAddStudent)
	}
	return "Students Added Successfully"
}
func (k Keeper) GetStudents(ctx sdk.Context, getStudents *types.GetStudentRequest) []*types.Student {
	store := ctx.KVStore(k.storeKey)

	var students []*types.Student
	itr := store.Iterator(types.SKey, nil)
	for ; itr.Valid(); itr.Next() {
		var t types.Student
		k.cdc.Unmarshal(itr.Value(), &t)
		students = append(students, &t)
	}
	return students
}
