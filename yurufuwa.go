package main

import (
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/yurufuwa/yurufuwa/client"
	"github.com/yurufuwa/yurufuwa/commands"
)

// Command のインターフェイス
type Command interface {
	Name() string
	Run(*github.Client)
}

func main() {
	commands := []Command{&commands.Members{}}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("サブコマンドを指定してください: member")
		os.Exit(1)
	}

	for _, cmd := range commands {
		if cmd.Name() == args[1] {
			client := client.CreateClient()
			cmd.Run(client)
			os.Exit(0)
		}
	}

	fmt.Printf("%s サブコマンドは存在しません\n", args[1])
	os.Exit(1)
}
