---
title: Dead Simple CI Introduction 
published: true
description: how to make use of dsci to write ci pipelines
tags: cicd, pipelines, forgejo
# cover_image: https://direct_url_to_image.jpg
# Use a ratio of 100:42 for best results.
# published_at: 2026-02-13 18:41 +0000
---

[Dsci](http://deadsimpleci.sparrowhub.io/doc/README) is a brand new kid on the cicd area. It allows use of general programming languages to create ci scenarios.

——-

Consider some imaginary case of building and deploying an application as docker container. So, logically we have two main stages:

- build docker image from source code 

- deploy docker image using docker pull and docker run commands

Let’s see how dsci could tackle such a scenario …

——-

# Pipeline file

The main concept we start with is so called ‘pipeline file’.


It’s just yaml file containing a list of jobs to be performed. Every job is executed on isolated docker container, execution logic is always sequential - one job, by another.

This pattern yet simple however covers the majority of use cases:

jobs.yaml

```yaml
jobs:
  -
     id: build
     path: build/
  -
     id: deploy
     path: deploy/    
```

Thus, every element of job list is a job, where job has some unique identifier and path to location of job file.

# Conditional job logic 

Usually there are some conditions when and what ci scripts are triggered. Those conditions may be based on different criteria, but usually they depend on branch names.

Dsci utilizes a special syntax to define those rules. Let’s say we only want to deploy if source branch is main, and perform build for all branches:

```yaml
jobs:
  -
     id: build
     path: build/
  -
     id: deploy
     path: deploy/ 
     only: .<ref> eq "refs/heads/main"
```

Read more on job conditions on dsci documentation - https://github.com/melezhik/DSCI/blob/main/job-cond.md

# Passing parameters to job

To pass parameters to job, just use `params:` key:

```yaml
jobs:
  -
     id: build
     path: build/
     params:
       foo: bar
       version: 0.1.0
  -
     id: deploy
     path: deploy/    
```

Those parameters could be thought as overrides for default ones, however dsci allows more then that - pass results ( states ) from one job to another, see later.

# Job file

The next important building block of the hierarchy is so called “job file” defining main job logic.

Job file needs to be named depending on language of choice, in our example we use Bash as it usually is enough for many use cases. But using other languages is also possible. If one needs more flexibility they may chose Python or Golang

build/task.bash

```bash
version=$(config version)
tag_version=$(date +%s).$version
docker build . -t app:$tag_version
docker push repo/app:$tag_version
```

In this simple example our job file is just a single task job. However if there is a need to split complex job into several tasks they may do so by using job.$ext file approach. This file needs to be named according language of choice. 

Let’s say we have three tasks - configure, build and test - incorporated into build job, we may run them one by one like this:

build/job.py

```python
#!/bin/python
run_task(“configure”);
run_task(“build”);
run_task(“test”);
```

And if we put those tasks under build/tasks/ directory like this:

build/tasks/configure/task.bash
build/tasks/build/tasks.bash
build/tasks/tests/task.bash

In this case we have modular setup of our ci job.  

The neat thing about this DSL, dsci provides the same SDK for all supported programming languages.


Read more about jobs and tasks on dsci documentation web site:

- http://deadsimpleci.sparrowhub.io/doc/task

- http://deadsimpleci.sparrowhub.io/doc/job

# Passing results between jobs

Let’s make our last job example more realistic and return docker image tag dynamically created by build job back to deploy job. We may want to do so deploy jobs know what tag to pull before deploy. All we need is to add extra line to save the job state into internal dsci cache:

build/task.bash

```bash
version=$(config version)
tag_version=$(date +%s).$version
docker build . -t app:$tag_version
docker push repo/app:$tag_version
update_state tag $tag_version
```

update_state() is very handy function allowing to pass states between different tasks and jobs. It’s implemented for all supported languages.

To pick up tag name in deploy job we can use already mentioned config() function:

deploy/task.bash

```bash
set -e
tag=$(config tag)
docker pull repo/app@$tag
docker stop -t 1 container || :
docker run -rm -name container -td repo/app@$tag
```

Read more about job and tasks states on dsci documentation ( the links above ^^ )

# Secrets 

Let’s modify build job example with pushing image to docker registry. To do a push we need to authenticate against a docker registry first. Dsci enables simple way to pass secrets to pipelines.

This time let’s rewrite our job file on Python for convenience:

build/job.py

```python
#!/bin/python
import os
password  = os.environ['password'])

run_task(“build”, { “password”: password })
```

And then modify build task to handle password parameter:

```bash
version=$(config version)
tag_version=$(date +%s).$version
docker build . -t app:$tag_version
update_state tag $tag_version

echo $password | docker login --username your-username --password-stdin

docker push repo/app:$tag_version
```

As we can see secrets passed into pipelines as environment variables.

And unlike other task and jobs parameters they never saved in job reports or cache files.  

—-


# Deployment on localhost 

By default dsci pipelines run inside some docker container, this fits situation when one needs run purely ci code - for example build and run some unit tests. CD part comes to play when build artifacts are ready for deploy. 

Dsci allows switch to deployment environment by using `localhost` switcher:

```yaml
jobs:
  -
     id: build
     path: build/
     params:
       foo: bar
       version: 0.1.0
  -
     id: deploy
     path: deploy/
     localhost: true 
```
In that case deployment occurs on VM running dsci orchestrator and which allows to restart docker container with new image version right on VM

——-

That is it. 

This simple but real life example shows how easy and one can write cicd pipelines using dsci framework and how flexible it is.

Hopefully you like it. 

—-

For comprehensive documentation and more information please visit dsci web site - http://deadsimpleci.sparrowhub.io/doc/README
