package internal_command

import (
	"github.com/virepri/clipman-desktop/config"
)

func PushClip(args []string) {
	buffer := CmdToBytes(Command{Cmd: PUSH_CLIP, Params: []string{config.ClipboardContent}})

	if config.Connection != nil && config.UserPerms {
		_, _ = config.Connection.Write(buffer)

		if len(args) > 0 && args[0] == "cli_request" {
			config.CLISuccess <- true
		}
	} else if len(args) > 0 && args[0] == "cli_request" {
		config.CLISuccess <- false
	}
}

func RefreshClip(args []string) {
	buffer := CmdToBytes(Command{Cmd: REFRESH_CLIP})

	if config.Connection != nil && config.UserPerms {
		_, _ = config.Connection.Write(buffer)

		if len(args) > 0 && args[0] == "cli_request" {
			config.CLISuccess <- true
		}
	} else if len(args) > 0 && args[0] == "cli_request" {
		config.CLISuccess <- false
	}
}
