package command

import (
	"net/http"
	"runtime"
	"storm/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// StormCli represents the storm command line client.
// Instances of the client can be returned from NewStormCli.
type StormCli struct {
	keyFile         string
	client          *client.Client
	hasExperimental bool
	defaultVersion  string
}

type ClientOptions struct {
	Version      bool
	Flags        *pflag.FlagSet
	UnixDomain   string
	HostPort     string
}

// HasExperimental returns true if experimental features are accessible.
func (cli *StormCli) HasExperimental() bool {
	return cli.hasExperimental
}

// DefaultVersion returns api.defaultVersion of STORM_API_VERSION if specified.
func (cli *StormCli) DefaultVersion() string {
	return cli.defaultVersion
}

// Client returns the APIClient
func (cli *StormCli) Client() *client.Client {
	return cli.client
}

// ShowHelp shows the command help.
func (cli *StormCli) ShowHelp(cmd *cobra.Command, args []string) error {
	cmd.HelpFunc()(cmd, args)
	return nil
}

// Initialize the stormCli runs initialization that must happen after command
// line flags are parsed.
func (cli *StormCli) Initialize(opts *ClientOptions) error {
	var err error
	cli.client, err = NewAPIClient(opts)
	if err != nil {
		return err
	}
	return nil
}

// NewStormCli returns a StormCli instance with IO output and error streams set by in, out and err.
func NewStormCli() *StormCli {
	return &StormCli{}
}

// NewAPIClient creates a new APIClient from command line flags
func NewAPIClient(opts *ClientOptions) (*client.Client, error) {
	var host string
	customHeaders := map[string]string{}
	customHeaders["User-Agent"] = UserAgent()

	if len(opts.HostPort) > 0{
		host = opts.HostPort
	}else{
		host = opts.UnixDomain
	}

	httpClient,err,proto,addr := newHTTPClient(host)
	if err != nil {
		return &client.Client{}, err
	}

	return client.NewClient(host,httpClient,customHeaders,proto,addr)
}

func newHTTPClient(host string) (*http.Client, error,string,string) {
	tr := &http.Transport{
	}
	proto, addr, _, err := client.ParseHost(host)
	if err != nil {
		return nil, err,"",""
	}

	client.ConfigureTransport(tr, proto, addr)

	return &http.Client{
		Transport: tr,
	}, nil,proto,addr
}

// UserAgent returns the user agent string used for making API requests
func UserAgent() string {
	return "storm-Client/" +  " (" + runtime.GOOS + ")"
}

