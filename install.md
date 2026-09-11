# Install

Self hosted installation guide

## dependencies

On machine (VM) running DSCI following dependencies should be installed:

- docker/podman
- rakupp

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

### Create `~/.dsci.toml`

Required parameters to set:

- DsciContainerRuntime

This sets container runtime - `docker` or `podman`

- GitAuthUser/GitAuthPassword

This sets login and password for git http basic authentication

- GitPathToHttpBackend

This sets path to your git http backend cli

- DsciAgentImage

This has to be set to `melezhik/dsci-agent-rakupp`.

Recommend to use this image as it has the fastest dsci rumtime

- DsciAgentSkipBootstrap

This has to be set to true

---

Final configuration file will be something like that:

```toml
DsciAgentImage = "melezhik/dsci-agent-rakupp"
DsciContainerRuntime = "podman"
GitServerAddress = "http://127.0.0.1"
GitAuthUser = "admin"
GitAuthPassword = "superSecret"
DsciAgentSkipBootstrap = true
```

For options explanation see [configuration](/configuration.md) doc

### Install dsci server

```bash
git clone https://github.com/melezhik/dsci-runner.git
cd dsci-runner
go mod tidy
go build
./dsci-runner
```

Go to http://127.0.0.1:8080

### Install dsci job runner

```bash
mkdir -p ~/.rakupp
cd ~/.rakupp
wget https://github.com/ash/rakupp/releases/download/v3.26.0/rakupp-linux-x86_64.tar.gz
tar -xzf rakupp-linux-x86_64.tar.gz
export PATH=~/.rakupo/rakupp/bin/:~/.raku/bin:$PATH
rakupp install --no-test Sparky::Minimal
echo 'export PATH=~/.rakupp/rakupp/bin/:~/.raku/bin:$PATH' >> ~/.bashrc
echo 'export PATH=~/.rakupp/rakupp/bin/:~/.raku/bin:$PATH' >> ~/.bash_profile 
```

Initialize job runner database

```
sparkyd-init
```

Run job runner

```
sparkyd
```

### Create git repository

web ui -> new repo 

`demo.git`

### Use git repository

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

Normally everything is done via dsci [pipelines](/pipeline.md) which in nutshell are just Bash or Python scripts, so developers just need to write them using guidelines or use some existing dsci [plugins](/bash-plugins.md) to get the job done
