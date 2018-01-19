package commands

import (

	"storm/cli/command"
	"storm/cli/command/test"
	"github.com/spf13/cobra"
)

// AddCommands adds all the commands from cli/command to the root command
func AddCommands(cmd *cobra.Command, stormCli *command.StormCli) {
	cmd.AddCommand(
		//add sub commands
		test.NewTestCommand(stormCli),
	)
}

func hide(cmd *cobra.Command) *cobra.Command {
	cmdCopy := *cmd
	cmdCopy.Hidden = true
	cmdCopy.Aliases = []string{}
	return &cmdCopy
}
