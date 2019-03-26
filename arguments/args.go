package arguments

import "os"

func CheckArgs() {
	for _, v := range os.Args {
		switch v {
		case "--admin":
			AutoAdmin = true
		case "--auth":
			AutoAuth = true
		case "--connect":
			AutoConnect = true
		}
	}
}
