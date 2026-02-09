# Install

Self hosted installation guide

## dependencies

On machine (VM) running DSCI following dependencies should be installed:

- docker
- sshd
- forgejo

### Install forgejo

Follow forgejo documentation

### Install dsci runner

Right now dsci runner needs to be  built from source code, ready to use binaries
for various architectures coming soon:

```bash
git clone https://github.com/melezhik/dsci-runner.git
cd dsci-runner
go mod tidy
go build -o dsci_runner main.go
```

### Configure forgejo and dsci runner

Follow ~[this documentation](/doc/forgejo-setup)

### Create first pipeline

* Create repo, setup web hook

* Create pipeline code

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

Infrastructure pipeline unlike regular pipelines run on VM, not in docker.

They are used by server administrators to setup server running DSCI stack

Requirements:

One needs to ensure following for server running dsci runner:

* enable passwordless sudo for user running dsci runner

* open 22 port on 127.0.0.1

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

Note: 

# Further setup

Normally everything is done via dsci [pipelines](~/doc/pipeline) which in nutshell are just Bash or Python scripts,
so developers just need to write them using guidelines or use some existing dsci plugins to get the job done
