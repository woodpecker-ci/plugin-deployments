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
		err := _forge.CreateDeployment(ctx, p.Metadata.Repository, deploymentURL, deploymentName, &p.Metadata.Curr)
		if err != nil {
			return err
		}
		log.Info().Msgf("Deployment created: %s (%s)", deploymentName, deploymentURL)

	case "delete":
		err := _forge.RemoveDeployment(ctx, p.Metadata.Repository, deploymentName)
		if err != nil {
			return err
		}
		log.Info().Msgf("Deployment removed: %s", deploymentName)

	default:
		return errors.New("invalid action")
	}

	return nil
}
