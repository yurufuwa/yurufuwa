package yurufuwa

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
)

// MembersCommand makes members subcommand.
func MembersCommand() *cli.Command {
	return &cli.Command{
		Name:  "members",
		Usage: "Show yurufuwa members.",
		Action: func(c *cli.Context) {
			client := CreateClient()
			listMembers(client)
		},
	}
}

func listMembers(client *github.Client) {
	members, _, err := client.Organizations.ListMembers("yurufuwa", &github.ListMembersOptions{})

	if err != nil {
		fmt.Println(err)
	}

	for index := range members {
		member := members[index]
		fmt.Println(*member.Login)
	}
}
