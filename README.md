# DSCI

[D]ead [S]imple CI


# Feature

\[Pros\]

* no painful YAML pipelines, use general purpose programming languages for your pipelines instead

* fully integrated with big guns like Forgejo/Codeberg

* repeatable steps are written as plugins on general purpose programming languages 
and used as is

* for developers - pipelines get run as normal scripts (useful for debugging)

\[Limitations\]

* no job dependencies

* ??? ( i don't know major ones, but let me know ;-)


# Show me the code

In your source code repository just drop few tasks:

```
.
├── job_one
│   ├── task.py
│   └── tasks
│       ├── task_one
│       │   └── task.bash
│       └── task_two
│           └── task.bash
└── job_two
    ├── task.bash
    └── tasks
        └── task_one
            └── task.bash
```


*job_one/task.py*

```python
#!/bin/python3
run_task(
  'task_one', {
      'foo' : 'foo value'
      'bar' : 'bar value'
    }
);

run_task(
  'task_two', {
      'foo' : 'foo value'
      'bar' : 'bar value'
    }
);
```


*job_one/task_one/task.bash*

```bash
#!/bin/bash

echo "hello from task one you passed: ${foo}|${bar}"
```

*jobs.yaml*

```yaml
# runs job_one, job_two sequentially
jobs:
  - job_one
  - job_two
```

## Passing states between tasks/jobs

Just use update\_state/get\_state functions within any tasks/* task:

```python
#!/usr/python3

update_state, {
  'out1' : 'out1 value'
  'out2' : 'out2 value'
}
```

Then pick it up within any other tasks/* task

```python
#!/bin/python

dict = get_state()
print(dict["out1"])

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

Choose the one you like an use for pipeline. Now extra code is required!

TBD - pipeline examples

# Forgejo integration

Use DSCI instead of YAML pipeline within your forgejo instance. Pipeline reports
are available directly via forgejo

TBD - installation/configuration howto

# Codeberg integration

TBD - installation/configuration howto

# Using job plugins

There are a plenty of job plugins for common tasks:

- installing database, services, etc

For example:

*jobs.yaml*

```
jobs:
  # install mysql database
  - job_plugin "mysql" 
  - job_one
  - job_two
```

Jon plugins acts as native jobs, for example can use get\_state/update\_state function.

TDB - example of how to write job plugins (pretty much as any job task)

# Containers support

Jobs are executed on either ephemeral containers on directly on VM


# Running pipelines locally

Just install dsci cli as docker container and then:

```bash
#!/bin/bash
docker run dsci job_one 
docker run dsci job_two
docker run dsci job_one task_two foo="foo_value"  bar="bar value"
```
