package webhook

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/google/go-github/v41/github"
)

// Handle webhook from github
func Handle(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("Handling pull request")
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
	if !(action == "synchronize" || action == "opened" || action == "edited" || action == "reopened") {
		return nil
	}

	return checkTitle(ctx, payload)
}

func checkTitle(ctx context.Context, payload *github.PullRequestEvent) error {
	client, err := createGithubClient(payload.GetInstallation().GetID())
	if err != nil {
		return fmt.Errorf("can not create GitHub client: %w", err)
	}

	repoName := payload.GetRepo().GetName()
	repoOwner := payload.GetRepo().GetOwner().GetLogin()
	title := payload.GetPullRequest().GetTitle()
	commit := payload.GetPullRequest().GetHead().GetSHA()

	content, _, _, err := client.Repositories.GetContents(ctx, repoOwner, repoName, ".entitle.yml", &github.RepositoryContentGetOptions{Ref: commit})
	if err != nil {
		return fmt.Errorf("can not read .entitle.yml file from repo: %w", err)
	}

	configStr, err := content.GetContent()
	if err != nil {
		return fmt.Errorf("can not decode content: %w", err)
	}

	cfg, err := parseConfig(configStr)
	if err != nil {
		return fmt.Errorf("can not parse config: %w", err)
	}

	state, messages := cfg.checkTitle(title)
	msgIdx := rand.Intn(len(messages))
	status := &github.RepoStatus{
		Context:     github.String("PR Title"),
		State:       &state,
		Description: &messages[msgIdx],
	}
	_, _, err = client.Repositories.CreateStatus(ctx, repoOwner, repoName, commit, status)
	return err
}
