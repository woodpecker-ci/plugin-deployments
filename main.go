package main

import (
	"codeberg.org/woodpecker-plugins/go-plugin"
	"github.com/urfave/cli/v3"
)

func main() {
	p := &Plugin{}
	p.Plugin = plugin.New(plugin.Options{
		Name:        "plugin-deployments",
		Description: "Update deployments in your forge",
		Version:     "v0.0.1",
		Execute:     p.execute,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "url",
				Usage:       "The URL of the deployment",
				Sources:     cli.EnvVars("PLUGIN_URL"),
				Destination: &p.settings.url,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "The name of the deployment",
				Sources:     cli.EnvVars("PLUGIN_NAME"),
				Destination: &p.settings.name,
			},
			&cli.StringFlag{
				Name:        "action",
				Usage:       "Should we create or delete the deployment?",
				Sources:     cli.EnvVars("PLUGIN_ACTION"),
				Destination: &p.settings.action,
			},
			&cli.StringFlag{
				Name:        "forge-token",
				Usage:       "The token to authenticate with your forge",
				Sources:     cli.EnvVars("PLUGIN_FORGE_TOKEN"),
				Destination: &p.settings.forgeToken,
			},
		},
	})

	p.Run()
}
