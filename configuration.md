# DSCI runner configuration variables

Dsci configuration settings are defined at `~/.dsci.toml` in TOML format

## DsciFeedbackUrl

Feedback URL. Default value is `http://127.0.0.1:8080`

Example:

```toml
DsciFeedbackUrl: "http://dsci-runner.sparrowhub.io"
```

## DsciAgentSkipBootstrap

Skip DSCI agent bootstrap

Example:

```bash
DsciAgentSkipBootstrap: true
```

##  DsciAgentImage

DSCI agent image. Default values is `alpine:latest`

Example:

```bash
DsciAgentImage: "dsci-agent-ubuntu:latest"
```

