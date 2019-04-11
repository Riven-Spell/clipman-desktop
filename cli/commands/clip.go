package commands

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/virepri/clipman-desktop/client"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
)

func clip(args []string) {
	if len(args) >= 1 {
		switch args[0] {
		case "force":
			if len(args) >= 2 {
				switch args[1] {
				case "push":
					client.Messages <- internal_command.Command{
						Cmd:    internal_command.PUSH_CLIP,
						Params: []string{"cli_request"},
					}

					if <-config.CLISuccess {
						fmt.Println("Successfully pushed the clipboard!")
					} else {
						fmt.Println("Couldn't push the clipboard. Are you connected or authorized?")
					}
				case "refresh":
					client.Messages <- internal_command.Command{
						Cmd:    internal_command.REFRESH_CLIP,
						Params: []string{"cli_request"},
					}

					if <-config.CLISuccess {
						fmt.Println("Successfully refreshed the clipboard!")
					} else {
						fmt.Println("Couldn't refresh the clipboard. Are you connected or authorized?")
					}
				}
			}
		case "recheck":
			if data, err := clipboard.ReadAll(); err == nil {
				if config.ClipboardContent != data {
					config.LockContent = true
					config.ClipboardContent = data

					client.Messages <- internal_command.Command{
						Cmd:    internal_command.PUSH_CLIP,
						Params: []string{"cli_request"},
					}

					if <-config.CLISuccess {
						fmt.Println("Clipboard was out of date! Pushed clipboard to server.")
					} else {
						fmt.Println("Clipboard was out of date. Failed to push to the server.")
					}
					config.LockContent = false
				}

				fmt.Println("Successfully performed clipboard recheck.")
			} else {
				fmt.Println(err.Error())
			}
		case "empty":
			if err := clipboard.WriteAll(""); err == nil {
				fmt.Println("Successfully emptied clipboard.")
			} else {
				fmt.Println(err.Error())
			}
		}
	}
}
