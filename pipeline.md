# Pipelines

To create a pipeline, create `jobs.yaml` file inside `.dsci` folder in repository root:

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

Jobs contain list of jobs to be executed sequentially. Every job has to have a unique id. If path is set job source is taken from path directory, see [~job](job)

Job list may contain plugins instead of regular jobs:

```yaml
jobs:
  -
    id: build_essential
    plugin: build-essential
```

Plugins are reusable jobs, see [~http://sparrowhub.io](https://sparrowhub.io/search?q=all) for available plugins

## Parameters

Jobs and plugins may take input parameters:

```yaml
# runs job1, job2 sequentially
jobs:
  - 
    id: job1
    path: job_one/
    params:
        foo: bar
        size: 10
        colors:
            - blue
            - green
            - red
 -
    id: database
    plugin: mariadb
    params:
      db_user: alpine
      db_pass: SecReet
      db_name: products
```

Input parameters within job tasks are handled by the use of `config()` function available for all supported by SDK languages:

Bash:

```bash
foo=$(config foo)
```

Python:

```python
cfg = config()
print(cfg["foo"])
```

Raku

```raku
say config()<foo>;
```

---

## Single task job

Sometimes all you need is a single task job.

To create single task job just drop `task.$ext` file under job directory, 
no job file is required:

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

# Further reading 

- [~Jobs](/doc/job)

- [~Predefined pipeline variables](/doc/predefined-vars)

- [~Pipeline secrets](/doc/secrets)
