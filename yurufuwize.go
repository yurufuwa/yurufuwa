package main

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/yurufuwa/tools/client"
)

func main() {
	client := client.CreateClient()

	members, _, err := client.Organizations.ListMembers("yurufuwa", &github.ListMembersOptions{})

	if err != nil {
		fmt.Println(err)
	}

	for index := range members {
		fmt.Println(*members[index].Login)
	}
}
