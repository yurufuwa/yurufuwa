package yurufuwa

import (
	"github.com/google/go-github/github"
	"github.com/satococoa/github-auth/client"
)

// CreateClient は認証済みの github.Client を返します
func CreateClient() *github.Client {
	return client.CreateClient("yurufuwa")
}
