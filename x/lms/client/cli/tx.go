package cli

import (
	"clms/x/lms/types"
	"log"
	"time"

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
		NewAcceptLeaveRequestCmd(),
	)
	return txCmd
}

// CLi to add /register a Admin

func NewRegisterAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "RegisterAdmin ",
		Short:   "RegisterAdmin [address] [name]",
		Long:    "To register new admin",
		Example: "./lmsa tx lms RegisterAdmin cosmos1amgv60vqvq9elx53xyu6fgx37xuaqcz6fltctl ADMIN --from validator-key --chain-id testnet",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}

			address := args[0]
			name := args[1]

			msg := types.NewRegisterAdminRequest(address, name)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CLI to add new students (can only be done by a admin)
func NewAddStudentRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "AddStudent ",
		Short:   "AddStudent [Adminaddress] [name] [id] [address]",
		Long:    "To Add new students",
		Example: "./lmsa tx lms AddStudent cosmos1k353axl2gjevx6u6rc8e0gd3sg2uhhyymnhj9r student 12345 cosmos1amgv60vqvq9elx53xyu6fgx37xuaqcz6fltctl --from validator-key --chain-id testnet",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			adminaddress, _ := sdk.AccAddressFromBech32(args[0])
			student := types.Student{
				Name:    args[1],
				Id:      args[2],
				Address: args[3],
			}
			msg := types.NewAddStudentReq(adminaddress, student)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CLI to apply leave requests from students which can be accessed by the admin
func NewApplyLeaveRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ApplyLeave ",
		Short:   "ApplyLeave [Adminaddress][User Addesss] [reason] [From] [to] ",
		Long:    "To apply leave requests",
		Example: "./lmsa tx lms ApplyLeave cosmos1k353axl2gjevx6u6rc8e0gd3sg2uhhyymnhj9r cosmos1amgv60vqvq9elx53xyu6fgx37xuaqcz6fltctl flu 2009-Feb-17 2009-Feb-19 --from validator-key --chain-id testnet",
		Args:    cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// a := "12-02-2022"
			// b := "13-02-2022"
			const format = "2006-Jan-06"
			ssfrom, _ := time.Parse(format, args[3])
			to, _ := time.Parse(format, args[4])
			adminaddress, _ := sdk.AccAddressFromBech32(args[0])
			//leaves := []*types.Leave{}
			leave := &types.Leave{
				Address: args[1],
				Reason:  args[2],
				Sfrom:   &ssfrom,
				To:      &to,
				Status:  "undefined",
			}
			msg := types.NewApplyLeaveReq(adminaddress, leave)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CLI to accept/aprove the leave requests (can only be done by admin)
func NewAcceptLeaveRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "AcceptLeave ",
		Short:   "AceeptLeave [Adminaddress] [address] [leaveId] [Status] ",
		Long:    "To Accept leave requests",
		Example: "./lmsa tx lms AcceptLeave cosmos1k353axl2gjevx6u6rc8e0gd3sg2uhhyymnhj9r cosmos1amgv60vqvq9elx53xyu6fgx37xuaqcz6fltctl accepted 1 --from validator-key --chain-id testnet",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			adminaddress, _ := sdk.AccAddressFromBech32(args[0])
			address := args[1]
			leaveid := args[2]
			status := args[3]
			msg := types.NewAcceptLeaveReq(adminaddress, address, leaveid, status)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
