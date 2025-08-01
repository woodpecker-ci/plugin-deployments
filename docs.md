---
name: Deployments plugin
author: Woodpecker Authors
description: Update deployments in your forge
tags: [env, semver]
containerImage: woodpeckerci/plugin-deployments
containerImageUrl: https://hub.docker.com/r/woodpeckerci/plugin-deployments
url: https://github.com/woodpecker-ci/plugin-deployments
---

# plugin-deployments

The deployments plugin allows you to update deployment entires in your forge. You can use it for example after deploying a review environment to publish the review environments url to your forge.

The below pipeline configuration demonstrates simple usage:

```yml
steps:
  extend-env:
    image: woodpeckerci/plugin-deployments
    settings:
      url: https://may-review-environment.example.com
      # action: create # This option is normally not necessary as its auto-detected by the pipeline event
      forge_token:
        from_secrets: github_token
```

## Settings

| Settings      | Default                                                                       | Description                         |
| ------------- | ----------------------------------------------------------------------------- | ----------------------------------- |
| `ACTION`      | `create` for all pipeline events apart from `pull_request_closed` => `delete` | `create` or `delete` a deployment   |
| `NAME`        | pull-requests: `pr-{pr-number}`, tag: `{tag-name}`, push: `{branch}`          | The name of your deployment         |
| `URL`         | _none_                                                                        | The url to the deployed environment |
| `FORGE_TOKEN` | _none_                                                                        | A token to access the forges api    |
