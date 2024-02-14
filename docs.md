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

The extend env plugin extends an existing or creates a new `.env` file with additional variables like semver information.

The below pipeline configuration demonstrates simple usage:

```yml
steps:
  extend-env:
    image: woodpeckerci/plugin-deployments
    settings:
      url: https://may-review-environment.example.com
      # action: create # This option is normally not necessary as its auto-detected by the pipeline event
      # forge_url: https://gitlab.com
      # forge_type: gitlab
      # forge_username: ignored
      # forge_password: personal-api-token
```
