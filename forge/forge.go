package forge

import (
	"context"

	"codeberg.org/woodpecker-plugins/go-plugin"
)

type Forge interface {
	CreateEnvironment(ctx context.Context, repo plugin.Repository, url, name string) error
	RemoveEnvironment(ctx context.Context, repo plugin.Repository, url, name string) error
}
