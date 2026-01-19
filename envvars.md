# DSCI Environment variables

## DSCI_FEEDBACK_URL

Feedback URL. Default value is `http://127.0.0.1:4000`

Example:

```bash
DSCI_FEEDBACK_URL=https://sparky.sparrowhub.io
```

## DSCI_ORCHESTRATOR_NO_CONTAINER

DSCI orchestrator is running in VM/Base metal host, not on container

Example:

```bash
DSCI_ORCHESTRATOR_NO_CONTAINER=1
```

## DSCI_AGENT_SKIP_BOOTSTRAP

Skip DSCI agent ( job runner ) bootstrap

Example:

```bash
DSCI_AGENT_SKIP_BOOTSTRAP=1
```

##  DSCI_AGENT_IMAGE

DSCI agent image. Default values is `alpine:latest`

Example:

```bash
DSCI_AGENT_IMAGE=melezhik/sparrow:ubuntu_arm
```

