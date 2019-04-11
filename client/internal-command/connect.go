package internal_command

import (
	"crypto/tls"
	"fmt"
	"github.com/virepri/clipman-desktop/config"
	"net"
)

func Connect(args []string) {
	if !config.ServerUsesTLS {
		fmt.Println("WARNING: Do not handle sensitive data without TLS enabled. This is unsafe and prone to MITM attacks.")

		if c, err := net.Dial("tcp", config.ServerIP); err == nil {
			config.Connection = c

			PushClip(make([]string, 0))

			if len(args) > 0 && args[0] == "cli_request" {
				config.CLISuccess <- true
			}
		} else {
			fmt.Println(err.Error())
			if len(args) > 0 && args[0] == "cli_request" {
				config.CLISuccess <- false
			}
		}
	} else {
		if config.TLSConfig.InsecureSkipVerify {
			fmt.Println("WARNING: Using TLS without verification of certificate. This is unsafe and prone to MITM attacks.")
			fmt.Println("Add the server's certificate to your certpool if you trust it (or get the server to switch to a cert authority)")
		}

		if c, err := tls.Dial("tcp", config.ServerIP, config.TLSConfig); err == nil {
			config.Connection = c

			PushClip(make([]string, 0))

			if len(args) > 0 && args[0] == "cli_request" {
				config.CLISuccess <- true
			}
		} else {
			fmt.Println(err.Error())
			if len(args) > 0 && args[0] == "cli_request" {
				config.CLISuccess <- false
			}
		}
	}
}

func Disconnect(args []string) {
	if config.Connection != nil {
		if err := config.Connection.Close(); err == nil {
			config.AdminPerms = false
			config.UserPerms = false

			config.Connection = nil

			if len(args) > 0 && args[0] == "cli_request" {
				config.CLISuccess <- true
			}
			return
		}
	}

	if len(args) > 0 && args[0] == "cli_request" {
		config.CLISuccess <- false
	}
}
