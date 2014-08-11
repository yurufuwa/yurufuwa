package yurufuwa

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
)

// MeetupsCommand makes members subcommand.
func MeetupsCommand() *cli.Command {
	return &cli.Command{
		Name:  "meetups",
		Usage: "Show upcoming yurufuwa meetups.",
		Action: func(c *cli.Context) {
			client := CreateClient()
			err := listMeetups(client)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
}

func listMeetups(client *github.Client) error {
	option := &github.IssueListByRepoOptions{
		State:     "open",
		Sort:      "created",
		Direction: "desc",
	}
	meetups, _, err := client.Issues.ListByRepo("yurufuwa", "meetups", option)
	if err != nil {
		return err
	}

	for index := range meetups {
		meetup := meetups[index]
		fmt.Printf(`# %s
%s

%s

`,
			*meetup.Title,
			*meetup.Body,
			*meetup.HTMLURL,
		)
	}
	return nil
}
