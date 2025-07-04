package forge

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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

func (g *Gitlab) getRepo(ctx context.Context, _repo plugin.Repository) (int, error) {
	// TODO: use repo-id to get the repo
	repo, _, err := g.Projects.GetProject(fmt.Sprintf("%s/%s", _repo.Owner, _repo.Name), nil, gitlab.WithContext(ctx))
	if err != nil {
		return 0, err
	}

	return repo.ID, nil
}

func (g *Gitlab) getEnvironment(ctx context.Context, repoID int, name string) (*gitlab.Environment, error) {
	envs, _, err := g.Environments.ListEnvironments(repoID, &gitlab.ListEnvironmentsOptions{
		Name: gitlab.Ptr(name),
	}, gitlab.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	if len(envs) == 1 {
		return envs[0], nil
	}

	return nil, nil
}

func (g *Gitlab) CreateDeployment(ctx context.Context, repo plugin.Repository, name, url string, metadata *plugin.Metadata) error {
	repoID, err := g.getRepo(ctx, repo)
	if err != nil {
		return err
	}

	env, err := g.getEnvironment(ctx, repoID, name)
	if err != nil {
		return err
	}

	if env == nil {
		_, _, err := g.Environments.CreateEnvironment(repoID, &gitlab.CreateEnvironmentOptions{
			Name:        gitlab.Ptr(name),
			ExternalURL: gitlab.Ptr(url),
		}, gitlab.WithContext(ctx))
		if err != nil {
			return err
		}
	} else {
		_, _, err = g.Environments.EditEnvironment(repoID, env.ID, &gitlab.EditEnvironmentOptions{
			ExternalURL: gitlab.Ptr(url),
		}, gitlab.WithContext(ctx))
		if err != nil {
			return err
		}
	}

	commit := metadata.Curr
	_, _, err = g.Deployments.CreateProjectDeployment(repoID, &gitlab.CreateProjectDeploymentOptions{
		Environment: gitlab.Ptr(name),
		Tag:         gitlab.Ptr(commit.Tag != ""),
		SHA:         gitlab.Ptr(commit.Sha),
		Ref:         gitlab.Ptr(commit.Ref),
		Status:      gitlab.Ptr(gitlab.DeploymentStatusValue("success")),
	}, gitlab.WithContext(ctx))

	if metadata.Pipeline.Event == "pull_request" {
		mergeRequestID, err := strconv.Atoi(metadata.Curr.PullRequest)
		if err != nil {
			return err
		}

		note, err := g.getComment(repoID, mergeRequestID, name)
		if err != nil {
			return err
		}

		if note == nil {
			_, _, err = g.Notes.CreateMergeRequestNote(repoID, mergeRequestID, &gitlab.CreateMergeRequestNoteOptions{
				Body: gitlab.Ptr(fmt.Sprintf("%s %s", g.getCommentPrefix(name), url)),
			}, gitlab.WithContext(ctx))
			if err != nil {
				return err
			}
		} else {
			_, _, err = g.Notes.UpdateMergeRequestNote(repoID, mergeRequestID, note.ID, &gitlab.UpdateMergeRequestNoteOptions{
				Body: gitlab.Ptr(fmt.Sprintf("%s %s", g.getCommentPrefix(name), url)),
			}, gitlab.WithContext(ctx))
			if err != nil {
				return err
			}
		}
	}

	return err
}

func (g *Gitlab) RemoveDeployment(ctx context.Context, repo plugin.Repository, name string) error {
	repoID, err := g.getRepo(ctx, repo)
	if err != nil {
		return err
	}

	env, err := g.getEnvironment(ctx, repoID, name)
	if err != nil {
		return err
	}

	if env != nil {
		env, _, err := g.Environments.StopEnvironment(repoID, env.ID, nil, gitlab.WithContext(ctx))
		if err != nil {
			return err
		}

		_, err = g.Environments.DeleteEnvironment(repoID, env.ID, gitlab.WithContext(ctx))
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Gitlab) getComment(projectID, mergeRequestID int, name string) (*gitlab.Note, error) {
	listMergeRequestNotesOptions := &gitlab.ListMergeRequestNotesOptions{
		Sort: gitlab.Ptr("asc"),
	}

	for {
		notes, resp, err := g.Client.Notes.ListMergeRequestNotes(projectID, mergeRequestID, listMergeRequestNotesOptions)
		if err != nil {
			return nil, err
		}

		for _, note := range notes {
			if strings.Contains(note.Body, g.getCommentPrefix(name)) {
				return note, nil
			}
		}

		// Exit the loop when we've seen all pages
		if resp.CurrentPage >= resp.TotalPages {
			break
		}

		// Update the page number to get the next page
		listMergeRequestNotesOptions.Page = resp.NextPage
	}

	return nil, nil
}

func (g *Gitlab) getCommentPrefix(name string) string {
	return fmt.Sprintf("🚀 Preview \"%s\" deployed to:", name)
}
