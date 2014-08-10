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

					err = addCollaborators(client, owner, repos)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
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

					err = removeCollaborators(client, owner, repos)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
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

func addCollaborators(client *github.Client, owner, repos string) error {
	members, err := FetchMembers(client)
	if err != nil {
		return err
	}

	for _, user := range members {
		client.Repositories.AddCollaborator(owner, repos, *user.Login)
		fmt.Println(*user.Login + "has successfully added to your repository!!")
	}
	return nil
}

func removeCollaborators(client *github.Client, owner, repos string) error {
	members, err := FetchMembers(client)
	if err != nil {
		return err
	}

	for _, user := range members {
		client.Repositories.RemoveCollaborator(owner, repos, *user.Login)
		fmt.Println(*user.Login + "has successfully removed from your repository!!")
	}
	return nil
}
