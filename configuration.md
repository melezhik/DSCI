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

# Git related options

## GitAuthUser

Sets http basic authentication user which is used for git push. Optional. Default value is `dsci`

Example:

```toml
GitAuthUser = "admin"
```

## GitAuthPassword

Sets http basic authentication password which is used for git push. Optional. Default value is `dsci`

Example:

```toml
GitAuthPassword = "seCret1001"
```

## GitServerAddress

Sets git server address as it appears for public users using clone. Optional. Default value is `http://127.0.0.1:8080`

Example:

```toml
GitServerAddress = "http://dsci.sparrowhub.io"
```

## GitPathToHttpBackend

Path to git-http-backend binary. Optional. Default value is `/usr/lib/git-core/git-http-backend`

```toml
GitPathToHttpBackend = "/Users/alex/homebrew/Cellar/git/2.33.0/libexec/git-core/git-http-backend"
```
