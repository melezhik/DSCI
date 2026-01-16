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

Jobs contain list of jobs to be executed sequentially. Every job has to have a unique id. If path is set job source is taken from path directory, see [job.md](job.md)

Job list may contain plugins instead of regular jobs:

```yaml
jobs:
  -
    id: build_essential
    plugin: build-essential
```

Plugins are reusable jobs, see http://sparrowhub.io/plugins for available plugins

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
    id: build_essential
    plugin: build-essential
    params:
      sudo: {}
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

# Further reading 

[Jobs](https://github.com/melezhik/DSCI/tree/main/job.md)
