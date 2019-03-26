package commands

import (
	"fmt"
	"github.com/virepri/clipman-desktop/client"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
)

func Connect(args []string) {
	client.Messages <- internal_command.Command{
		Cmd: internal_command.CONNECT,
		Params: []string{"cli_request"},
	}

	if <- Success {
		fmt.Println("Successfully connected to " + config.ServerIP + "!")
	} else {
		fmt.Println("Failed to connect to " + config.ServerIP + ". Are you sure it's a valid IP or is serving?")
	}
}