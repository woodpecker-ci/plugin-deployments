package forge

import (
	"context"

	"codeberg.org/woodpecker-plugins/go-plugin"
	"github.com/xanzy/go-gitlab"
)

type Gitlab struct {
	*gitlab.Client
}

func NewGitlab(url, token string) (Forge, error) {
	git, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		return nil, err
	}

	return &Gitlab{git}, nil
}

func (g *Gitlab) CreateEnvironment(ctx context.Context, repo plugin.Repository, url, name string) error {
	// Your code here
	return nil
}

func (g *Gitlab) RemoveEnvironment(ctx context.Context, repo plugin.Repository, url, name string) error {
	// Your code here
	return nil
}
