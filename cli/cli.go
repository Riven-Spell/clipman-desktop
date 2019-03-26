package cli

import (
	"bufio"
	"fmt"
	"github.com/virepri/clipman-desktop/cli/commands"
	"github.com/virepri/clipman-desktop/config"
	"os"
	"strings"
)

var Reader *bufio.Reader

func StartCLI() {
	defer config.WaitGroup.Done()

	fmt.Println("clipman-desktop version ", config.Version)
	fmt.Println("Enter \"help\" for information on various commands.")

	Reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input := strings.Split(readLine(), " ")

		if v, ok := commands.Aliases[input[0]]; ok {
			v(input[1:])
		} else {
			fmt.Println(input[0], "is not an available command. Check help.")
		}
	}
}

func readLine() string {
	if str, err := Reader.ReadString('\n'); err == nil {
		return strings.Trim(str, "\n")
	} else {
		fmt.Println(err.Error())
		return ""
	}
}

func SetupChannels() {
	config.CLISuccess = make(chan bool)
}
