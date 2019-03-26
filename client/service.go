package client

import (
	"fmt"
	"github.com/virepri/clipman-desktop/client/external-command"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
)

func StartClient() {
	defer config.WaitGroup.Done()

	//var connection net.Conn
	//fmt.Println(connection) //TODO: handle connection stuff.

	for {
		select {
		case cmd := <- Messages:
			fmt.Println(cmd)
		case cmd := <- ExternalMessages:
			fmt.Println(cmd)
		}
	}
}

func SetupChannels() {
	Messages = make(chan internal_command.Command)
	ExternalMessages = make(chan external_command.Command)
}