package client

import (
	"fmt"
	"net"
)

func StartClient() {
	var connection net.Conn
	fmt.Println(connection) //TODO: handle connection stuff.

	for {
		select {
		case cmd := <- Messages:
			fmt.Println(cmd)
		case cmd := <- ExternalMessages:
			fmt.Println(cmd)
		}
	}
}
