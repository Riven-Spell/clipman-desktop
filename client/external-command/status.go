package external_command

import "github.com/virepri/clipman-desktop/config"

func Success(args []string) {
	config.CLISuccess <- true
}

func Failure(args []string) {
	config.CLISuccess <- false
}

//It's not like these are going to be used any other ways.
