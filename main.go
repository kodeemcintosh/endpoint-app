package main

import (
	"os"
	"strings"
)

var (
	commandFile string
	root        Node
)

func main() {
	root = Node{
		Name:     "root",
		Depth:    0,
		Children: []*Node{},
	}

	ParseFlags()

	if commandFile != "" {
		var strCommands = ReadCommandFile(commandFile)

		for _, c := range strCommands {
			str := strings.Fields(c)

			cmd := Command{
				Type: strings.ToUpper(str[0]),
			}

			if len(str) != 0 {
				cmd.Args = str[1:]
			}

			cmd.Execute()
		}
		// } else if len(os.Args) >= 2 {
	} else {
		cmd := &Command{
			Type: strings.ToUpper(os.Args[1]),
			Args: os.Args[2:],
		}

		cmd.Execute()
	}
}
