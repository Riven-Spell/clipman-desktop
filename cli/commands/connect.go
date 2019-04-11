package commands

import (
	"fmt"
	"github.com/virepri/clipman-desktop/client"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
)

func connect(args []string) {
	client.Messages <- internal_command.Command{
		Cmd:    internal_command.CONNECT,
		Params: []string{"cli_request"},
	}

	if <-config.CLISuccess {
		fmt.Println("Successfully connected to " + config.ServerIP + "!")
	} else {
		fmt.Println("Failed to connect to " + config.ServerIP + ". Are you sure it's a valid IP or is serving?")
	}
}

func disconnect(args []string) {
	client.Messages <- internal_command.Command{
		Cmd:    internal_command.DISCONNECT,
		Params: []string{"cli_request"},
	}

	if <-config.CLISuccess {
		fmt.Println("Successfully disconnected.")
	} else {
		fmt.Println("Failed to disconnect. Were you even connected to start with?")
	}
}
