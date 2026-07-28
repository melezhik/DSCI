---
title: DSCI series / Rakulang CI
published: true
description: How to build CI pipelines for Rakulang using DSCI
tags: rakulang, ci, devops, automation
# cover_image: https://direct_url_to_image.jpg
# Use a ratio of 100:42 for best results.
# published_at: 2026-07-28 08:02 +0000
---

In this series I'm going to show how one can use [DSCI](https://github.com/melezhik/DSCI) to build CI/CD pipelines for various languages and stacks. And in today's episode we talk about [Raku](https://raku.org)

---

Building Raku module usually involves installing dependencies and running unit tests:

`.dsci/jobs.yaml`

```yaml
jobs:
  - 
    id: ci
    path: .
```

`.dsci/task.bash`

```bash
set -e
cd ../
ls -l
zef install . --deps-only
zef test .
echo "done"
```

In this simple Bash task we install module dependencies first and then run unit tests for the module. Easy step. Also zef and Raku is preinstalled on DSCI job runner, which makes things even easier for you.

---

Test coverage.

Say, we'd like run test coverage report to make it sure module has enough test coverage. We are going to use Test::Coverage (which excellently explained in Liz [post](https://dev.to/lizmat/towards-more-coverage-fne) ) for that. It's logical to add separate task that installs Test::Coverage first, before doing anything. Let's refactor out pipeline, as DSCI is just a tiny YAML layer and everything else is regular programming language, so it's easy to do that:

`.dsci/job.raku` - holds all the tasks executed for a job:

```perl
run_task "deps";
run_task "ci";
```

Tasks should reside under `tasks/` folder, so we have two tasks:


`.dsci/tasks/deps/task.bash`

```shell
zef install --/test Test::Coverage
```

And `.dsci/tasks/ci/task.bash` modified (from `.dsci/task.bash`) to:

```bash
set -e
cd ../
ls -l
zef install . --deps-only
raku -I. xt/coverage.rakutest
echo "done"
```

When I run the pipeline against my dummy rakudist-teddy-bear module, I get this expected [results](http://dsci-runner.sparrowhub.io/report/ui/job.run/boiqvgzxfjhwnedluapm.3403):


```console
09:59:18 :: All candidates are currently installed
09:59:18 :: # Failed test 'Coverage 33.33% >= 80%'
09:59:18 :: # at xt/coverage.rakutest line 5
09:59:18 :: # You failed 1 test of 2
09:59:18 :: task exit status: 1
09:59:18 :: task . FAILED
```

---

Switching on/off test coverage 

Sometimes module author might want to disable running test coverage on their code, for example when you're contributing to someone's module adding some lines of code, but still have not done with new tests for them yet, DSCI allows to tune pipeline in many ways and one of the simplest method is to rely on commit messages. So, let's say if a committer add `no_coverage` string into commit message, then test coverage logic is disabled, let's modify the pipeline according to new requirements:


`.dsci/job.raku`:

```perl
unless %*ENV<DSCI_MESSAGE> ~~ /no_coverage/ {
  run_task "deps";
}
run_task "ci", %(
  skip_coverage => %*ENV<DSCI_MESSAGE> ~~ /no_coverage/
);
``` 

And `.dsci/tasks/ci/task.bash`:

```bash
set -e
cd ../
ls -l
zef install . --deps-only
if test "$skip_coverage" = "Yes"; then
  zef test .
else 
 raku -I. xt/coverage.rakutest
fi
echo "done"
```

---

That's it. Thanks for reading. Stay tuned for other languages and stacks examples. As one can see DSCI is extremely flexible and easy to work with as it's just a normal programming languages experience with a tiny YAML layer on top of it (to configure workflow). Combining the best of two worlds - declarative DSL and regular programming.

