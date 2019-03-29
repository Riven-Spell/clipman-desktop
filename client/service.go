package client

import (
	"github.com/virepri/clipman-desktop/client/external-command"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
)

func StartClient() {
	defer config.WaitGroup.Done()

	for {
		select {
		case cmd := <-Messages:
			if v, ok := internal_command.Commands[cmd.Cmd]; ok {
				v(cmd.Params)
			}
		case cmd := <-ExternalMessages:
			if v, ok := external_command.Commands[cmd.Cmd]; ok {
				v(cmd.Params)
			}
		}
	}
}

func SetupChannels() {
	Messages = make(chan internal_command.Command)
	ExternalMessages = make(chan external_command.Command)
}
