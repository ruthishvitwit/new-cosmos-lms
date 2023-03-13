package keeper_test

import (
	"clms/x/lms/keeper"
	"clms/x/lms/types"
	"fmt"
	"sync"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
)

type TestSuite struct {
	suite.Suite
	ctx         sdk.Context
	stdntKeeper keeper.Keeper
	*assert.Assertions
	mu      sync.RWMutex
	require *require.Assertions
	t       *testing.T
}

func (s *TestSuite) SetupTest() {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	encCfg := simapp.MakeTestEncodingConfig()
	lmsKey := sdk.NewKVStoreKey(types.StoreKey)
	ctx := testutil.DefaultContext(lmsKey, sdk.NewTransientStoreKey("transient_test"))
	keeper := keeper.NewKeeper(lmsKey, encCfg.Codec)
	cms.MountStoreWithDB(lmsKey, storetypes.StoreTypeIAVL, db)
	s.Require().NoError(cms.LoadLatestVersion())
	s.stdntKeeper = keeper
	s.ctx = ctx
}

func (suite *TestSuite) T() *testing.T {
	suite.mu.RLock()
	defer suite.mu.RUnlock()
	return suite.t
}

// SetT sets the current *testing.T context.
func (suite *TestSuite) SetT(t *testing.T) {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	suite.t = t
	suite.Assertions = assert.New(t)
	suite.require = require.New(t)
}

// Require returns a require context for suite.
func (suite *TestSuite) Require() *require.Assertions {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	if suite.require == nil {
		suite.require = require.New(suite.T())
	}
	return suite.require
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

///////////////////// Register Admin Tests ////////////////////////

func (s *TestSuite) TestRegisterAdmin() {
	type registerAdminTest struct {
		arg1     types.RegisterAdminRequest
		expected error
	}

	var registerAdminTests = []registerAdminTest{
		{
			arg1: types.RegisterAdminRequest{
				Name:    "Hemanthsai",
				Address: sdk.AccAddress("abcdef").String(),
			},
			expected: nil,
		},
		{
			arg1: types.RegisterAdminRequest{
				Name:    "Sai",
				Address: sdk.AccAddress("sakjhfdd").String(),
			},
			expected: nil,
		},
		{
			arg1: types.RegisterAdminRequest{
				Name:    "Vishal",
				Address: "",
			},
			expected: types.Err1,
		},
		{
			arg1: types.RegisterAdminRequest{
				Name:    "",
				Address: sdk.AccAddress("kgjdk").String(),
			},
			expected: types.Err2,
		},
	}

	require := s.Require()
	for _, test := range registerAdminTests {

		if output := s.stdntKeeper.RegisterAdmin(s.ctx, &test.arg1); output != test.expected {
			require.Equal(test.expected, output)
		}
		// s.stdntKeeper.GetAdmin(s.ctx, sdk.AccAddress("sakjhfdd").String())
	}

}
func (s *TestSuite) TestAddStudent() {
	students := types.Student{
		Address: sdk.AccAddress("lms1").String(),
		Name:    "hemanth1",
		Id:      "1",
	}
	req := types.AddStudentRequest{
		Admin:    "Hemanthsai",
		Students: &students,
	}
	s.stdntKeeper.AddStudent(s.ctx, &req)
}

func (s *TestSuite) TestGetStudents() {
	s.TestAddStudent()
	res := s.stdntKeeper.GetStudents(s.ctx, &types.GetStudentRequest{})
	fmt.Println("Get Students Response: ")
	fmt.Println(res)
}
func (s *TestSuite) TestApplyLeave() {
	students := types.Leave{
		Address: "add.stud",
		Reason:  "sick",
		Leaveid: "KYFJJY",
		Status:  "pending",
	}
	req := types.ApplyLeaveRequest{
		Admin:  "Hemanthsai",
		Leaves: &students,
	}
	s.stdntKeeper.ApplyLeave(s.ctx, &req)
}
func (s *TestSuite) TestGetLeaves() {
	s.TestApplyLeave()
	res := s.stdntKeeper.GetLeaves(s.ctx, &types.GetLeavesRequest{})
	fmt.Println("Get Students Response: ")
	fmt.Println(res)
}

func (s *TestSuite) TestAcceptLeave() {
	s.TestApplyLeave()
	s.stdntKeeper.AcceptLeave(s.ctx, &types.AcceptLeaveRequest{Admin: "hi", Adress: "add.stud", LeaveId: "1", Status: "accepted"})
	fmt.Println("Get Students Response: ")
	req := s.stdntKeeper.GetLeaves(s.ctx, &types.GetLeavesRequest{})
	fmt.Println(req)
}
