package main

import (
	"codeberg.org/woodpecker-plugins/go-plugin"
)

func main() {
	p := &Plugin{}
	p.Plugin = plugin.New(plugin.Options{
		Name:        "extend-env",
		Description: "Extend .env file with additional variables like semver",
		Version:     "v0.0.1",
		Execute:     p.execute,
	})

	p.Run()
}
