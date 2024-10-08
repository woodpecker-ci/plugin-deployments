package forge

import (
	"context"
	"net/url"

	"codeberg.org/woodpecker-plugins/go-plugin"
	"github.com/google/go-github/v66/github"
	"golang.org/x/oauth2"
)

type Github struct {
	*github.Client
}

func NewGithub(_url, token string) (Forge, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.TODO(), ts)
	client := github.NewClient(tc)
	client.BaseURL, _ = url.Parse(_url)
	return &Github{client}, nil
}

func (g *Github) getDeployment(ctx context.Context, repo plugin.Repository, name string) (*github.Deployment, error) {
	envs, _, err := g.Repositories.ListDeployments(ctx, repo.Owner, repo.Name, &github.DeploymentsListOptions{
		Environment: name,
	})
	if err != nil {
		return nil, err
	}

	if len(envs) == 1 {
		return envs[0], nil
	}

	return nil, nil
}

func (g *Github) getEnvironment(ctx context.Context, repo plugin.Repository, name string) (*github.Deployment, error) {
	envs, _, err := g.Repositories.ListDeployments(ctx, repo.Owner, repo.Name, &github.DeploymentsListOptions{
		Environment: name,
	})
	if err != nil {
		return nil, err
	}

	if len(envs) == 1 {
		return envs[0], nil
	}

	return nil, nil
}

func (g *Github) CreateDeployment(ctx context.Context, repo plugin.Repository, name, url string, metadata *plugin.Metadata) error {
	environment, err := g.getEnvironment(ctx, repo, name)
	if err != nil {
		return err
	}

	if environment == nil {
		_, _, err := g.Repositories.CreateUpdateEnvironment(ctx, repo.Owner, repo.Name, name, &github.CreateUpdateEnvironment{})
		if err != nil {
			return err
		}
	}

	deployment, err := g.getDeployment(ctx, repo, name)
	if err != nil {
		return err
	}

	if deployment == nil {
		deployment, _, err = g.Repositories.CreateDeployment(ctx, repo.Owner, repo.Name, &github.DeploymentRequest{
			Ref:              github.String(metadata.Curr.Ref),
			Environment:      github.String(name),
			Description:      github.String("🚀 Deployment created by woodpecker"),
			RequiredContexts: &[]string{}, // empty array to skip checks
		})
		if err != nil {
			return err
		}
	}

	_, _, err = g.Repositories.CreateDeploymentStatus(ctx, repo.Owner, repo.Name, *deployment.ID, &github.DeploymentStatusRequest{
		Environment:    github.String(name),
		EnvironmentURL: github.String(url),
		State:          github.String("success"),
		AutoInactive:   github.Bool(true),
		LogURL:         github.String(metadata.Pipeline.Link),
	})

	return err
}

func (g *Github) RemoveDeployment(ctx context.Context, repo plugin.Repository, name string) error {
	deployment, err := g.getDeployment(ctx, repo, name)
	if err != nil {
		return err
	}

	if deployment != nil {
		_, err := g.Repositories.DeleteDeployment(ctx, repo.Owner, repo.Name, 1)
		if err != nil {
			return err
		}
	}

	_, err = g.Repositories.DeleteEnvironment(ctx, repo.Owner, repo.Name, name)
	return err
}
