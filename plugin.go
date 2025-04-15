package main

import (
	"context"
	"errors"
	"fmt"

	"codeberg.org/woodpecker-plugins/go-plugin"
	"github.com/rs/zerolog/log"
	"github.com/woodpecker-ci/plugin-deployments/forge"
)

type Plugin struct {
	*plugin.Plugin
	settings struct {
		url        string
		name       string
		action     string
		forgeToken string
	}
}

func (p *Plugin) execute(ctx context.Context) error {
	token := p.settings.forgeToken
	_forge, err := forge.GetForge(p.Metadata.Forge, token)
	if err != nil {
		return err
	}

	deploymentURL := p.settings.url
	if deploymentURL == "" {
		return errors.New("deployment url is required")
	}

	deploymentName := p.settings.name
	if deploymentName == "" {
		switch p.Metadata.Pipeline.Event {
		case "pull_request", "pull_request_closed":
			deploymentName = fmt.Sprintf("pr-%s", p.Metadata.Commit.PullRequest)
		case "tag":
			deploymentName = p.Metadata.Commit.Tag
		case "push":
			deploymentName = p.Metadata.Commit.Branch
		default:
			return errors.New("please set a deployment name")
		}
	}

	action := p.settings.action
	if action == "" {
		if p.Metadata.Pipeline.Event == "pull_request_closed" {
			action = "delete"
		} else {
			action = "create"
		}
	}

	switch action {
	case "create":
		err := _forge.CreateDeployment(ctx, p.Metadata.Repository, deploymentName, deploymentURL, &p.Metadata)
		if err != nil {
			return err
		}
		log.Info().Msgf("🚀 Deployment created: %s (%s)", deploymentName, deploymentURL)

	case "delete":
		err := _forge.RemoveDeployment(ctx, p.Metadata.Repository, deploymentName)
		if err != nil {
			return err
		}
		log.Info().Msgf("🛌🏽 Deployment removed: %s", deploymentName)

	default:
		return errors.New("invalid action")
	}

	return nil
}
