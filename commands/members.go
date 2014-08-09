package commands

import (
	"fmt"

	"github.com/google/go-github/github"
)

// Members サブコマンド
type Members struct {
}

// Name コマンドの名前を返す
func (m *Members) Name() string {
	return "members"
}

// Run メンバー一覧を表示
func (m *Members) Run(client *github.Client) {
	displayMembers(client)
}

func displayMembers(client *github.Client) {
	members, _, err := client.Organizations.ListMembers("yurufuwa", &github.ListMembersOptions{})

	if err != nil {
		fmt.Println(err)
	}

	for index := range members {
		member := members[index]
		fmt.Printf("%d\t%s\n", *member.ID, *member.Login)
	}
}
