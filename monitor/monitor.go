package monitor

import (
	"github.com/atotto/clipboard"
	"os"
	"time"
)

var clipboardContent string

func StartMonitor() {
	ticker := time.NewTicker(time.Second / 2)

	for  {
		<- ticker.C

		if data, err := clipboard.ReadAll(); err == nil {
			if clipboardContent != data {
				clipboardContent = data

				//TODO: ping the web client.
			}
		} else {
			os.Exit(0)
		}
	}
}
