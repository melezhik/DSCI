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

* ??? ( i don't major ones, but let me know ;-)


# Show me the code

In your source code repository just drop few tasks:

```
.
├── task.py
└── tasks
    ├── task_one
    │   └── task.bash
    └── task_two
        └── task.bash
```


*task.py*

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


*task_one/task.bash*

```bash
#!/bin/bash

echo "hello from task one you passed: ${foo}|${bar}"
```
