# Tasks

To create a task, create a task folder within `tasks/` directory and within this folder task file named as `task.$ext`, where $ext is one of supported by SDK languages extension, for example:

*tasks/task_one/task.python*

```python
print("hello from task")
```

*tasks/task_two/task.bash*

```bash
echo "hello from task"
```

Following is a list of tasks file names for different languages:

```
+------------+--------------+
| Language   | File         |
+------------+--------------+
| Raku       | task.raku    |
| Perl       | task.pl      |
| Bash       | task.bash    |
| Python     | task.py      |
| Ruby       | task.rb      |
| Powershell | task.ps1     |
| Golang     | task.go      |
+------------+--------------+
```

Tasks are called from [job file](https://github.com/melezhik/DSCI/tree/main/job.md) using `run_task()` function, it takes the first input parameter as a path to sub directory within `tasks/` directory:

```python
#!/bin/python3
run_task("task_one"); # will run task inside tasks/task_one/ folder
run_task("task_two"); # will run task inside tasks/task_one/ folder
```

Tasks may take some input variables by passing the second argument to `run_task()`function:

```python
run_task("task_one", { "foo" : "bar" });
```

Task variable is handled inside tasks by calling `task_var()` function

```python
foo = task_var("foo")
```

## Job parameters

Tasks may also get an access to job input parameters by using `config()` function, say job has a parameter named `param1` then task may access the parameter like this:


Python example:

```python
job_params = config()
print(job_params["param1"])
```

Bash example:

```bash
param1=$(config param1)
```

`config()` function signature for supported languages:

```
+-------------+-------------------+
| Language    | signature         |
+-------------+-------------------+
| Raku        | config()          |
| Perl        | config()          |
| Bash        | config()          |
| Python      | config()          |
| Ruby        | config()          |
| Powershell  | config()          |
+-------------+-------------------+
```

## Exchange state between tasks

Tasks may exchange data between tasks within the same job by using `update_state()`,`get_state()` functions.

Say, task_one set two parameters (out1, out2) like this:

```python
#!/usr/python3

update_state({
  'out1' : 'out1 value',
  'out2' : 'out2 value'
})
```

Then any other task may pick it up by using `get_state()` function

```python
#!/bin/python

dict = get_state()
print(dict["out1"])
```

`update_state` function signatures for Sparrow6 supported languages:

```
+-------------+-----------------------------+
| Language    | signature                   |
+-------------+-----------------------------+
| Raku        | update_state(array|hash)    |
| Perl        | update_state(array|hash)    |
| Bash(*)     | update_state(key,value)     |
| Python      | update_state(array|hash)    |
| Ruby        | update_state(array|hash)    |
| Powershell  | update_state(array|hash)    |
+-------------+-----------------------------+
```

(*) Bash has a limited key/value support only:

```bash
#!bash
update_state "cnt" 100
```

`get_state()` function signature for supported languages:

```
+-------------+-------------------+
| Language    | signature         |
+-------------+-------------------+
| Raku        | get_state()       |
| Perl        | get_state()       |
| Bash        | not supported     |
| Python      | get_state()       |
| Ruby        | get_state()       |
| Powershell  | get_state()       |
+-------------+-------------------+
```

# Further reading 

* [Pipelines](https://github.com/melezhik/DSCI/tree/main/pipeline.md)

* [Jobs](https://github.com/melezhik/DSCI/tree/main/job.md)
