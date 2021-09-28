package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/ksbomj/sfkp/services/order/cmd"

	"fmt"
	"log"
	"os"
)

var opts struct {
	ServerCmd cmd.ServerCommand `command:"server"`
}

var version = "unknown"

func main() {
	fmt.Printf("ORDER-SERVICE ver: %s\n\n", version)

	logger := log.New(os.Stdout, "ORDER-SERVICE ", log.Lmicroseconds|log.Llongfile|log.Lshortfile)

	if err := run(logger); err != nil {
		fmt.Printf("Cannot execute a program: %s\n", err)
		os.Exit(1)
	}
}

func run(logger *log.Logger) error {
	p := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash|flags.HelpFlag)

	p.CommandHandler = func(command flags.Commander, args []string) error {
		// commands implement CommonOptionsCommander to allow passing set of extra options defined for all commands
		c := command.(cmd.CommonOptionsCommander)
		c.SetCommon(cmd.CommonOpts{
			Version: version,
			Logger:  logger,
		})

		err := c.Execute(args)
		if err != nil {
			logger.Printf("failed with %+v", err)
		}
		return err
	}

	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			return err
		}
	}

	return nil
}
