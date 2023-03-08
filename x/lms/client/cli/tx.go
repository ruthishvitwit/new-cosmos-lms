package cli

import (
	"clms/x/lms/types"
	"log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func handleError(err error) {
	log.Fatal(err)
}

func TxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.Modulename,
		Short:                      "LMS",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 4,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		NewRegisterAdminCmd(),
		NewAddStudentRequestCmd(),
		NewApplyLeaveRequestCmd(),
	)
	return txCmd
}

func NewRegisterAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registeradmin [name] [address]",
		Short: "To register new admin",
		Long:  "To register new admin",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}

			address := args[0]
			name := args[1]

			msg := types.NewRegisterAdminRequest(address, name)
			// panic("called1")
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewAddStudentRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "addstudents",
		Short: "This is used to add new students",
		Long:  "This is used to add new students",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			adminaddress, _ := sdk.AccAddressFromBech32(args[0])
			students := []*types.Student{}
			student := &types.Student{
				Name:    args[1],
				Id:      args[2],
				Address: args[3],
			}
			students = append(students, student)

			msg := types.NewAddStudentReq(adminaddress, students)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func NewApplyLeaveRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aplyleave",
		Short: "...",
		Long:  "This is used to apply leaves",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// a := "12-02-2022"
			// b := "13-02-2022"
			// const format = "2022-Feb-07"
			// ssfrom, _ := time.Parse(format, "2022-Feb-07")
			// to, _ := time.Parse(format, args[3])
			adminaddress, _ := sdk.AccAddressFromBech32(args[0])
			leaves := []*types.Leave{}
			leave := &types.Leave{
				Address: args[1],
				Reason:  args[2],
				// Sfrom:   a,
				// To:      b,
			}
			leaves = append(leaves, leave)

			msg := types.NewApplyLeaveReq(adminaddress, leaves)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
