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
	app.Commands = []cli.Command{*yurufuwa.MembersCommand()}

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}
	app.Run(os.Args)
}
