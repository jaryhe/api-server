package test
import (
	"storm/cli/command"
	"github.com/spf13/cobra"
)

// NewNodeCommand returns a cobra command for `node` subcommands
func NewTestCommand(stormCli *command.StormCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "just test function",
		RunE:  stormCli.ShowHelp,
	}
	cmd.AddCommand(
		newinfoCommand(stormCli),
	)
	return cmd
}

