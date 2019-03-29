package external_command

import (
	"github.com/atotto/clipboard"
	"github.com/virepri/clipman-desktop/config"
)

func SetClip(args []string) {
	config.LockContent = true
	if err := clipboard.WriteAll(args[0]); err == nil {
		config.ClipboardContent = args[0]
	}
	config.LockContent = false
}
