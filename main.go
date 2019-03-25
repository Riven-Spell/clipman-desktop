package main

import (
	"fmt"
	"github.com/virepri/clipman-desktop/client"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
	"github.com/virepri/clipman-desktop/monitor"
	"os/user"
)

func main() {
	if u, err := user.Current(); err == nil {
		config.CfgLocation = u.HomeDir + "/.config/clipman-desktop.cfg"
		fmt.Println(u.HomeDir)
	} else {
		fmt.Println(err.Error())
		return
	}

	if config.LoadCFG() {
		client.Messages <- internal_command.Command{
			Cmd: internal_command.CONNECT,
		}

		client.Messages <- internal_command.Command{
			Cmd: internal_command.AUTH_USER,
		}
	}
	config.WaitGroup.Add(2)

	go monitor.StartMonitor()
	go client.StartClient()

	config.WaitGroup.Wait()
}
