package client

import (
	"fmt"
	"github.com/virepri/clipman-desktop/client/external-command"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
	"time"
)

func StartClient() {
	defer config.WaitGroup.Done()

	go func() {
		for {
			if config.Connection == nil {
				time.Sleep(100 * time.Millisecond)
				continue
			}

			buffer := make([]byte, config.Buffer)

			if _, err := config.Connection.Read(buffer); err != nil {
				continue
			}

			ExternalMessages <- external_command.ParseCmd(buffer)
		}
	}()

	for {
		select {
		case cmd := <-Messages:
			if v, ok := internal_command.Commands[cmd.Cmd]; ok {
				v(cmd.Params)
			}
		case cmd := <-ExternalMessages:
			if v, ok := external_command.Commands[cmd.Cmd]; ok {
				fmt.Println(cmd)
				v(cmd.Params)
			}
		}
	}
}



func SetupChannels() {
	Messages = make(chan internal_command.Command)
	ExternalMessages = make(chan external_command.Command)
}
