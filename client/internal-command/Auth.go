package internal_command

import (
	"fmt"
	"github.com/virepri/clipman-desktop/config"
)

func AuthUser(args []string) {
	buffer := []byte{4, 10}
	buffer = append(buffer, []byte(args[0])...)
	buffer = append(buffer, 0)

	if _, err := config.Connection.Write(buffer); err != nil {
		fmt.Println(err)
	}
}

func AuthAdmin(args []string) {
	buffer := []byte{0, 10}
	buffer = append(buffer, []byte(args[0])...)
	buffer = append(buffer, 0)

	if _, err := config.Connection.Write(buffer); err != nil {
		fmt.Println(err)
	}

	//TODO: handle success/failure conditions.
}
