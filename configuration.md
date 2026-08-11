# DSCI runner configuration variables

Dsci configuration settings are defined at `~/.dsci.toml` in TOML format

## DsciContainerRuntime

Container runtime, should be one of these: `docker|podman`, docker is used if not set

Example:

```toml
DsciContainerRuntime = "podman"
```

## DsciAgentSkipBootstrap

Skip DSCI agent bootstrap

Example:

```toml
DsciAgentSkipBootstrap = true
```

##  DsciAgentImage

DSCI agent image. Default values is `alpine:latest`

Example:

```toml
DsciAgentImage = "dsci-agent-ubuntu:latest"
```

## GitPathToHttpBackend

Path to git-http-backend binary. Optional. Default value is `/usr/lib/git-core/git-http-backend`

```toml
GitPathToHttpBackend = "/Users/alex/homebrew/Cellar/git/2.33.0/libexec/git-core/git-http-backend"
```
