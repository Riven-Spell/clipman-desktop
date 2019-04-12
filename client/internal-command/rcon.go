package internal_command

import "github.com/virepri/clipman-desktop/config"

func Rcon(args []string) {
	if config.Connection != nil {
		buff := CmdToBytes(Command{Cmd: RCON_COMMAND, Params: args})

		buff[len(buff)-1] = 0

		_, _ = config.Connection.Write(buff)

		config.CLISuccess <- true
	} else {
		config.CLISuccess <- false
	}
}
