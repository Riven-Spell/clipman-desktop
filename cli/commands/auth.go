package commands

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/virepri/clipman-desktop/client"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
)

var Success chan bool

func auth(args []string) {
	if len(args) >= 1 {
		switch args[0] {
		case "user":
			client.Messages <- internal_command.Command{
				Cmd: internal_command.AUTH_USER,
				Params: []string{"cli_request"},
			}

			if <- Success {
				fmt.Println("Successfully authenticated as a user.")
			} else {
				fmt.Println("Failed to authenticate, maybe your password is incorrect?")
			}
		case "admin":
			client.Messages <- internal_command.Command{
				Cmd: internal_command.AUTH_ADMIN,
				Params: []string{"cli_request"},
			}

			if <- Success {
				fmt.Println("Successfully authenticated as an admin.")
			} else {
				fmt.Println("Failed to authenticate, maybe your password is incorrect?")
			}
		case "pass":
			if len(args) >= 3 {
				strptr := &config.UserHash

				if args[1] == "admin" {
					strptr = &config.AdminHash
				}

				hash := sha256.New()
				hash.Write([]byte(args[2]))
				*strptr = hex.EncodeToString(hash.Sum(nil))
				fmt.Println("Updated", args[1], "password.")
			}
		}
	}
}