package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
    "github.com/spf13/pflag"
    "storm/cli"
	"strings"
	"storm/cli/command"
	"storm/cli/command/commands"

	"storm-master/stormversion"
)

type daemonOptions struct {
	version      bool
	configFile   string
	flags        *pflag.FlagSet
	unixDomain   string
	host         string
	port         string
	logFile      string
	logLevel     string
}

func newDaemonCommand() *cobra.Command {
    opts := daemonOptions{
    	version: false,
    }

    cmd := &cobra.Command{
	Use:           "stormd",
	Short:         "A self-sufficient runtime for storm.",
	SilenceUsage:  true,
	SilenceErrors: true,
	Args:          cli.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
	    opts.flags = cmd.Flags()
	    return runDaemon(opts)
	},
    }
    cli.SetupRootCommand(cmd)

    flags := cmd.Flags()
    flags.BoolVarP(&opts.version, "version", "v", false, "Print version information and quit")
    flags.StringVarP(&opts.configFile, "configFile", "c","/etc/storm/daemon.json", "Daemon configuration file")
    flags.StringVarP(&opts.unixDomain,"unixDomain","u","/run/storm.sock","Set unix sock file")
    flags.StringVarP(&opts.host,"ip","i","","Which ip socket to listen")
	flags.StringVarP(&opts.port,"port","p","9090","Which port socket to listen")
	flags.StringVarP(&opts.logFile,"log","l","/var/log/storm.log","Log file location")
	flags.StringVarP(&opts.logLevel,"level","L","info","Log level(eg:panic,fatal,error,warn,info,debug)")

    return cmd
}

func runDaemon(opts daemonOptions) error {
    if opts.version {
	    showVersion()
	    return nil
    }
    daemonCli := NewDaemonCli()
    err := daemonCli.start(opts)
    return err
}

func showVersion() {
	fmt.Printf("Storm version %s, build %s\n", Version, GitCommit)
}

func newStormCommand(stormCli *command.StormCli) *cobra.Command {
	opts := &command.ClientOptions{}
	var flags *pflag.FlagSet

	cmd := &cobra.Command{
		Use:              "storm [OPTIONS] COMMAND [ARG...]",
		Short:            "A self-sufficient runtime for containers",
		SilenceUsage:     true,
		SilenceErrors:    true,
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Version {
				showVersion()
				return nil
			}
			return stormCli.ShowHelp(cmd, args)
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// flags must be the top-level command flags, not cmd.Flags()
			if err := stormCli.Initialize(opts); err != nil {
				return err
			}
			return nil
		},
	}
	cli.SetupRootCommand(cmd)

	flags = cmd.Flags()
	flags.BoolVarP(&opts.Version, "version", "v", false, "Print version information and quit")
	flags.StringVarP(&opts.UnixDomain, "unixDomain","d","unix:///run/storm.sock", "unix domain file location")
	flags.StringVarP(&opts.HostPort, "host","H","", "Daemon socket(s) to connect to (eg:tcp://127.0.0.1:9090)")
	commands.AddCommands(cmd, stormCli)

	return cmd
}

func main(){
	commandName := os.Args[0]
	if strings.HasSuffix(commandName,"stormd") {
		cmd := newDaemonCommand()
		cmd.SetOutput(os.Stdout)

		if err := cmd.Execute(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}else if strings.HasSuffix(commandName,"storm"){
		stormCli := command.NewStormCli()
		cmd := newStormCommand(stormCli)

		if err := cmd.Execute(); err != nil {
			if sterr, ok := err.(cli.StatusError); ok {
				if sterr.Status != "" {
					fmt.Println(sterr.Status)
				}
				if sterr.StatusCode == 0 {
					os.Exit(1)
				}
				os.Exit(sterr.StatusCode)
			}
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

