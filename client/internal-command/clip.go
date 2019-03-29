package internal_command

import (
	"github.com/virepri/clipman-desktop/config"
)

func PushClip(args []string) {
	buffer := []byte{3, 10}
	buffer = append(buffer, []byte(config.ClipboardContent)...)
	buffer = append(buffer, 0)

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
	buffer := []byte{2, 0}

	if config.Connection != nil && config.UserPerms {
		_, _ = config.Connection.Write(buffer)

		if len(args) > 0 && args[0] == "cli_request" {
			config.CLISuccess <- true
		}
	} else if len(args) > 0 && args[0] == "cli_request" {
		config.CLISuccess <- false
	}
}
