package main

import (
	"fmt"
	"github.com/virepri/clipman-desktop/arguments"
	"github.com/virepri/clipman-desktop/cli"
	"github.com/virepri/clipman-desktop/client"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
	"github.com/virepri/clipman-desktop/monitor"
	"os/user"
)

func main() {
	arguments.CheckArgs()
	client.SetupChannels()
	cli.SetupChannels()

	if u, err := user.Current(); err == nil {
		config.CfgLocation = u.HomeDir + "/.config/clipman-desktop.cfg"
	} else {
		fmt.Println(err.Error())
		return
	}

	if config.LoadCFG() {
		if arguments.AutoConnect {
			client.Messages <- internal_command.Command{
				Cmd: internal_command.CONNECT,
			}

			if arguments.AutoAuth {
				client.Messages <- internal_command.Command{
					Cmd: internal_command.AUTH_USER,
				}

				if arguments.AutoAdmin {
					client.Messages <- internal_command.Command{
						Cmd: internal_command.AUTH_ADMIN,
					}
				}
			}
		}
	}
	config.WaitGroup.Add(3)

	go monitor.StartMonitor()
	go client.StartClient()
	go cli.StartCLI()

	config.WaitGroup.Wait()
}
