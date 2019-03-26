package internal_command

import (
	"fmt"
	"github.com/virepri/clipman-desktop/cli/commands"
	"github.com/virepri/clipman-desktop/config"
	"net"
)

func Connect(args []string) {
	if c, err := net.Dial("tcp", config.ServerIP); err == nil {
		config.Connection = c

		if args[0] == "cli_request" {
			commands.Success <- true
		}
	} else {
		fmt.Println(err.Error())
		if args[0] == "cli_request" {
			commands.Success <- false
		}
	}
}