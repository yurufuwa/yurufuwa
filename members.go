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
			err := listMembers(client)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
}

func listMembers(client *github.Client) error {
	members, err := FetchMembers(client)

	if err != nil {
		return err
	}

	for index := range members {
		member := members[index]
		fmt.Println(*member.Login)
	}
	return nil
}

// FetchMembers returns Yurufuwa members as array of github.User
func FetchMembers(client *github.Client) ([]github.User, error) {
	members, _, err := client.Organizations.ListMembers("yurufuwa", &github.ListMembersOptions{})

	return members, err
}
