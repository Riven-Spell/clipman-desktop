package external_command

import (
	"github.com/atotto/clipboard"
	"github.com/virepri/clipman-desktop/config"
	"strings"
)

func SetClip(args []string) {
	config.LockContent = true
	if err := clipboard.WriteAll(strings.Join(args, "\n")); err == nil {
		config.ClipboardContent = strings.Join(args, "\n")
	}
	config.LockContent = false
}
