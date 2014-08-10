package yurufuwa

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"strings"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
	"github.com/howeyc/gopass"
	"github.com/jmcvetta/napping"
)

// CreateClient は認証済みの github.Client を返します
func CreateClient() *github.Client {
	// 設定ファイルがあれば読み込む
	var accessToken string
	var err error

	user, _ := user.Current()
	conf := user.HomeDir + "/.yurufuwa.conf"
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

	return oauthClient(accessToken)
}

func oauthClient(accessToken string) *github.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: accessToken},
	}
	return github.NewClient(t.Client())
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
		ID        int
		URL       string
		Scopes    []string
		Token     string
		App       map[string]string
		Note      string
		NoteURL   string `json:"note_url"`
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
	if err != nil {
		return "", err
	}
	if resp.Status() == 201 {
		return res.Token, nil
	}

	fmt.Println("Bad response status from Github server")
	fmt.Printf("\t Status:  %v\n", resp.Status())
	fmt.Printf("\t Message: %v\n", e.Message)
	fmt.Printf("\t Errors: %v\n", e.Message)
	return "", errors.New("既に Yurufuwize という Personal API Token が存在する場合は削除してください。")
}
