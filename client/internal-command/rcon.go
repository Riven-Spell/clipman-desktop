package internal_command

import "github.com/virepri/clipman-desktop/config"

func Rcon(args []string) {
	if config.Connection != nil {
		buff := []byte{1, 10}

		for _, v := range args {
			buff = append(buff, append([]byte(v), 10)...)
		}

		buff[len(buff)-1] = 0

		_, _ = config.Connection.Write(buff)

		config.CLISuccess <- true
	} else {
		config.CLISuccess <- false
	}
}
