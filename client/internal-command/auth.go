package internal_command

import (
	"fmt"
	"github.com/virepri/clipman-desktop/config"
)

func AuthUser(args []string) {
	buffer := []byte{4, 10}
	buffer = append(buffer, []byte(config.UserHash)...)
	buffer = append(buffer, 0)

	if config.Connection != nil {
		if _, err := config.Connection.Write(buffer); err != nil {
			fmt.Println(err)
		}

		//TODO: handle success/failure conditions
	} else if args[0] == "cli_request" {
		config.CLISuccess <- false
	}
}

func AuthAdmin(args []string) {
	buffer := []byte{0, 10}
	buffer = append(buffer, []byte(config.AdminHash)...)
	buffer = append(buffer, 0)

	if config.Connection != nil {
		if _, err := config.Connection.Write(buffer); err != nil {
			fmt.Println(err)
		}

		//TODO: handle success/failure conditions.
	} else if args[0] == "cli_request" {
		config.CLISuccess <- false
	}
}
