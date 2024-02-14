package main

import (
	"context"
	"errors"
	"fmt"

	"codeberg.org/woodpecker-plugins/go-plugin"

	"github.com/woodpecker-ci/plugin-deployments/forge"
)

type Plugin struct {
	*plugin.Plugin
}

func (p *Plugin) execute(ctx context.Context) error {
	var _forge forge.Forge

	token := "" // TODO

	switch p.Metadata.Forge.Type {
	case "gitlab":
		var err error
		_forge, err = forge.NewGitlab(p.Metadata.Forge.Link, token)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported forge type")
	}

	action := "" // TODO
	if action == "" {
		if p.Metadata.Pipeline.Event == "pull_request_closed" {
			action = "delete"
		} else {
			action = "create"
		}
	}

	deploymentURL := "" // TODO
	if deploymentURL == "" {
		return errors.New("deployment url is required")
	}

	deploymentName := "" // TODO
	if deploymentName == "" {
		if p.Metadata.Pipeline.Event == "pull_request" || p.Metadata.Pipeline.Event == "pull_request_closed" {
			deploymentName = fmt.Sprintf("pr-%s", p.Metadata.Curr.PullRequest)
		} else if p.Metadata.Pipeline.Event == "tag" {
			deploymentName = p.Metadata.Curr.Tag
		} else if p.Metadata.Pipeline.Event == "push" {
			deploymentName = p.Metadata.Curr.Branch
		} else {
			return errors.New("please set a deployment name")
		}
	}

	switch action {
	case "create":
		err := _forge.CreateEnvironment(ctx, p.Metadata.Repository, deploymentURL, deploymentName)
		if err != nil {
			return err
		}

	case "delete":
		err := _forge.RemoveEnvironment(ctx, p.Metadata.Repository, deploymentURL, deploymentName)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid action")
	}

	return nil
}
