package internal_command

import (
	"fmt"
	"github.com/virepri/clipman-desktop/config"
	"net"
)

func Connect(args []string) {
	if c, err := net.Dial("tcp", config.ServerIP); err == nil {
		config.Connection = c

		if args[0] == "cli_request" {
			config.CLISuccess <- true
		}
	} else {
		fmt.Println(err.Error())
		if args[0] == "cli_request" {
			config.CLISuccess <- false
		}
	}
}

func Disconnect(args []string) {
	if config.Connection != nil {
		if err := config.Connection.Close(); err == nil {
			config.AdminPerms = false
			config.UserPerms = false

			if args[0] == "cli_request" {
				config.CLISuccess <- true
			}
			return
		}
	}

	if args[0] == "cli_request" {
		config.CLISuccess <- false
	}
}
