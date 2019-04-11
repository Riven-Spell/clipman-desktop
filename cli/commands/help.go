package commands

import "fmt"

func help(args []string) {
	if len(args) >= 1 {
		if v, ok := HelpList[args[0]]; ok {
			fmt.Println(args[0] + ":\n" + v + "\n")
		} else {
			fmt.Println(args[0], "is not a valid command.")
		}
	} else {
		for k, v := range HelpList {
			fmt.Println(k + ":\n" + v + "\n")
		}
	}
}
