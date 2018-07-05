package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

var (
	githubKey      = strings.TrimSpace(strings.Replace(os.Getenv("GITHUB_APP_KEY"), "#", "\n", -1))
	githubAppID, _ = strconv.Atoi(os.Getenv("GITHUB_APP_ID"))
)

func createGithubClient(installationID int64) (*github.Client, error) {
	itr, err := ghinstallation.New(http.DefaultTransport, githubAppID, int(installationID), []byte(githubKey))
	if err != nil {
		return nil, err
	}

	return github.NewClient(&http.Client{Transport: itr}), nil
}
