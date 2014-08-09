package main

import (
	"bufio"
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/howeyc/gopass"
	"github.com/jmcvetta/napping"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// 設定ファイルがあれば読み込む
	var accessToken string
	var err error
	conf := ".yurufuwa.conf"
	buf, err := ioutil.ReadFile(conf)
	if err != nil {
		// 認証して accessToken を取得する
		accessToken, err = fetchAccessToken()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(conf, []byte(accessToken), 0644)
	} else {
		// ファイルから accessToken を読み込む
		accessToken = fmt.Sprintf("%s", buf)
	}

	client := createClient(accessToken)

	members, _, err := client.Organizations.ListMembers("yurufuwa", &github.ListMembersOptions{})
	for index := range members {
		fmt.Println(members[index])
	}
}

func fetchAccessToken() (string, error) {
	scan := func() string {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}

	fmt.Print("Username: ")
	login := scan()
	fmt.Print("Password: ")
	password := strings.TrimSpace(fmt.Sprintf("%s", gopass.GetPasswd()))
	fmt.Print("Two Factor Auth: ")
	tfaToken := scan()

	payload := struct {
		Scopes []string `json:"scopes"`
		Note   string   `json:"note"`
	}{
		Scopes: []string{"repo", "public_repo", "read:org"},
		Note:   "Yurufuwize",
	}

	res := struct {
		Id        int
		Url       string
		Scopes    []string
		Token     string
		App       map[string]string
		Note      string
		NoteUrl   string `json:"note_url"`
		UpdatedAt string `json:"updated_at"`
		CreatedAt string `json:"created_at"`
	}{}

	e := struct {
		Message string
		Errors  []struct {
			Resource string
			Field    string
			Code     string
		}
	}{}

	header := &http.Header{}
	header.Add("X-GitHub-OTP", tfaToken)

	s := napping.Session{
		Userinfo: url.UserPassword(login, password),
		Header:   header,
	}

	url := "https://api.github.com/authorizations"

	resp, err := s.Post(url, &payload, &res, &e)
	if resp.Status() == 201 {
		return res.Token, nil
	} else {
		fmt.Println("Bad response status from Github server")
		fmt.Printf("\t Status:  %v\n", resp.Status())
		fmt.Printf("\t Message: %v\n", e.Message)
		fmt.Printf("\t Errors: %v\n", e.Message)
		return "", err
	}
}

func createClient(accessToken string) *github.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: accessToken},
	}
	return github.NewClient(t.Client())
}
