package commands

import (
	"fmt"
	"github.com/virepri/clipman-desktop/config"
	"strconv"
)

func configCmd(args []string) {
	if len(args) >= 1 {
		switch args[0] {
		case "server":
			if len(args) >= 2 {
				config.ServerIP = args[1]
			} else {
				fmt.Println(config.ServerIP)
			}
		case "buffer":
			if len(args) >= 2 {
				if n, err := strconv.Atoi(args[1]); err == nil {
					config.Buffer = n

					fmt.Println("Successfully changed buffer size.")
				} else {
					fmt.Println("Bufffer size is an integer.")
				}
			} else {
				fmt.Println("buffer size:", config.Buffer)
			}
		case "save":
			config.SaveCFG()
			fmt.Println("Saved the config to the disk.")
		case "reload":
			config.LoadCFG()
			fmt.Println("Successfully reloaded the config.")
		}
	}
}
