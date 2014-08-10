package yurufuwa

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
)

// CollaboratorsCommand makes members subcommand.
func CollaboratorsCommand() *cli.Command {
	return &cli.Command{
		Name: "collaborators",
		Usage: `Manage yurufuwa members for your repo.

    ## コラボレータに追加
    $ yurufuwa collaborators add your/repos

    ## コラボレータから削除
    $ yurufuwa collaborators remove your/repos
    `,
		Subcommands: []cli.Command{
			{
				Name:  "add",
				Usage: "Add yurufuwa members to repo.",
				Action: func(c *cli.Context) {
					client := CreateClient()
					owner, repos, err := extractReposName(c.Args())
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					addedMembers, err := addCollaborators(client, owner, repos)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					fmt.Printf("These members have been added to %s/%s!!\n", owner, repos)
					for _, member := range addedMembers {
						fmt.Printf(" - %s\n", *member.Login)
					}
				},
			},
			{
				Name:  "remove",
				Usage: "Remove yurufuwa members from repo.",
				Action: func(c *cli.Context) {
					client := CreateClient()
					owner, repos, err := extractReposName(c.Args())
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					removedMembers, err := removeCollaborators(client, owner, repos)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					fmt.Printf("These members have been removed from %s/%s!!\n", owner, repos)
					for _, member := range removedMembers {
						fmt.Printf(" - %s\n", *member.Login)
					}
				},
			},
		},
	}
}

func extractReposName(args cli.Args) (string, string, error) {
	var org, name string
	if len(args) > 0 {
		slice := strings.Split(args[0], "/")
		org = slice[0]
		name = slice[1]
	} else {
		return "", "", errors.New("リポジトリを指定してください")
	}

	return org, name, nil
}

func addCollaborators(client *github.Client, owner, repos string) ([]github.User, error) {
	members, err := FetchMembers(client)
	if err != nil {
		return nil, err
	}

	var addedMembers []github.User
	for _, user := range members {
		_, err := client.Repositories.AddCollaborator(owner, repos, *user.Login)
		if err == nil {
			addedMembers = append(addedMembers, user)
		}
	}
	return addedMembers, nil
}

func removeCollaborators(client *github.Client, owner, repos string) ([]github.User, error) {
	members, err := FetchMembers(client)
	if err != nil {
		return nil, err
	}

	var removedMembers []github.User
	for _, user := range members {
		_, err := client.Repositories.RemoveCollaborator(owner, repos, *user.Login)
		if err == nil {
			removedMembers = append(removedMembers, user)
		}
	}
	return removedMembers, nil
}
