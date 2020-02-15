package webhook

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v29/github"
)

var (
	githubKey      = strings.TrimSpace(strings.Replace(os.Getenv("GITHUB_APP_KEY"), "#", "\n", -1))
	githubAppID, _ = strconv.ParseInt(os.Getenv("GITHUB_APP_ID"), 10, 64)
)

func createGithubClient(installationID int64) (*github.Client, error) {
	itr, err := ghinstallation.New(http.DefaultTransport, githubAppID, installationID, []byte(githubKey))
	if err != nil {
		return nil, err
	}

	return github.NewClient(&http.Client{Transport: itr}), nil
}
