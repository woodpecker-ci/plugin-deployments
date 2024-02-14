# plugin-deployments

<p align="center">
  <a href="https://ci.woodpecker-ci.org/woodpecker-ci/plugin-extend-env" title="Build Status">
    <img src="https://ci.woodpecker-ci.org/api/badges/woodpecker-ci/plugin-extend-env/status.svg" alt="Build Status">
  </a>
  <a href="https://goreportcard.com/badge/github.com/woodpecker-ci/plugin-extend-env" title="Go Report Card">
    <img src="https://goreportcard.com/badge/github.com/woodpecker-ci/plugin-extend-env" alt="Go Report Card">
  </a>
  <a href="https://godoc.org/github.com/woodpecker-ci/plugin-extend-env" title="GoDoc">
    <img src="https://godoc.org/github.com/woodpecker-ci/plugin-extend-env?status.svg" alt="GoDoc">
  </a>
  <a href="https://hub.docker.com/r/woodpeckerci/plugin-extend-env" title="Docker pulls">
    <img src="https://img.shields.io/docker/pulls/woodpeckerci/plugin-extend-env" alt="Docker pulls">
  </a>
  <a href="https://opensource.org/licenses/Apache-2.0" title="License: Apache-2.0">
    <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="License: Apache-2.0">
  </a>
</p>

The deployments plugin allows you to update deployment entires in your forge. You can use it for example after deploying a review environment to publish the review environments url to your forge.

## Usage

```yml
steps:
  add-deployment:
    image: woodpeckerci/plugin-deployments
    settings:
      url: https://may-review-environment.example.com
      # action: create # This option is normally not necessary as its auto-detected by the pipeline event
```
