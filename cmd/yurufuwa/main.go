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

	client := yurufuwa.CreateClient()
	app.Commands = []cli.Command{
		{
			Name:  "members",
			Usage: "Show yurufuwa members.",
			Action: func(c *cli.Context) {
				yurufuwa.ListMembers(client)
			},
		},
	}

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}
	app.Run(os.Args)
}
