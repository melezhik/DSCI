# Install

Self hosted installation guide

## dependencies

On machine (VM) running DSCI following dependencies should be installed:

- docker/podman
- sshd


### Configure docker/podman

For docker:

No specific setup

For podman:

- make sure you've logged into public docker registry

```bash
podman login  -u user 
```

- make sure podman containers persist even when user sessions are gone

```
loginctl enable-linger $USER
```

- Set DsciContainerRuntime in `~/.dsci.toml` file

see [configuration](/configuration.md)

- Set GitAuthUser/GitAuthPassword in `~/.dsci.toml` file

see [configuration](/configuration.md)

### Install dsci server

```bash
git clone https://github.com/melezhik/dsci-runner.git
cd dsci-runner
go mod tidy
go build
./dsci-runner
```

Go to http://127.0.0.1:8080

### Create git repository

Right now, this is done via terminal, in the future this will be available via UI

To clone existing repository:

```bash
cd dsci-runner/.repositories
git clone --bare https://github.com/foo/demo.git
```

To create empty repository

```bash
cd dsci-runner/.repositories
git init --bare demo.git
```

Now you can close your repositories and work with them:

```
git clone http://127.0.0.1:8080/demo.git
cd demo
touch file.txt
git add file.txt && git commit -a -m "add file"
git push
```

### Create first pipeline

* Create pipeline code

In git repo, create `.dsci` folder with pipeline:

`.dsci/jobs.yaml`

```yaml
jobs:
    -
        id: job1
        path: .
```

`.dsci/task.py`

```python
print("hello world")
```

* Push it

* See results

### Setup infrastructure pipeline

Infrastructure pipelines unlike regular pipelines run on VM, not in docker.

They are used by server administrators to setup server running DSCI stack

Requirements:

One needs to ensure following for server running dsci runner:

* enable passwordless sudo for user running dsci runner

* open 22 port on 127.0.0.1

* add git repository with infrastructure pipeline to the white list

Set DsciAllowLocalhostModeRepos in  ~/.dsci.toml and restart dsci runner

```toml
# repos should be added in owner/repo
# format
DsciAllowLocalhostModeRepos = [
  "root/test3",
  "root/test4",
  "root/test2"
]
```

to make pipeline run on localhost use `global.localhost` modifier:

`.dsci/jobs.yaml`

```yaml
global:
  localhost: true
jobs:
  -
    id: job1
    path: .
```

# Further setup

Normally everything is done via dsci [pipelines](/pipeline.md) which in nutshell are just Bash or Python scripts, so developers just need to write them using guidelines or use some existing dsci [~plugins](/bash-plugins.md) to get the job done
