# DSCI runner configuration variables

Dsci configuration settings are defined at `~/.dsci.toml` in TOML format

## DsciContainerRuntime

Container runtime, should be one of these: `docker|podman`, docker is used if not set

Example:

```toml
DsciContainerRuntime: "podman"
```

## DsciFeedbackUrl

Feedback URL. Default value is `http://127.0.0.1:8080`

Example:

```toml
DsciFeedbackUrl: "http://dsci-runner.sparrowhub.io"
```

## DsciAgentSkipBootstrap

Skip DSCI agent bootstrap

Example:

```toml
DsciAgentSkipBootstrap: true
```

##  DsciAgentImage

DSCI agent image. Default values is `alpine:latest`

Example:

```toml
DsciAgentImage: "dsci-agent-ubuntu:latest"
```

