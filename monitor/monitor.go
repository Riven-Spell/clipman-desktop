package monitor

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/virepri/clipman-desktop/client"
	"github.com/virepri/clipman-desktop/client/internal-command"
	"github.com/virepri/clipman-desktop/config"
	"time"
)

var ClipboardContent string

func StartMonitor() {
	ticker := time.NewTicker(time.Second / 10)
	defer config.WaitGroup.Done()

	for  {
		<- ticker.C

		if data, err := clipboard.ReadAll(); err == nil {
			if ClipboardContent != data {
				ClipboardContent = data

				client.Messages <- internal_command.Command{
					Cmd:internal_command.PUSH_CLIP,
				}
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}
