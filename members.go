package yurufuwa

import (
	"fmt"

	"github.com/google/go-github/github"
)

// ListMembers lists yurufuwa orgnization members.
func ListMembers(client *github.Client) {
	members, _, err := client.Organizations.ListMembers("yurufuwa", &github.ListMembersOptions{})

	if err != nil {
		fmt.Println(err)
	}

	for index := range members {
		member := members[index]
		fmt.Println(*member.Login)
	}
}
