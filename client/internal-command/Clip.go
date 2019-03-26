package internal_command

import (
	"github.com/virepri/clipman-desktop/cli/commands"
	"github.com/virepri/clipman-desktop/config"
	"github.com/virepri/clipman-desktop/monitor"
)

func PushClip(args []string) {
	buffer := []byte{2, 10}
	buffer = append(buffer, []byte(monitor.ClipboardContent)...)
	buffer = append(buffer, 0)

	if config.Connection != nil && commands.UserPerms {
		_, _ = config.Connection.Write(buffer)
		commands.Success <- true
	} else {
		commands.Success <- false
	}
}

func RefreshClip(args []string) {
	buffer := []byte{2, 0}

	if config.Connection != nil && commands.UserPerms {
		_, _ = config.Connection.Write(buffer)
		commands.Success <- true
	} else {
		commands.Success <- false
	}
}
