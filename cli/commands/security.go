package commands

import (
	"fmt"
	"github.com/virepri/clipman-desktop/client"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
)

func security(args []string) {
	if len(args) >= 1 {
		/*
			TODO: implement all security features planned:

			-Encryption across all clients
				-Perhaps with a simple plain-text password to generate the private/public keys, and then store them in the config?
				-Use machine-specific salt to securely store the priv/pub keys
			-Toggle to toggle the sync on/off (DONE)
			-Move TLS setup into here (DONE)
		*/

		switch args[0] {
		case "sync":
			if len(args) >= 2 {
				switch args[1] {
				case "toggle":
					config.NoSync = !config.NoSync
					if !config.NoSync {
						client.Messages <- internal_command.Command{
							Cmd:    internal_command.REFRESH_CLIP,
							Params: []string{"cli_request"},
						}

						if <-config.CLISuccess {
							fmt.Println("Refreshed clipboard after sync resumed.")
						} else {
							fmt.Println("Failed to refresh clipboard after sync resumed.")
						}
					}
				}
			}

			fmt.Println("Sync on?:", !config.NoSync)
		case "tls":
			if len(args) >= 2 {
				switch args[1] {
				case "toggle":
					config.ServerUsesTLS = !config.ServerUsesTLS
				case "toggleVerify":
					config.TLSConfig.InsecureSkipVerify = !config.TLSConfig.InsecureSkipVerify
				}
			}

			fmt.Println("TLS Enabled?", config.ServerUsesTLS)
			fmt.Println("TLS Insecure?:", config.TLSConfig.InsecureSkipVerify)

			if config.TLSConfig.InsecureSkipVerify {
				fmt.Println("WARNING: TLS verification is off. This is susceptible to man-in-the-middle attacks and not recommended.")
				fmt.Println("Do not handle sensitive data.")
				fmt.Println("If you trust this server's certificate, add it to your cert pool.")
			}

			if !config.ServerUsesTLS {
				fmt.Println("WARNING: TLS is disabled. This is susceptible to man-in-the-middle attacks and not recommended.")
				fmt.Println("Do not handle sensitive data.")
			}
		}
	} else {
		//TODO: print the status of all security measures
	}
}
