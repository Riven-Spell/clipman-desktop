package monitor

import (
	"github.com/atotto/clipboard"
	"os"
	"time"
)

var ClipboardContent string

func StartMonitor() {
	ticker := time.NewTicker(time.Second / 2)

	for  {
		<- ticker.C

		if data, err := clipboard.ReadAll(); err == nil {
			if ClipboardContent != data {
				ClipboardContent = data

				//TODO: ping the web client.
			}
		} else {
			os.Exit(0)
		}
	}
}
