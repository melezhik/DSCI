# DSCI

[D]ead [S]imple CI

---

We are on discord, [~join](https://discord.gg/KSMRTZ9F) us!

# Feature

\[Pros\]

* no painful YAML pipelines, use general purpose programming languages for your pipelines instead

* fully integrated with big guns like [~Forgejo](/doc/forgejo-setup)/Codeberg/Gitea/Gitlab/Github

* repeatable steps are written as plugins on general purpose programming languages 
and used as is

* for developers - pipelines get run as normal scripts (useful for debugging)

\[Limitations\]

* no job dependencies

* ??? ( i don't know major ones, but let me know ;-)


# Show me the code

In your source code repository just drop few tasks under `.dsci` directory:

```
.
├── job_one
│   ├── job.py
│   └── tasks
│       ├── task_one
│       │   └── task.py
│       └── task_two
│           └── task.bash
├── job_two
│   ├── job.bash
│   └── tasks
│       └── task_one
│           └── task.bash
└── jobs.yaml
```


*job_one/job.py*

```python
#!/bin/python3
run_task(
  'task_one', {
      'foo' : 'foo value',
      'bar' : 'bar value'
    }
);

run_task(
  'task_two', {
      'foo' : 'foo value',
      'bar' : 'bar value'
    }
);
```

*job_one/tasks/task_one/task.py*

```python
print(task_var("foo"))
print(task_var("bar"))
```

*job_one/tasks/task_two/task.bash*

```bash
#!/bin/bash

echo "hello from task one you passed: ${foo}|${bar}"
```

*job_one/job.bash*

```bash
#!/bin/bash
run_task "task_one"
```

*jobs.yaml*

```yaml
# runs job1, job2 sequentially
jobs:
  - 
    id: job1
    path: job_one/
  - 
    id: job2
    path: job_two/
```

## Passing states between tasks

Just use update\_state function within any tasks/* task to set some data:

```python
#!/usr/python3

update_state({
  'out1' : 'out1 value',
  'out2' : 'out2 value'
})
```

Then pick it up within any other tasks/* task by using get\_state function

```python
#!/bin/python

dict = get_state()
print(dict["out1"])
```

## Passing states between jobs

Just use update\_state function within any tasks/* task to set some data.

For example within job1, task_one:

```python
#!/usr/python3

update_state({
  'out1' : 'out1 value',
  'out2' : 'out2 value'
})
```

Then pick it up within any other jobs tasks/* task by using config function

```python
#!/bin/python

dict = config()

print(dict["_dsci_"]["job1"]["out1"])
```

# Programming languages supported

The same SDK for those programming languages:

* Bash
* Python
* Golang
* Ruby
* Powershell
* Raku
* Perl5

Choose the one you like and use it for pipeline. No extra code is required!

TBD - pipeline examples

# Forgejo integration

Use DSCI instead of YAML pipeline within your forgejo instance. Pipeline reports
are available directly via forgejo. See [~instruction](/doc/forgejo-setup)

# Codeberg/Gitea/Gitlab/Github integration

TBD - installation/configuration howto

# Using job plugins

There are a plenty of job plugins for common tasks:

- installing database, services, etc

For example:

*jobs.yaml*

```yaml
jobs:
  - 
    # install mariadb database
    id: db
    plugin: mariadb
  - 
    id: job1
    path: job_one/
   - 
    id: job2
    path: job_two/
  - 
```

Job plugins act as native jobs, for example can use get\_state/update\_state function.

TDB - example of how to write job plugin (pretty much as any job)

# Containers support

Jobs are executed either on ephemeral containers (alpine:latest)

# Running pipelines locally

* enable debug option for a job, f.e.

```yaml
jobs:
  - 
    id: job1
    path: job_one/
    debug: true
```

The run job and copy job effective configuration from output.

* Paste configuration into some file, f.e.`.config.json`

* Run job locally

```bash
docker run  \
--env SP6_TASK_CONFIG_FROM=.config.json 
\-it 
\--entrypoint /bin/bash \
-v $PWD:/opt/job dsci  \
-c "cd /opt/job/; s6 --task-run ."
```

# Pipelines vs Jobs vs Tasks vs Plugins

- [~Pipeline](/doc/pipeline) is a list of jobs executed sequentially

- [~Job](/doc/job) is a list of tasks executed sequentially

- Isolation:

Jobs are executed on isolated environments, while [~tasks](/doc/task) within a certain job are executed in the same enviorment

- Environments are represented by ephemeral containers

- Jobs/tasks can share/pass states

- Plugins are the same as jobs, but reusable jobs published to https://sparrowhub.io

# Screenshots

* DSCI reports

![UI](https://raw.githubusercontent.com/melezhik/DSCI/refs/heads/main/screenshots/ui.jpeg)

* Forgejo commit status via DSCI

![Commit check](https://raw.githubusercontent.com/melezhik/DSCI/refs/heads/main/screenshots/commit-check.jpeg)

# Examples

See [~examples](https://github.com/melezhik/DSCI/tree/main/examples) folder
