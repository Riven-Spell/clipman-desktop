package main

import (
	"fmt"
	"github.com/virepri/clipman-desktop/config"
	"github.com/virepri/clipman-desktop/monitor"
	"os/user"
)

func main() {
	if u, err := user.Current(); err == nil {
		config.CfgLocation = u.HomeDir + "/.config/clipman-desktop.cfg"
		fmt.Println(u.HomeDir)
	} else {
		fmt.Println(err.Error())
		return
	}

	config.LoadCFG()
	config.WaitGroup.Add(1)

	go monitor.StartMonitor()

	config.WaitGroup.Wait()
}
