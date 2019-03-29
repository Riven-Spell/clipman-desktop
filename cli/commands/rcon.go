package commands

import (
	"fmt"
	"github.com/virepri/clipman-desktop/client"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
)

func rcon(args []string) {
	client.Messages <- internal_command.Command{
		Cmd: internal_command.RCON_COMMAND,
		Params: args,
	}

	if <- config.CLISuccess {
		fmt.Println("Successfully issued rcon command.")
	} else {
		fmt.Println("Failed to issue rcon command. Are you connected to the server?")
	}
}
