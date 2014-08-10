package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/yurufuwa/yurufuwa"
)

func main() {
	app := cli.NewApp()
	app.Name = "yurufuwa"
	app.Usage = "Admin tools to manage Yurufuwa organization."
	app.Commands = []cli.Command{
		*yurufuwa.MembersCommand(),
		*yurufuwa.CollaboratorsCommand(),
	}
	app.Run(os.Args)
}
