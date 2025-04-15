# plugin-deployments

[![Build status](https://ci.woodpecker-ci.org/api/badges/woodpecker-ci/plugin-deployment/status.svg)](https://ci.woodpecker-ci.org/woodpecker-ci/plugin-deployment)
[![Docker Image Version (latest by date)](https://img.shields.io/docker/v/woodpeckerci/plugin-deployment?label=DockerHub%20latest%20version&sort=semver)](https://hub.docker.com/r/woodpeckerci/plugin-deployment/tags)

Woodpecker plugin that allows you to update deployment entires in your forge. You can use it for example after deploying a review environment to publish the review environments url to your forge.

## Build

Build the Docker image with the following command:

```sh
docker build -f Dockerfile -t woodpeckerci/plugin-codecov:next .
```
