## Cargo, Jest project

> Let’s say, i’d want to run just test in this repo: https://codeberg.org/lutra/lutra/ How do I install just, cargo-insta & python version 3.14?

jobs.yaml

```yaml
global:
  container_image: python:3.14-alpine
jobs:
  -
    id: install_and_test
    path: test/
    
```

test/job.bash

```bash
run_task "install"
run_task "test"
```

test/tasks/install/task.bash

```bash
apk add cargo-insta just
```

test/tasks/install/task.bash

```bash
just test
```
