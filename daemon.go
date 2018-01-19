package main

import (
	"github.com/spf13/pflag"
	"storm/listeners"
	"github.com/sirupsen/logrus"
	"fmt"
	//"strings"
	"storm/api/server/router"
	apiserver "storm/api/server"
	"storm/api/server/router/test"
	"storm/log"
	"storm/api/server/middleware"
	"os/signal"
	"os"
	"storm/api/server/router/management"
)
type Config struct {

	Logging     bool
	CorsHeaders string
	Version     string
	SocketGroup string
	Hosts     []string `json:"hosts,omitempty"`
}

type DaemonCli struct {
	*Config
	configFile *string
	flags      *pflag.FlagSet
	api        *apiserver.Server
	//d               *Daemon
}


// NewDaemonCli returns a daemon CLI
func NewDaemonCli() *DaemonCli {
	return &DaemonCli{}
}

func (cli *DaemonCli) start(opts daemonOptions) (err error) {
	log.LogInit(opts.logFile,opts.logLevel)
	if cli.Config, err = loadDaemonCliConfig(opts); err != nil {
		return err
	}
	cli.configFile = &opts.configFile
	cli.flags = opts.flags

	serverConfig := &apiserver.Config{
		Logging:     true,
		Version:     Version,
	}

	api := apiserver.New(serverConfig)
	cli.api = api

	ls, err := listeners.Init("unix", opts.unixDomain)
	if err != nil {
		return err
	}
	api.Accept(opts.unixDomain, ls...)

	addr := opts.host + ":" + opts.port
	ls, err = listeners.Init("tcp", addr)
	if err != nil {
		return err
	}
	api.Accept(addr, ls...)


	logrus.WithFields(logrus.Fields{
		"version":    Version,
		"commit":     GitCommit,
	}).Info("storm daemon")

	// initMiddlewares needs cli.d to be populated. Dont change this init order.
	if err := cli.initMiddlewares(api, serverConfig); err != nil {
		logrus.Fatalf("Error creating middlewares: %v", err)
	}

	initRouter(api)

	// The serve API routine never exits unless an error occurs
	// We need to start it as a goroutine and wait on it so
	// daemon doesn't exit
	serveAPIWait := make(chan error)
	sinalWait := make(chan os.Signal)
	go api.Wait(serveAPIWait)
	signal.Notify(sinalWait)

	select {
		case errAPI := <- serveAPIWait:
			if errAPI != nil {
				return fmt.Errorf("Shutting down due to ServeAPI error: %v", errAPI)
			}
			case s := <- sinalWait:
				fmt.Println("Signal recive",s)
	}
	defer func() {
		api.Close()
	}()
	return nil
}

func loadDaemonCliConfig(opts daemonOptions) (*Config, error) {
	return nil,nil
}

func initRouter(s *apiserver.Server) {
	routers := []router.Router{
		//register router
		test.NewRouter(),
		management.NewRouter(),
	}
	s.InitRouter(routers...)
}


func (cli *DaemonCli) initMiddlewares(s *apiserver.Server, cfg *apiserver.Config) error {

	exp := middleware.NewTestMiddleware("hello world")
	s.UseMiddleware(exp)

	return nil
}




