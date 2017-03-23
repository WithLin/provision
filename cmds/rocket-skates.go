// Package main Rocket Skates Server
//
// An RestFUL API-driven Provisioner and DHCP server
//
package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/rackn/rocket-skates/server"
)

var c_opts server.ProgOpts

func main() {
	parser := flags.NewParser(&c_opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	server.Server(&c_opts)
}
