package forge

import (
	"context"
	"errors"

	"codeberg.org/woodpecker-plugins/go-plugin"
)

type Forge interface {
	CreateDeployment(ctx context.Context, repo plugin.Repository, name, url string, commit *plugin.Commit) error
	RemoveDeployment(ctx context.Context, repo plugin.Repository, name string) error
}

func GetForge(forge plugin.Forge, token string) (Forge, error) {
	switch forge.Type {
	case "gitlab":
		var err error
		_forge, err := NewGitlab(forge.Link, token)
		return _forge, err
	case "github":
		var err error
		_forge, err := NewGithub(forge.Link, token)
		return _forge, err
	default:
		return nil, errors.New("unsupported forge type")
	}
}
