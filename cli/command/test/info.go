package test

import (
	"github.com/spf13/cobra"
	"storm/cli/command"
	"golang.org/x/net/context"
)


func newinfoCommand(stormCli *command.StormCli) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "test info",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInfo(stormCli, args)
		},
	}
}

func runInfo(stromcli *command.StormCli, args []string) error {
	ctx := context.Background()
	stromcli.Client().TestInfo(ctx)
	return nil
}
