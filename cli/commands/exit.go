package commands

import (
	"fmt"
	"github.com/virepri/clipman-desktop/config"
	"os"
)

func exit(args []string) {
	fmt.Println("Shutting down server.")
	config.SaveCFG()
	os.Exit(0)
}
