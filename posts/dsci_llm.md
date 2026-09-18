# Title

You can create any workflow or pipeline you want with DSCI and LLM

# Description

How to create any sort of pipelines using DSCI and LLM

# Tags

devops, ai, ci, LLM

---

[DSCI](https://deadsimpleci.sparrowhub.io) - is a self hosted git server with embedded CI runner. The cool thing about it that advances it position among other tools:

- You can use many popular programming languages to create CI scenarios, no need entangled YAML

- You can generate CI pipeline by talking to AI agent using public [DSCI sandbox](http://dsci.sparrowhub.io:8080/)

---

Here the thing. Just ask LLM what do you need and you'll get your answer in seconds:

```
Create DSCI pipeline to build golang project with unit tests and go vet
and publishing resulted container image based on Alpine to external container
registry

```

If you don't believe me those are examples I've created for myself, by just playing with DSCI and LLM:

---

* [Python / selenium tests ci pipeline](http://dsci.sparrowhub.io:8080/file_view/dsci/1c9fb55.1789580532/answer.md)

* [Execute chef cookbook on remote host](http://dsci.sparrowhub.io:8080/file_view/dsci/007aefe.1789618706/answer.md)

* [Execute Ansible playbook on remote host](http://dsci.sparrowhub.io:8080/file_view/dsci/007aefe.1789617896/answer.md)

* [C++ with cmake](http://dsci.sparrowhub.io:8080/file_view/dsci/007aefe.1789628812/answer.md)

* [Node.js with npm project](http://dsci.sparrowhub.io:8080/file_view/dsci/007aefe.1789654909/answer.md)

* [Create DSCI pipeline to deploy AI agent on localhost , use some popular Agent framework and LLM](http://dsci.sparrowhub.io:8080/file_view/dsci/007aefe.1789718838/answer.md)

* [Python+pytest+coverage](http://dsci.sparrowhub.io:8080/file_view/dsci/1c9fb55.1789570454/answer.md)

* [Golang+unit tests+ coverage + docker image push](http://dsci.sparrowhub.io:8080/file_view/dsci/1c9fb55.1789570619/answer.md)

* [Perl/Mariadb service/Sqlite/Curl](http://dsci.sparrowhub.io:8080/file_view/dsci/1c9fb55.1789573431/answer.md)

Now, what amazing about all those examples:

* They are pretty accurate in a sense of correctness of code they generated ( believe me I am - DSCI author )

* They are built in seconds 

* I use my free GROQ tokens to build them ( now this is part of DSCI sandbox service ) - that means - not a lot of tokens is required to build such a code

---

How to use it ?

Use this link - [http://dsci.sparrowhub.io:8080/repo/pipeline-generator.git](http://dsci.sparrowhub.io:8080/repo/pipeline-generator.git)

To build pipeline for your requirements, just say something:

> Build DSCI pipeline for typical Python project with pytest unit tests and code coverage more then X

Then just for a few seconds, go to the builds/ page and pick up the latest `dsci@*` build, click on artifacts links and get your answer.md

You can do more than that. AI agent is smart enough to nail it down for you.

---
