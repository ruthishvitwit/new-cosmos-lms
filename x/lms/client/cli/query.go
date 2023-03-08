package cli

import (
	"clms/x/lms/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func QueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.Modulename,
		Short:                      "LMS",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 4,
		RunE:                       client.ValidateCmd,
	}
	queryCmd.AddCommand(
		GetStudentsCmd(),
		GetLeavesCmd(),
		GetaStudentCmd(),
	)
	return queryCmd
}

func GetStudentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getStudents",
		Short: "getStudents ",
		Long:  "get list of all Students",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			//getStudentRequest := &types.GetStudentRequest{}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.QueryGetStudent(cmd.Context(), &types.GetStudentRequest{})
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
func GetLeavesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getLeaves",
		Short: "getLeaves ",
		Long:  "get list of all leaves",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			//getStudentRequest := &types.GetStudentRequest{}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.QueryGetLeaves(cmd.Context(), &types.GetLeavesRequest{})
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetaStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getaStudent",
		Short: "getStudents ",
		Long:  "get list of all Students",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			id := args[0]
			//getStudentRequest := &types.GetStudentRequest{}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.QueryGetaStudent(cmd.Context(), &types.GetaStudentRequest{
				Id: id,
			})
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
