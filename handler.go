package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

func handleHook(w http.ResponseWriter, r *http.Request) {
	rawPayload, err := github.ValidatePayload(r, []byte(os.Getenv("GITHUB_HOOK_SECRET")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	hookType := github.WebHookType(r)
	payload, err := github.ParseWebHook(hookType, rawPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if hookType == "pull_request" {
		err = handlePullRequest(r.Context(), payload.(*github.PullRequestEvent))
	} else if hookType == "ping" {
		// do nothing
	} else {
		err = fmt.Errorf("Unknown event: %s", hookType)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "OK")
}

func handlePullRequest(ctx context.Context, payload *github.PullRequestEvent) error {
	action := payload.GetAction()
	if !(action == "synchronized" || action == "opened" || action == "edited" || action == "reopened") {
		return nil
	}

	client, err := createGithubClient(payload.GetInstallation().GetID())
	if err != nil {
		return err
	}

	repoName := payload.GetRepo().GetName()
	repoOwner := payload.GetRepo().GetOwner().GetLogin()
	title := payload.GetPullRequest().GetTitle()
	commit := payload.GetPullRequest().GetHead().GetSHA()

	content, _, _, err := client.Repositories.GetContents(ctx, repoOwner, repoName, ".github/entitle.yml", nil)
	if err != nil {
		return err
	}

	configStr, err := content.GetContent()
	if err != nil {
		return err
	}

	fmt.Printf("Got %s %s\n%s\n\n", title, commit, configStr)
	return nil
}
